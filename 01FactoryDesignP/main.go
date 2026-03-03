package main

import "fmt"

//interface
type Payment interface {
	Pay(amount float64)
}

// implementations
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) {
	fmt.Println("Payment done using Credit card:::", amount)
}

type UPI struct{}

func (u *UPI) Pay(amount float64) {
	fmt.Println("Payment done using UPI:::", amount)
}

// create factory method
func GetPaymentMethod(method string) Payment {

	switch method {
	case "creditcard":
		return &CreditCard{}
	case "upi":
		return &UPI{}
	default:
		return nil
	}
}

func main() {

	payment := GetPaymentMethod("upi")

	payment.Pay(599)

}
