package main

import "fmt"

// PaymentStrategy defines the interface for payment methods.
type PaymentStrategy interface {
    Pay(amount float64)
}
// PayPal represents a payment method using PayPal.
type PayPal struct {
    email string
}

func (p *PayPal) Pay(amount float64) {
    fmt.Printf("Paid %.2f using PayPal account: %s\n", amount, p.email)
}

// CreditCard represents a payment method using a credit card.
type CreditCard struct {
    cardNumber string
}

func (c *CreditCard) Pay(amount float64) {
    fmt.Printf("Paid %.2f using Credit Card: %s\n", amount, c.cardNumber)
}

// BankTransfer represents a payment method using bank transfer.
type BankTransfer struct {
    bankAccount string
}

func (b *BankTransfer) Pay(amount float64) {
    fmt.Printf("Paid %.2f using Bank Transfer: %s\n", amount, b.bankAccount)
}


// PaymentProcessor acts as a context to process payments with a selected strategy.
type PaymentProcessor struct {
    strategy PaymentStrategy
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
    p.strategy = strategy
}

func (p *PaymentProcessor) Pay(amount float64) {
    p.strategy.Pay(amount)
}

func main() {
    processor := &PaymentProcessor{}

    // Use PayPal payment method
    processor.SetStrategy(&PayPal{email: "user@example.com"})
    processor.Pay(100.00)

    // Switch to Credit Card payment method
    processor.SetStrategy(&CreditCard{cardNumber: "1234-5678-9012-3456"})
    processor.Pay(200.00)

    // Switch to Bank Transfer payment method
    processor.SetStrategy(&BankTransfer{bankAccount: "ABCD-1234"})
    processor.Pay(300.00)
}
