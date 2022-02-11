package main

import "fmt"

func square(x int) int {
	return x * x
}

func main() {
	input := 9

	fmt.Printf("%d * %d = %d\n", input, input, square(input))
}
