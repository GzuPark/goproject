package main

import "fmt"

func main() {
	const C int = 10
	var b int = C * 20

	fmt.Println(b)

	//cannot assign to C (declared const)
	// C = 20

	// cannot take the address of C
	fmt.Println(&C)
}
