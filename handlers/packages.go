package handlers

import (
	"log"
	"strings"
	"th3-sh0p-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"google.golang.org/api/idtoken"
)

func PostImagePack(params operations.PostImagePackParams, principal interface{}) middleware.Responder {
	imagePack := *params.Body.ImagePack
	response := operations.NewPostImagePackOK()
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

	imagePackCost := 0
	if imagePack == "pack_1" {
		imagePackCost = 100
	} else if imagePack == "pack_2" {
		imagePackCost = 300
	} else if imagePack == "pack_3" {
		imagePackCost = 500
	}

	piparams := &stripe.PaymentIntentParams{
		Amount:       stripe.Int64(int64(imagePackCost)),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		ReceiptEmail: stripe.String(email),
	}
	pi, err := paymentintent.New(piparams)
	if err != nil {
		log.Printf("Failed to create payment intent, %v", err)
		return internalErr
	}
	response.SetPayload(&operations.PostImagePackOKBody{
		PaymentIntent: pi.ClientSecret,
	})

	return response
}
