package main

import "fmt"

type Account struct {
	balance  int
	firtName string
	lastName string
}

func (a1 *Account) withdrawPointer(amount int) { a1.balance -= amount }
func (a2 Account) withdrawValue(amount int)    { a2.balance -= amount }
func (a3 Account) withdrawReturnValue(amount int) Account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *Account = &Account{100, "Joe", "Park"}
	fmt.Println("mainA start:", mainA.balance)

	mainA.withdrawPointer(30)
	fmt.Println("mainA with pointer:", mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println("mainA with value:", mainA.balance)

	var mainB Account = mainA.withdrawReturnValue(20)
	fmt.Println("mainB start:", mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println("mainB with pointer:", mainB.balance)
}
