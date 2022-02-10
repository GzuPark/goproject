package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("Before:", slice)

	sort.Ints(slice)
	fmt.Println("After", slice)
}
