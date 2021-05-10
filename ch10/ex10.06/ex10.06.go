package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	age := 0

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Printf("My age is %d\t%p\n", age, &age)
	}

	fmt.Printf("age is %d\t%p\n", age, &age)
}
