// services/inventory/main.go
package main

import (
	"encoding/json"
	"log"

	"github.com/anubhav100rao/lld_design_patterns/event_driven/events"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nc.Subscribe("orders.created", func(msg *nats.Msg) {
		var evt events.OrderCreated
		_ = json.Unmarshal(msg.Data, &evt)
		log.Printf("Inventory: reserving stock for order %s", evt.OrderID)

		// *** business logic to reserve stock ***
		success := true // pretend outcome

		var out []byte
		if success {
			resp := events.StockReserved{OrderID: evt.OrderID}
			out, _ = json.Marshal(resp)
			nc.Publish("stock.reserved", out)
			log.Println("Inventory: stock reserved", evt.OrderID)
		} else {
			resp := events.StockFailed{OrderID: evt.OrderID, Reason: "out of stock"}
			out, _ = json.Marshal(resp)
			nc.Publish("stock.failed", out)
			log.Println("Inventory: stock failed", evt.OrderID)
		}
	})

	select {} // block forever
}
