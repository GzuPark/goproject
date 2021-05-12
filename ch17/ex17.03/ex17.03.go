package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

// return radom int from 1 to 100
func GetRandomIntValue() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	return n + 1
}

func main() {
	target := GetRandomIntValue()
	cnt := 0

	for {
		fmt.Printf("숫자값을 입력하세요: ")
		cnt += 1

		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if n > target {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < target {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Printf("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: %d\n", cnt)
				break
			}
		}
	}
}
