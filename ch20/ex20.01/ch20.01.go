package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Hello! I'm %s and %d years old.", s.Name, s.Age)
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer = student

	fmt.Printf("%s\n", stringer.String())
}
