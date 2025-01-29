package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

type razerPay struct{}

func (r razerPay) pay(amount float32) {
	fmt.Println("Paying using RazerPay", amount)
}
type strapi struct{

}
func (s strapi) pay(amount float32) {
	fmt.Println("Paying using Strapi", amount)
}

func inf() {
    rPay := razerPay{}
		sPay := strapi{}
		newPayGw := payment{
			gateway: rPay,
		}
		newPayGw.gateway.pay(42.27)

}