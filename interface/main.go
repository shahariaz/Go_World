package main

import "fmt"

type payment struct {
}

func (p payment) makePayment(amount float32)  {
	razerPaymentGw := razerPay{}
	razerPaymentGw.pay(amount)
}

type razerPay struct{}

func (r razerPay) pay(amount float32) {
	fmt.Println("Paying using RazerPay",amount)
}
func main() {
	p  := payment{}
	p.makePayment(42.27)
} 