package main

import (
	"fmt"
	"log"

	//"os"

	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/api"
	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/webhook"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	BasisAPI := api.BasisAPI{
		URL:              "API_HOST",
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	BasisAPIWebhook := webhook.BasisAPIWebhook{
		BasisAPI: BasisAPI,
	}

	BasisAPIWebhook.Start(func(body map[string]interface{}) {
		typeWebhook := body["typeWebhook"]
		if typeWebhook == "incomingMessageReceived" {
			senderData := body["senderData"]
			chatId := senderData.(map[string]interface{})["chatId"]

			response, err := BasisAPI.Methods().Sending().SendMessage(map[string]interface{}{
				"chatId":  chatId,
				"message": "Any message",
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(response)

			BasisAPIWebhook.Stop()
		}
	})
}
