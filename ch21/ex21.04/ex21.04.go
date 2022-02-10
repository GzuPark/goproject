package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int { return a + b }
	} else if op == "*" {
		return func(a, b int) int { return a * b }
	} else {
		return nil
	}
}

func main() {
	var addition = getOperator("+")(3, 4)
	var multiply = getOperator("*")(3, 4)

	fmt.Println("Add:", addition)
	fmt.Println("Mul:", multiply)
}
