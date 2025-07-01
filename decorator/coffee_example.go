package decorator

import "fmt"

// Component interface
type Beverage interface {
	Description() string
	Cost() float64
}

// Concrete Component
type Espresso struct{}

func (e *Espresso) Description() string { return "Espresso" }
func (e *Espresso) Cost() float64       { return 1.50 }

// Decorator base — in Go we usually use embedding
type CondimentDecorator struct {
	beverage Beverage
}

func (cd *CondimentDecorator) Description() string { return cd.beverage.Description() }
func (cd *CondimentDecorator) Cost() float64       { return cd.beverage.Cost() }

// Concrete Decorators
type Milk struct{ CondimentDecorator }

func NewMilk(b Beverage) *Milk {
	return &Milk{CondimentDecorator{beverage: b}}
}

func (m *Milk) Description() string { return m.beverage.Description() + ", Milk" }
func (m *Milk) Cost() float64       { return m.beverage.Cost() + 0.50 }

type Sugar struct{ CondimentDecorator }

func NewSugar(b Beverage) *Sugar {
	return &Sugar{CondimentDecorator{beverage: b}}
}

func (s *Sugar) Description() string { return s.beverage.Description() + ", Sugar" }
func (s *Sugar) Cost() float64       { return s.beverage.Cost() + 0.25 }

func RunCoffeeExample() {
	var order Beverage = &Espresso{}
	fmt.Printf("%s: $%.2f\n", order.Description(), order.Cost())

	order = NewMilk(order)
	fmt.Printf("%s: $%.2f\n", order.Description(), order.Cost())

	order = NewSugar(order)
	fmt.Printf("%s: $%.2f\n", order.Description(), order.Cost())

	order = NewMilk(order) // another milk
	fmt.Printf("%s: $%.2f\n", order.Description(), order.Cost())
}
