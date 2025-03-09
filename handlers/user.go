package handlers

import (
	"log"
	"strings"
	"th3-sh0p-api/database"
	"th3-sh0p-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"google.golang.org/api/idtoken"
)

func GetUserCredit(params operations.GetUserCreditParams, principal interface{}) middleware.Responder {
	response := operations.NewGetUserCreditOK()
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

	credit, err := database.UserCredit(email)
	if err != nil {
		log.Printf("Failed to parse payload, %v", err)
		return internalErr
	}
	response.SetPayload(&operations.GetUserCreditOKBody{
		ImageCredit: credit,
	})
	return response
}
