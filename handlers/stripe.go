package handlers

import (
	"th3-sh0p-api/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetPubKey(params operations.GetPubKeyParams) middleware.Responder {
	return operations.NewGetPubKeyOK().WithPayload("pk_test_51R0E3oRfakCijU1OfYANR05o9c37Vu5rIu9eI8pTpib0HqtJ21XxxItdSk4alLSdsN0ihg8PZ5v9Bm21nzH7DsuH00TsXibysm")
}
