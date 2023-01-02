package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"wowtools/internal"
	"wowtools/pkg/utilities"
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
		log.Fatal("Error unmarshalling viper config: ", unmarshallErr)
	}

	blizzAuth := utilities.GetBlizzApiAuth(c.ClientID, c.ClientSecret)
	internal.TokenPriceAlert(blizzAuth, c.WebhookUri)
}
