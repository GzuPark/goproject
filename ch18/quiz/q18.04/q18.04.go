package main

import (
	"fmt"
	"sort"
)

type Player struct {
	Name string
	Age  int
	Goal int
	Pass float32
}

type Players []Player

func (p Players) Len() int           { return len(p) }
func (p Players) Less(i, j int) bool { return p[i].Goal < p[j].Goal }
func (p Players) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	p := Players{
		{"나통키", 13, 45, 78.4},
		{"오영태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}

	fmt.Println("Before", p)
	sort.Sort(p)
	fmt.Println("After", p)
}
