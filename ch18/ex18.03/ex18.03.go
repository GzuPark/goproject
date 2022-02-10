package main

import "fmt"

func printSlice(slice1 []int, slice2 []int) {
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}

func main() {
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)
	printSlice(slice1, slice2)

	//
	slice1[1] = 100
	fmt.Println("After change second element")
	printSlice(slice1, slice2)

	//
	slice1 = append(slice1, 500)
	fmt.Println("After append element")
	printSlice(slice1, slice2)

	//
	fmt.Println()
	slice3 := []int{1, 2, 3}
	slice4 := append(slice3, 4, 5)
	printSlice(slice3, slice4)

	//
	slice3[1] = 100
	fmt.Println("After change second element")
	printSlice(slice3, slice4)

	//
	slice3 = append(slice3, 500)
	fmt.Println("After append element")
	printSlice(slice3, slice4)
}
