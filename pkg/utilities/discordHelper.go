package utilities

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"log"
)

func DiscordPost(webhookUri string) webhook.Client {

	client, err := webhook.NewWithURL(webhookUri)
	if err != nil {
		log.Fatal("Unable to create client: ", err)
	}
	return client
}

func MessageSend(client webhook.Client, embed discord.Embed) {
	if _, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().SetEmbeds(embed).Build()); err != nil {
		log.Fatal(err)
	}
}
