package main

import "fmt"

type Account struct {
	balance int
}

func withdrawFunc(a *Account, amount int)  { a.balance -= amount }
func (a *Account) withdrawFunc(amount int) { a.balance -= amount }

func main() {
	a := &Account{100}
	fmt.Println("Start:", a.balance)

	withdrawFunc(a, 30)
	fmt.Println("Function:", a.balance)

	a.withdrawFunc(50)
	fmt.Println("Method", a.balance)
}
