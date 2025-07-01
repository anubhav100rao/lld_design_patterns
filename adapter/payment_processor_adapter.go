package adapter

import "fmt"

// Target interface
type PaymentProcessor interface {
	Pay(amount float64) error
}

// Adaptee #1: Stripe SDK
type StripeSDK struct{}

func (s *StripeSDK) ChargeCents(cents int) {
	fmt.Printf("Stripe: charged %d cents\n", cents)
}

// Adaptee #2: PayPal SDK
type PayPalSDK struct{}

func (p *PayPalSDK) SendPayment(dollars float64) {
	fmt.Printf("PayPal: sent payment of $%.2f\n", dollars)
}

// Adapter for Stripe
type StripeAdapter struct {
	stripe *StripeSDK
}

func NewStripeAdapter(s *StripeSDK) *StripeAdapter {
	return &StripeAdapter{stripe: s}
}

func (a *StripeAdapter) Pay(amount float64) error {
	cents := int(amount * 100)
	a.stripe.ChargeCents(cents)
	return nil
}

// Adapter for PayPal
type PayPalAdapter struct {
	paypal *PayPalSDK
}

func NewPayPalAdapter(p *PayPalSDK) *PayPalAdapter {
	return &PayPalAdapter{paypal: p}
}

func (a *PayPalAdapter) Pay(amount float64) error {
	a.paypal.SendPayment(amount)
	return nil
}

// Client code
func processPayment(p PaymentProcessor, amt float64) {
	if err := p.Pay(amt); err != nil {
		fmt.Println("Payment failed:", err)
	}
}

func RunPaymentProcessorDemo() {
	stripeSDK := &StripeSDK{}
	paypalSDK := &PayPalSDK{}

	stripeProcessor := NewStripeAdapter(stripeSDK)
	paypalProcessor := NewPayPalAdapter(paypalSDK)

	processPayment(stripeProcessor, 29.99) // uses Stripe under the hood
	processPayment(paypalProcessor, 49.50) // uses PayPal under the hood
}
