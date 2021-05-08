1 | 다음에서 설명하는 함수 정의를 작성하세요.
:--:|:--
조건 | <ul><li>함수 이름은 Multiple 입니다.</li><li>입력으로 int 타입 2개를 받고, int 타입값 1개를 반환합니다.</li><li>두 입력값을 곱한 결괏값을 반환합니다.</li></ul>

<details>
<summary> Answer </summary>

```go
func Multiple(a, b int) int {
    return a * b
}
```

</details>

---

2 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func AAA() {
    fmt.Println("start AAA()")
    BBB()
    fmt.Println("end AAA()")
}

func BBB() {
    fmt.Println("BBB()")
}

func main() {
    AAA()
}
```

<details>
<summary> Answer </summary>

```sh
start AAA()
BBB()
end AAA()
```

</details>

---

3 | 다음은 피보나치 수열의 정의와 재귀 호출을 이용해서 구현한 예제입니다. 주석 부분에 탈출 조건을 명시해서 함수를 완성하세요.
:--:|:--
조건 | <ul><li>N번째 피보나치 수열 값을 F(N)이라고 합니다.</li><li>F(0) 값은 0입니다.</li><li>F(1) 값은 1입니다.</li><li>F(N) 값은 F(N-2) + F(N-1)입니다.</li></ul>

```go
package main

import "fmt"

func F(n int) int {
    // 여기에 탈출 조건을 채우세요.
    return F(n-2) + F(n-1)
}

func main() {
    // 피보나치 수열 9번째 값을 출력합니다.
    fmt.Println(F(9))
}
```

<details>
<summary> Answer </summary>

```go
// q07.03.go
package main

import "fmt"

func F(n int) int {
    if n < 2 {
        return n
    }
    return F(n-2) + F(n-1)
}

func main() {
    fmt.Println(F(9))
}
```

[q07.03.go](./q07.03/q07.03.go)

</details>
