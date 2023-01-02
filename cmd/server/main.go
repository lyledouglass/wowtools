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
		AccessToken string `mapstructure:"blizzard_access_token"`
		WebhookUri  string `mapstructure:"discord_webhook_url"`
	}
	var c config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("Error unmarshalling viper config: ", err)
	}
	accessToken := c.AccessToken
	discordWebhook := c.WebhookUri

	//if tokenPriceCheck {
	internal.TokenPriceAlert(accessToken, discordWebhook)
	//}

}
