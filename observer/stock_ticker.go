package observer

import "fmt"

// Observer interface
type Observer interface {
	Update(price float64)
}

// Subject interface
type Subject interface {
	Register(o Observer)
	Unregister(o Observer)
	Notify()
}

// ConcreteSubject
type Stock struct {
	symbol    string
	price     float64
	observers map[Observer]struct{}
}

func NewStock(symbol string) *Stock {
	return &Stock{
		symbol:    symbol,
		observers: make(map[Observer]struct{}),
	}
}

func (s *Stock) Register(o Observer) {
	s.observers[o] = struct{}{}
}

func (s *Stock) Unregister(o Observer) {
	delete(s.observers, o)
}

func (s *Stock) Notify() {
	for o := range s.observers {
		o.Update(s.price)
	}
}

func (s *Stock) SetPrice(p float64) {
	fmt.Printf("\n%s price changed to %.2f\n", s.symbol, p)
	s.price = p
	s.Notify()
}

// ConcreteObservers
type ConsoleDisplay struct {
	name string
}

func NewConsoleDisplay(name string) *ConsoleDisplay {
	return &ConsoleDisplay{name}
}

func (c *ConsoleDisplay) Update(price float64) {
	fmt.Printf("[%s] Current Price: %.2f\n", c.name, price)
}

type AlertService struct {
	threshold float64
}

func NewAlertService(th float64) *AlertService {
	return &AlertService{th}
}

func (a *AlertService) Update(price float64) {
	if price > a.threshold {
		fmt.Printf("[ALERT] Price %.2f exceeded threshold %.2f!\n", price, a.threshold)
	}
}

// Client
func RunStockTicker() {
	stock := NewStock("GOOG")

	console1 := NewConsoleDisplay("Display#1")
	console2 := NewConsoleDisplay("Display#2")
	alert := NewAlertService(1500.00)

	stock.Register(console1)
	stock.Register(console2)
	stock.Register(alert)

	stock.SetPrice(1490.50)
	stock.SetPrice(1505.75)

	stock.Unregister(console2)
	stock.SetPrice(1510.00)
}
