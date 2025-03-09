package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"th3-sh0p-api/database"
	"th3-sh0p-api/filesystem"
	"th3-sh0p-api/models"
	"th3-sh0p-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"google.golang.org/api/idtoken"
)

type Data struct {
	Data []*models.Image
}

func CreateImage(params operations.PostImageParams, principal interface{}) middleware.Responder {
	response := operations.NewPostImageOK()
	internalErr := operations.NewGetImagesInternalServerError()
	bearerHeader := params.HTTPRequest.Header.Get("Authorization")

	idToken := strings.Split(bearerHeader, " ")[1]
	payload, err := idtoken.ParsePayload(idToken)
	if err != nil {
		log.Printf("Failed to parse payload, %v", err)
		return internalErr
	}

	claims := payload.Claims
	email := claims["email"].(string)

	credit, err := database.ReduceUserCredit(email)
	if err != nil {
		log.Printf("Failed to user credit, %v", err)
		return internalErr
	}

	if credit == 0 {
		log.Println("User has no credit left")
		return operations.NewPostImageBadRequest()
	}

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
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPEN_AI"))
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
		ImageCredit: credit,
	})
	return response
}

func GetImages(params operations.GetImagesParams) middleware.Responder {
	page := params.Page
	response := operations.NewGetImagesOK()
	internalErr := operations.NewGetImagesInternalServerError()
	images, err := database.GetImages(page)
	if err != nil {
		log.Printf("Failed to get images, %v", err)
		return internalErr
	}
	response.SetPayload(images)
	return response
}

func GetImagesPages(params operations.GetImagesPagesParams) middleware.Responder {
	response := operations.NewGetImagesPagesOK()
	internalErr := operations.NewGetImagesInternalServerError()
	count, err := database.GetImageCount()
	if err != nil {
		log.Printf("Failed to get image count, %v", err)
		return internalErr
	}

	if count == float64(0) {
		count = float64(1)
	}

	pages := math.Ceil(count / 10)

	response.SetPayload(int64(pages))
	return response
}
