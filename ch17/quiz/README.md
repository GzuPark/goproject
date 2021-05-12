1 | 프로그램이 실행될 때마다 다른 랜덤값이 생성되도록 다음 공란을 채우세요.
:--:|:--

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.________(time.Now().________())

    n := rand.Intn(100)
    fmt.Println(n)
}
```

<details>
<summary> Answer </summary>

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    n := rand.Intn(100)
    fmt.Println(n)
}
```

</details>

---

2 | 간단한 슬롯머신 게임을 만들어보겠습니다.
:--:|:--
규칙1 | 가진 돈은 1000원으로 시작합니다.
규칙2 | 1 ~ 5 사이의 값을 입력받습니다. 그런 뒤 1 ~ 5 사이 랜덤한 값을 선택합니다.
규칙3 | 만약 입력한 값과 랜덤한 값이 같으면 가진 돈에 500원을 추가하고 축하한다는 메시지와 잔액을 표시합니다.
규칙4 | 다를 경우 가진 돈에서 100원을 빼고 아쉽다는 메시지와 잔액을 표시합니다.
규칙5 | 다시 1 ~ 5 사이의 값을 입력받는 것을 반복하다가 가진 돈이 0원 이하가 되거나 5000원 이상이 되면 게임을 종료합니다.

<details>
<summary> Answer </summary>

```go
// q17.02.go
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
```

[q17.02.go](./q17.02/q17.02.go)

</details>
