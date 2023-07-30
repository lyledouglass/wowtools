package utilities

import (
	"log"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
)

func DiscordPost(webhookUri string) webhook.Client {

	client, err := webhook.NewWithURL(webhookUri)
	if err != nil {
		Log.WithError(err).Error("Unable to create discord client")
	}
	return client
}

func MessageSend(client webhook.Client, embed discord.Embed) {
	if _, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().SetEmbeds(embed).Build()); err != nil {
		log.Fatal(err)
		Log.WithError(err).Error("Error sending discord message")
	}
}
