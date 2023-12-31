package main

import (
	"flag"
	"wowtools/internal"
	"wowtools/pkg/utilities"

	"github.com/spf13/viper"
)

func main() {
	utilities.LoadConfig(".")

	var (
		tokenPriceCheck bool
	)
	// Flags
	flag.BoolVar(&tokenPriceCheck, "token-price-check", false, "checks the price of the WoW token")

	type config struct {
		AccessToken  string `mapstructure:"blizzard_access_token"`
		ClientID     string `mapstructure:"blizzard_clientid"`
		ClientSecret string `mapstructure:"blizzard_clientsecret"`
		WebhookUri   string `mapstructure:"discord_webhook_url"`
	}
	var c config
	unmarshallErr := viper.Unmarshal(&c)
	if unmarshallErr != nil {
		utilities.Log.WithError(unmarshallErr).Error("Error unmarshalling viper config")
	}

	blizzAuth := utilities.GetBlizzApiAuth(c.ClientID, c.ClientSecret)
	internal.TokenPriceAlert(blizzAuth, c.WebhookUri)
}
