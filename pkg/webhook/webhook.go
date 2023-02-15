package webhook

import (
	"log"

	"github.com/Yard-Team/whatsapp-api-client-golang/pkg/api"
)

type BasisAPIWebhook struct {
	BasisAPI api.BasisAPI
}

var running = true

func (w BasisAPIWebhook) Start(handler func(map[string]interface{})) {
	for running {
		response, err := w.BasisAPI.Methods().Receiving().ReceiveNotification()
		if err != nil {
			log.Fatal(err)
		}

		if response == nil {
			continue
		} else {
			body := response["body"]
			handler(body.(map[string]interface{}))

			receiptId := int(response["receiptId"].(float64))
			response, err = w.BasisAPI.Methods().Receiving().DeleteNotification(receiptId)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (w BasisAPIWebhook) Stop() {
	running = false
}
