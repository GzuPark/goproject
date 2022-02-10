package main

import "fmt"

func printArrayAndSlice(array [5]int, slice []int) {
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]
	printArrayAndSlice(array, slice)

	//
	array[1] = 100
	fmt.Println("After change second element")
	printArrayAndSlice(array, slice)

	//
	slice = append(slice, 500)
	fmt.Println("After append the element")
	printArrayAndSlice(array, slice)

	//
	fmt.Println()
	array2 := [5]int{1, 2, 3, 4, 5}
	slice2 := array2[1:4]
	slice3 := slice2[1:2]
	slice4 := array2[:]
	slice5 := array2[1:4:4]
	fmt.Println("array2:", array2)
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
	fmt.Println("slice5:", slice5, len(slice5), cap(slice5))
}
