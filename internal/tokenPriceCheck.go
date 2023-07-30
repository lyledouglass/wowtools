package internal

import (
	"context"
	"encoding/json"
	"strconv"
	"time"
	"wowtools/pkg/utilities"

	"github.com/FuzzyStatic/blizzard/v3"
	"github.com/disgoorg/disgo/discord"
	"github.com/spf13/viper"
)

func TokenPriceAlert(blizzAuthClient *blizzard.Client, discordWebhook string) {
	minTokenPrice := viper.GetInt("wow_token_min")
	tokenPrice, _, err := blizzAuthClient.WoWToken(context.Background())
	if err != nil {
		utilities.Log.WithError(err).Error("Error accessing token price via API")
	}
	tokenResp, err := json.MarshalIndent(tokenPrice, "", "  ")
	if err != nil {
		utilities.Log.WithError(err).Error("Error running json.MarshalIndent on TokenPrice")
	}

	type tokenData struct {
		Price int `json:"price"`
	}
	var response tokenData
	jsonErr := json.Unmarshal(tokenResp, &response)
	if jsonErr != nil {
		utilities.Log.WithError(jsonErr).Error("Error unmarshalling json")
	}
	if response.Price == 0 {
		utilities.Log.Error("Error getting token data from Blizzard API")
	}
	if response.Price > minTokenPrice {
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
			Description: "The WoW Token is above " + strconv.Itoa(minTokenPrice/100000),
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
		utilities.Log.Info(tokenPrice)
	}
}
