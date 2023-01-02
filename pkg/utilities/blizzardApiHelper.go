package utilities

import (
	"context"
	"github.com/FuzzyStatic/blizzard/v3"
	"log"
	"net/http"
)

// GetBlizzApiAuth Uses the FuzzyStatic Blizzard API wrapper package to create an access token
func GetBlizzApiAuth(clientId string, clientSecret string) *blizzard.Client {

	usBlizzClient, err := blizzard.NewClient(blizzard.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		HTTPClient:   http.DefaultClient,
		Region:       blizzard.US,
		Locale:       blizzard.EnUS,
	})
	if err != nil {
		log.Fatal("Error creating Blizzard client: ", err)
	}

	tokenErr := usBlizzClient.AccessTokenRequest(context.Background())
	if tokenErr != nil {
		log.Fatal("Error creating access token request: ", tokenErr)
	}
	return usBlizzClient
}
