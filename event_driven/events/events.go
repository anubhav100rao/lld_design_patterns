// events/events.go
package events

type OrderCreated struct {
	OrderID    string
	CustomerID string
	Items      []struct {
		SKU string
		Qty int
	}
}

type StockReserved struct {
	OrderID string
}

type StockFailed struct {
	OrderID, Reason string
}

type PaymentCompleted struct {
	OrderID string
	Amount  float64
}

type PaymentFailed struct {
	OrderID, Reason string
}

type ShipmentCreated struct {
	OrderID    string
	TrackingID string
}
