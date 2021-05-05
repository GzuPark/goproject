package main

import "fmt"

func main() {
	fmt.Println(3*4^7<<2+3*5 == 7)
	fmt.Println((((3 * 4) ^ (7 << 2)) + (3 * 5)) == 7) // 가독성을 위해 더 좋은 코드
}
