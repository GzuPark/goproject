package main

import "fmt"

const Y int = 3

func main() {
	x := 5
	a := [x]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := [Y]int{1, 2, 3}
	fmt.Println(b)

	var c [10]int

	fmt.Println(c)
}
