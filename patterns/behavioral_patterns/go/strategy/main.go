package main

import "strategy_method/strategy"

func main() {

	cart := strategy.NewShoppingCart(100.0)

	creditCard := strategy.NewCreditCardPayment("1234-5678-9012-3456", "John Doe", "123", "12/25")
	cart.SetPaymentStrategy(creditCard)
	cart.Checkout()

	// Use PayPal payment
	paypal := strategy.NewPaypalPayment("john@example.com")
	cart.SetPaymentStrategy(paypal)
	cart.Checkout()
}
