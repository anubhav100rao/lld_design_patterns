// services/shipping/main.go
package main

import (
	"encoding/json"
	"log"

	"github.com/anubhav100rao/lld_design_patterns/event_driven/events"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("payment.completed", func(msg *nats.Msg) {
		var evt events.PaymentCompleted
		_ = json.Unmarshal(msg.Data, &evt)
		tracking := uuid.NewString()
		log.Printf("Shipping: creating shipment %s for order %s", tracking, evt.OrderID)

		resp := events.ShipmentCreated{OrderID: evt.OrderID, TrackingID: tracking}
		data, _ := json.Marshal(resp)
		nc.Publish("shipment.created", data)
		log.Println("Shipping: shipment created", evt.OrderID)
	})

	select {}
}
