package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	fmt.Printf("a: %d, b: %d\n", a, b)

	a, b = b, a

	fmt.Printf("a: %d, b: %d", a, b)
}
