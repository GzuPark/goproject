package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct{}

func (f *File) Read() {}

func ReadFile(reader Reader) {
	if c, ok := reader.(Closer); ok {
		c.Close()
		fmt.Println("Success")
	} else {
		fmt.Println("Fail")
	}
}

func main() {
	file := &File{}
	ReadFile(file)
}
