package handlers

import (
	"os"
	"th3-sh0p-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetPubKey(params operations.GetPubKeyParams) middleware.Responder {
	return operations.NewGetPubKeyOK().WithPayload(os.Getenv("STRIPE_PUB_KEY"))
}
