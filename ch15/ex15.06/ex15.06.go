package main

import "fmt"

func main() {
	str := "Hello 월드"
	runes := []rune(str)

	fmt.Printf("str = %s\tlen(str) = %d\n", str, len(str))
	fmt.Printf("runes = %s\tlen(runes) = %d\n", string(runes), len(runes))
}
