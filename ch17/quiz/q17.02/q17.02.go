package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	InitMoney int = 1000
	GetMoney  int = 500
	PayMoney  int = 100
	Win       int = 5000
	Lose      int = 0
)

var stdin = bufio.NewReader(os.Stdin)

func InputValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

// return random int from 1 to 5
func GetRandomValue() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5)
	return n + 1
}

func IsEndGame(money int) bool {
	if money <= Lose {
		fmt.Println("게임 패배")
		return true
	} else if money >= Win {
		fmt.Println("게임 승리")
	}
	return false
}

func main() {
	var money int = InitMoney
	var cnt int = 1

	fmt.Println("슬롯머신 게임에 참여하신 것을 환영합니다.")
	fmt.Printf("시작할 때 지급되는 돈은 %d 원입니다.\n\n", money)

	for {
		fmt.Printf("%d 회차 게임을 시작합니다. 1 ~ 5 사이의 값을 입력하세요: ", cnt)
		target := GetRandomValue()
		n, err := InputValue()

		if err != nil {
			fmt.Println("1 ~ 5 사이의 숫자를 입력하세요.")
		} else {
			if n == target {
				money += GetMoney
				fmt.Printf("축하합니다. 잔액은 %d원 입니다.\n\n", money)
			} else {
				money -= PayMoney
				fmt.Printf("아쉽군요. 잔액은 %d원 입니다.\n\n", money)
			}

			if IsEndGame(money) {
				break
			}
		}
		cnt++
	}
}
