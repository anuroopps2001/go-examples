package main

import "fmt"

type payment interface {
	pay(amount float64) string
}

type UPI struct {
	ID string
}

type CreditCard struct {
	number string
}

type Wallet struct {
	balance float64
}

func main() {
	ID := UPI{ID: "anuroop@ybl"}

	Number := CreditCard{number: "**** 3473"}

	Wallet := &Wallet{balance: 2000}

	checkout(ID, 499)
	checkout(Number, 499)
	checkout(Wallet, 499)
}

func (u UPI) pay(amount float64) string {
	return fmt.Sprintf("Paid ₹%.2f via UPI (ID: %s)", amount, u.ID)
}
func (cc CreditCard) pay(amount float64) string {
	return fmt.Sprintf("Paid ₹%.2f using Credit Card (Number: %s)", amount, cc.number)
}

func (w *Wallet) pay(amount float64) string {
	if w.balance < amount {
		return "Wallet payment failed: insufficient balance"
	}
	w.balance = w.balance - amount
	return fmt.Sprintf("Paid ₹%.2f using Wallet (Balance remaining: %.2f)", amount, w.balance)
}

func checkout(p payment, amount float64) {
	fmt.Println("Processing payment...")
	fmt.Println(p.pay(amount))
	fmt.Println()
}
