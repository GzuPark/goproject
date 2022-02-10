package main

import "fmt"

func main() {
	// remove at
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", slice1)

	removeIdx := 2
	slice1 = append(slice1[:removeIdx], slice1[removeIdx+1:]...)
	fmt.Println("After:", slice1)

	// insert at
	fmt.Println()
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", slice2)

	insertIdx := 3
	slice2 = append(slice2, 0)
	copy(slice2[insertIdx+1:], slice2[insertIdx:])
	slice2[insertIdx] = 100
	fmt.Println("After:", slice2)
}
