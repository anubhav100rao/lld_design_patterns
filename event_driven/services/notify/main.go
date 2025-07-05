// services/notify/main.go
package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	subjects := []string{"orders.created", "stock.reserved", "payment.completed", "shipment.created"}
	for _, subj := range subjects {
		nc.Subscribe(subj, func(msg *nats.Msg) {
			log.Printf("Notification: event %s received: %s", msg.Subject, string(msg.Data))
			// send email/SMS...
		})
	}

	select {}
}
