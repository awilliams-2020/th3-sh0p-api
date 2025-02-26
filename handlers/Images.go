package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"th3-sh0p-api/database"
	"th3-sh0p-api/filesystem"
	"th3-sh0p-api/models"
	"th3-sh0p-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type Data struct {
	Data []*models.Image
}

func CreateImage(params operations.PostImageParams) middleware.Responder {
	response := operations.NewPostImageOK()
	internalErr := operations.NewGetImagesInternalServerError()

	request := struct {
		Model   string `json:"model"`
		Prompt  string `json:"prompt"`
		Quality string `json:"quality"`
		Number  int    `json:"n"`
		Size    string `json:"size"`
	}{
		"dall-e-3",
		*params.Body.Prompt,
		"standard",
		1,
		"1024x1024",
	}

	_bytes, err := json.Marshal(request)
	if err != nil {
		log.Printf("Failed to marshal request into byes, %v", err)
		return internalErr
	}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/images/generations", bytes.NewBuffer(_bytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer")
	if err != nil {
		log.Printf("Failed to post generated image from open api, %v", err)
		return internalErr
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed do request for generated image from open api, %v", err)
		return internalErr
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body in %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: %v", string(body))
		return operations.NewPostImageBadRequest()
	}

	var openAI Data
	err = json.Unmarshal(body, &openAI)
	if err != nil {
		log.Printf("Failed to unmarshal response, %v", err)
		return internalErr
	}

	fileName, err := filesystem.SaveImage(openAI.Data[0].URL)
	if err != nil {
		log.Printf("Failed to save image to filesystem, %v", err)
		return internalErr
	}

	url := fmt.Sprintf("https://th3-sh0p.com/images/%s", fileName)

	ID, err := database.SaveImage(url)
	if err != nil {
		log.Printf("Failed to save image, %v", err)
		return internalErr
	}

	response.SetPayload(&operations.PostImageOKBody{
		Image: &models.Image{
			ID:  ID,
			URL: url,
		},
	})
	return response
}

func GetImages(params operations.GetImagesParams) middleware.Responder {
	response := operations.NewGetImagesOK()
	internalErr := operations.NewGetImagesInternalServerError()
	images, err := database.GetImages()
	if err != nil {
		log.Printf("Failed to get images, %v", err)
		return internalErr
	}
	response.SetPayload(images)
	return response
}
