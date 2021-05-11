package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	C int8
	A int8
	E int8
	B int
	D int
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
