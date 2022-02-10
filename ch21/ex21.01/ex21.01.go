package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	fmt.Printf("type of nums: %T\n", nums)

	for _, n := range nums {
		sum += n
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
