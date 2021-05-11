package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A int8
	G int8
	D uint16
	F float32
	B int
	C float64
	E int
}

func main() {
	var padding Padding
	fmt.Println(unsafe.Sizeof(padding))
}
