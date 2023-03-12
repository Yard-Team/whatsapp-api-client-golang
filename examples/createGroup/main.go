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
		URL:              "API_HOST",
		IDInstance:       "IDInstance",
		APITokenInstance: "APITokenInstance",
	}

	response, err := BasisAPI.Methods().Groups().CreateGroup("groupName", []string{
		"11001234567@c.us",
		"11002345678@c.us",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
