package main

import "fmt"

type payment struct {

	geteway stripe
	
}

func (p payment ) makepayment( amount float64) {
	// paymentGatway := razorpay{}
	// paymentGatway.pay(amount)

	// paymentGatway := stripe{}
	// paymentGatway.pay(amount)

	p.geteway.pay(amount)
	
}

type razorpay struct {}
func (r razorpay) pay(amount float64) {
	fmt.Println("Payment made of razorpay", amount)
}

type stripe struct {}
func (s stripe) pay(amount float64) {
	fmt.Println("Payment made of stripe", amount)
}

func main() {
	// payment := payment{}
	// payment.makepayment(100)

	payment := payment{
		geteway: stripe{},
	}
	payment.makepayment(100)

	
}
