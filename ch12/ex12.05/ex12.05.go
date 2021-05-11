package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\t%p\n", i, v, &a[i])
	}
	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\t%p\n", i, v, &b[i])
	}
	fmt.Println()

	b = a

	for i, v := range b {
		fmt.Printf("b[%d] = %d\t%p\n", i, v, &b[i])
	}
}
