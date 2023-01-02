package internal

import (
	"context"
	"encoding/json"
	"github.com/FuzzyStatic/blizzard/v3"
	"github.com/disgoorg/disgo/discord"
	"github.com/spf13/viper"
	"log"
	"time"
	"wowtools/pkg/utilities"
)

func TokenPriceAlert(blizzAuthClient *blizzard.Client, discordWebhook string) {
	tokenPrice, _, err := blizzAuthClient.WoWToken(context.Background())
	if err != nil {
		log.Fatal("Error accessing token price via API: ", err)
	}
	tokenResp, err := json.MarshalIndent(tokenPrice, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	type tokenData struct {
		Price int `json:"price"`
	}
	var response tokenData
	jsonErr := json.Unmarshal(tokenResp, &response)
	if jsonErr != nil {
		log.Fatal("Error unmarshalling json: ", jsonErr)
	}
	if response.Price == 0 {
		log.Fatalln("Error getting token data from Blizzard API")
	}
	if response.Price > viper.GetInt("wow_token_min") {
		image := discord.EmbedResource{
			URL:      viper.GetString("discord_embed_image"),
			ProxyURL: "",
			Height:   0,
			Width:    0,
		}
		discordClient := utilities.DiscordPost(discordWebhook)
		currentTime := time.Now()
		embedContent := discord.Embed{
			Title:       "wowtools.io",
			Type:        "",
			Description: "The WoW Token is above 200k",
			URL:         "",
			Timestamp:   &currentTime,
			Color:       0,
			Footer:      nil,
			Image:       &image,
			Thumbnail:   nil,
			Video:       nil,
			Provider:    nil,
			Author:      nil,
			Fields:      nil,
		}

		utilities.MessageSend(discordClient, embedContent)
		println(tokenPrice)
	}
}
