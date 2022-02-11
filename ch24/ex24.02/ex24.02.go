package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(20)
	for i := 1; i <= 20; i++ {
		go SumAtoB(1, 100000000*i)
	}

	wg.Wait()
	fmt.Println("Complete")
}
