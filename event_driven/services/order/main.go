// services/order/main.go
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

	// Simulate receiving an HTTP request to create an order:
	newOrder := events.OrderCreated{
		OrderID:    uuid.NewString(),
		CustomerID: "cust‑1234",
		Items: []struct {
			SKU string
			Qty int
		}{{"widget-x", 2}, {"gadget-y", 1}},
	}

	data, _ := json.Marshal(newOrder)
	if err := nc.Publish("orders.created", data); err != nil {
		log.Fatal("publish error:", err)
	}
	log.Println("Published OrderCreated:", newOrder.OrderID)
}
