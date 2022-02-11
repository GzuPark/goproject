package main

import (
	"errors"
	"fmt"
	"strings"
)

func Atoi(input string) (int, error) {
	rst := 0
	negative := false
	input = strings.TrimSpace(input)

	if input[0] == '-' {
		negative = true
		input = input[1:]
	}

	for _, c := range input {
		if c >= '0' && c <= '9' {
			rst *= 10
			rst += int(c - '0')
		} else {
			return 0, errors.New("cannont conver to int")
		}
	}

	if negative {
		rst *= -1
	}

	return rst, nil
}

func main() {
	n, err := Atoi("0")
	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	fmt.Println(n)
}
