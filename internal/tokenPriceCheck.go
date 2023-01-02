package internal

import (
	"github.com/disgoorg/disgo/discord"
	"time"
	"wowtools/pkg/utilities"
)

func TokenPriceAlert(accessToken string, discordWebhook string) {

	tokenPrice := utilities.GetTokenPrice(accessToken)
	//if tokenPrice > 2000000000 {
	// Discord Logic Here
	// }

	image := discord.EmbedResource{
		URL:      "https://media.tenor.com/f42nTKr9aggAAAAC/excited-ronpaul.gif",
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
