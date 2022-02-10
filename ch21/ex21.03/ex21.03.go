package main

import "fmt"

func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }

type opFunc func(int, int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
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
