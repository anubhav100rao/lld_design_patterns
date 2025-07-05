// services/payment/main.go
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

	nc.Subscribe("stock.reserved", func(msg *nats.Msg) {
		var evt events.StockReserved
		_ = json.Unmarshal(msg.Data, &evt)
		log.Printf("Payment: charging for order %s", evt.OrderID)

		// *** charge the customer ***
		charged := true // pretend outcome

		var out []byte
		if charged {
			resp := events.PaymentCompleted{OrderID: evt.OrderID, Amount: 123.45}
			out, _ = json.Marshal(resp)
			nc.Publish("payment.completed", out)
			log.Println("Payment: completed", evt.OrderID)
		} else {
			resp := events.PaymentFailed{OrderID: evt.OrderID, Reason: "card declined"}
			out, _ = json.Marshal(resp)
			nc.Publish("payment.failed", out)
			log.Println("Payment: failed", evt.OrderID)
		}
	})

	select {}
}
