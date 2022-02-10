package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3}
	var slice2 = make([]int, 3)

	for i := 11; i <= 15; i++ {
		slice1 = append(slice1, i)
	}

	slice2 = append(slice2, 21, 22, 23, 24, 25)

	fmt.Println(slice1)
	fmt.Println(slice2)
}
