package internal

import (
	"context"
	"log"
	"strings"

	"google.golang.org/api/idtoken"
)

var audience string = "889504103274-5fu49653t4beaotjhnn4ql0fj8i5vpd3.apps.googleusercontent.com"

func ValidateToken(bearerHeader string) (interface{}, error) {
	idToken := strings.Split(bearerHeader, " ")[1]
	_, err := idtoken.Validate(context.Background(), idToken, audience)
	if err != nil {
		log.Printf("%v", err)
		return false, err
	}
	return true, nil
}
