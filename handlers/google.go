package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"th3-sh0p-api/database"
	"th3-sh0p-api/models"
	"th3-sh0p-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetGoogleProfile(params operations.GetGoogleProfileParams) middleware.Responder {
	accessToken := params.AccessToken
	response := operations.NewGetGoogleProfileOK()
	internalErr := operations.NewGetImagesInternalServerError()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://www.googleapis.com/oauth2/v1/userinfo", nil)
	if err != nil {
		log.Printf("Failed to create new request, %v", err)
		return internalErr
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to get tradesperson google info, %v", err)
		return internalErr
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read body in %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: %v", string(body))
		return internalErr
	}

	var results map[string]interface{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Printf("Failed to unmarshal response, %v", err)
		return internalErr
	}

	email := results["email"].(string)
	picture := results["picture"].(string)

	credit, err := database.SaveUser(email)
	if err != nil {
		log.Printf("Failed to save user, %v", err)
		return internalErr
	}

	profile := models.Profile{
		Email: email,
		Image: picture,
	}
	response.SetPayload(&operations.GetGoogleProfileOKBody{
		Profile:     &profile,
		ImageCredit: int64(credit),
	})
	return response
}
