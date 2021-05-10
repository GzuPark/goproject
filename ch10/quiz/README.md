1 | 다음 공란에 들어갈 키워드를 적으세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    day := 1

    ________ day {
    ________ 1:
        fmt.Println("First day")
    ________ 2:
        fmt.Println("Second day")
    ________:
        fmt.Println("Another dya")
    }
}
```

<details>
<summary> Answer </summary>

```go
package main

import "fmt"

func main() {
    day := 1

    switch day {
    case 1:
        fmt.Println("First day")
    case 2:
        fmt.Println("Second day")
    default:
        fmt.Println("Another dya")
    }
}
```

</details>

---

2 | 다음 예제의 결과를 적어주세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    score := 85
    count := 15

    switch {
    case count < 10:
        fmt.Println("평가 수가 모자랍니다.")
    case count < 20 && score > 80:
        fmt.Println("긍정적인 평가")
    case count < 20:
        fmt.Println("판단할 단계가 아닙니다.")
    case score > 90:
        fmt.Println("좋은 평가입니다.")
    case score > 80:
        fmt.Println("살 만한 물건입니다.")
    case score > 70:
        fmt.Println("신중히 생각해보세요.")
    default:
        fmt.Println("좋은 물건이 아닙니다.")
    }
}
```

<details>
<summary> Answer </summary>

```sh
긍정적인 평가
```

</details>

---

3 | switch문을 사용하여 다음 조건에 만족하도록 함수를 작성하세요.
:--:|:--
조건 | <ul><li>함수명은 GetDirection</li><li>함수 매개변수는 angle float64를 받음</li><li>함수 결과는 Direction 타입 반환</li></ul><ul><li>angle이 315 이상이면 North 반환</li><li>angle이 0 이상 45 보다 작으면 North 반환</li><li>angle이 45 이상 135보다 작으면 East 반환</li><li>angle이 135 이상 225보다 작으면 South 반환</li><li>angle이 225 이상 315보다 작으면 West 반환</li><li>모든 조건이 만족하지 않으면 None 반환</li></ul>

```go
package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToSTring(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

// 이곳에 GetDirection 함수를 만드세요.

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
}
```

<details>
<summary> Answer </summary>

```go
// q10.03.go
package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirection(angle float64) Direction {
	switch {
	case angle >= 315, angle >= 0 && angle < 45:
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle >= 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return West
	default:
		return None
	}
}

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
	fmt.Println(DirectionToString(GetDirection(235.8)))
	fmt.Println(DirectionToString(GetDirection(94.2)))
	fmt.Println(DirectionToString(GetDirection(-30)))
}
```

[q10.03.go](./q10.03/q10.03.go)

</details>
