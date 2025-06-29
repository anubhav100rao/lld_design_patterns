package factory

import (
	"fmt"
)

type PaymentProcessor interface {
	Charge(amount float64) error
}

type StripeProcessor struct{}

func (s *StripeProcessor) Charge(amount float64) error {
	fmt.Printf("Charged $%.2f using Stripe\n", amount)
	return nil
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) Charge(amount float64) error {
	fmt.Printf("Charged $%.2f using PayPal\n", amount)
	return nil
}

// factory function
func NewPaymentProcessor(provider string) (PaymentProcessor, error) {
	switch provider {
	case "stripe":
		return &StripeProcessor{}, nil
	case "paypal":
		return &PayPalProcessor{}, nil
	default:
		return nil, fmt.Errorf("unsupported payment provider: %s", provider)
	}
}

func RunPaymentProcessorDemo() {
	processor, err := NewPaymentProcessor("stripe")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = processor.Charge(100.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
