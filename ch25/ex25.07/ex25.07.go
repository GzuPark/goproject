package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Number struct{}

func customContext(ctx context.Context) {
	if v := ctx.Value(Number{}); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d\n", n*n)
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), Number{}, 9)
	go customContext(ctx)

	wg.Wait()
}
