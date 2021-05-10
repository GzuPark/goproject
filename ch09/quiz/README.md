1 | 다음 공란에 들어갈 키워드를 적으세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    age := 22

    ________ age < 10 {
        fmt.Println("You are a child")
    } ________ age >= 20 && age < 30 {
        fmt.Println("Best age of your life")
    } ________ {
        fmt.Println("You are beautiful")
    }
}
```

<details>
<summary> Answer </summary>

```go
package main

import "fmt"

func main() {
    age := 22

    if age < 10 {
        fmt.Println("You are a child")
    } else if age >= 20 && age < 30 {
        fmt.Println("Best age of your life")
    } else {
        fmt.Println("You are beautiful")
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
    score := 78
    count := 20

    if count < 10 {
        fmt.Println("평가 수가 모자랍니다.")
    } else if count < 20 {
        if score > 80 {
            fmt.Println("긍정적인 평가")
        } else {
            fmt.Println("판단할 단계가 아닙니다.")
        }
    } else {
        if score > 90 {
            fmt.Println("좋은 평가입니다.")
        } else if score > 80 {
            fmt.Println("살만한 물건입니다.")
        } else if score > 70 {
            fmt.Println("신중히 생각해보세요.")
        } else {
            fmt.Println("좋은 물건이 아닙니다.")
        }
    }
}
```

<details>
<summary> Answer </summary>

```sh
신중히 생각해보세요.
```

</details>

---

3 | 다음 예시를 if 문을 사용해서 완성하고 출력 결과를 적어보세요. temp 변수는 낮 최고 기온을 나타내고 rain 변수는 강수확률을 나타냅니다. 출력은 fmt 패키지의 Println을 사용합니다.
:--:|:--
조건 | <ul><li>일기예보에서 낮 최고 기온이 25도 이상이고 강수확률이 80% 이상이면 "덥고 비가 옵니다." 출력, 낮 최고 기온이 25도 이상이고 강수확률이 20% 이상이면 "덥고 습합니다", 낮 최고 기온이 25도 이상이고 강수확률이 20% 미만이면 "야외 활동하기 좋습니다."를 출력하세요.</li><li>기온이 25도 이상이 아니고 기온이 10도 미만이거나 강수확률이 80% 이상이면 "야외 활동하기 좋지 않습니다."를 출력하세요. 그렇지 않다면 "좋은 날씨입니다."를 출력하세요.</li></ul>

```go
package main

import "fmt"

func main() {
    temp := 30
    rain := 40

    // 여기를 채워보세요.
}
```

<details>
<summary> Answer </summary>

```go
// q09.03.go
package main

import "fmt"

func main() {
	temp := 30
	rain := 40

	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("덥고 비가 옵니다.")
		} else if rain >= 20 {
			fmt.Println("덥고 습합니다.")
		} else {
			fmt.Println("야외 활동하기 좋습니다.")
		}
	} else if temp < 10 || rain >= 80 {
		fmt.Println("야외 활동하기 좋지 않습니다.")
	} else {
		fmt.Println("좋은 날씨입니다.")
	}
}
```

```sh
덥고 습합니다.
```

[q09.03.go](./q09.03/q09.03.go)

</details>
