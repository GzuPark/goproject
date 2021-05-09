1 | 다음과 같은 상수를 선언하세요.
:--:|:--
조건 | <ul><li>상수 이른은 Gravity 입니다.</li><li>상숫값은 9.80665 입니다.</li><li>타입 없는 상수로 선언하세요.</li></ul>

<details>
<summary> Answer </summary>

```go
const Gravity = 9.80665
```

</details>

---

2 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

const (
    C1 = iota
    C2
    C3
)

const (
    D1 = iota + 1
    D2
    D3
)

func main() {
    fmt.Println(C3, D3)
}
```

<details>
<summary> Answer </summary>

```sh
2 3
```

</details>
