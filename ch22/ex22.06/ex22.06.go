package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[461] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete
	delete(m, 78)
	v, ok := m[78]
	if ok {
		fmt.Println("Value:", v)
	} else {
		fmt.Println("Do not exist")
	}

	v, ok = m[100]
	if ok {
		fmt.Println("Value:", v)
	} else {
		fmt.Println("Do not exist")
	}
}
