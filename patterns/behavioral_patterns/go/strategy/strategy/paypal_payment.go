package strategy

import "fmt"

type PayPalPayment struct {
	email string
}

func NewPaypalPayment(email string) *PayPalPayment {

	return &PayPalPayment{
		email: email,
	}
}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("%.2f paid using Paypal", amount)
}
