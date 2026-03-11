package strategy

import "fmt"

type CreditCardPayment struct {
	cardNumber string
	name       string
	cvv        string
	dateOfExp  string
}

func NewCreditCardPayment(cardNumber, name, cvv, dateOfExp string) *CreditCardPayment {

	return &CreditCardPayment{
		cardNumber: cardNumber,
		name:       name,
		cvv:        cvv,
		dateOfExp:  dateOfExp,
	}
}

func (p *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("%.2f paid using Credit Card", amount)
}
