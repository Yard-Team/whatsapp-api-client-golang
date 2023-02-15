package main

import (
	"fmt"
	"log"

	//"os"

	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/api"
)

func main() {
	//You can set environment variables in your OS
	//
	//IDInstance := os.Getenv("ID_INSTANCE")
	//APITokenInstance := os.Getenv("API_TOKEN_INSTANCE")

	BasisAPI := api.BasisAPI{
		IDInstance:       "9901739550",
		APITokenInstance: "691f0b48a6e344d699cd986d61ea16a288c6538d584b494c8c",
	}

	response, err := BasisAPI.Methods().Sending().SendMessage(map[string]interface{}{
		"chatId":  "11001234567@c.us",
		"message": "Any message",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
