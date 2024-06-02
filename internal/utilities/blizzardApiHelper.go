package utilities

import (
	"context"
	"log"
	"net/http"

	"github.com/FuzzyStatic/blizzard/v3"
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
		Log.WithError(err).Error("Error creating Blizzard client")
	}

	tokenErr := usBlizzClient.AccessTokenRequest(context.Background())
	if tokenErr != nil {
		log.Fatal("Error creating access token request: ", tokenErr)
		Log.WithError(tokenErr).Error("Error creating access token request")
	}
	return usBlizzClient
}
