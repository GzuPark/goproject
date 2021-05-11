1 | 다음에서 설명하는 구조체를 정의하세요.
:--:|:--
조건 | <ul><li>구조체 이름은 `Product`입니다.</li><li>`string` 타입 `Name` 필드가 있습니다.</li><li>`int` 타입의 `Price` 필드가 있습니다.</li><li>`float64` 타입의 `ReviewScore` 필드가 있습니다.</li></ul>

<details>
<summary> Answer </summary>

```go
type Product struct {
    Name        string
    Price       int
    ReviewScore float64
}
```

</details>

---

2 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

type Actor struct {
    Name  string
    HP    int
    Speed float64
}

type Monster struct {
    Actor
    Attack int
    Speed  int
}

func main() {
    var monster = Monster{
        Actor{"NPCA", 100, 8.7},
        500,
        200,
    }

    fmt.Println(monster.Speed)
    fmt.Println(monster.Actor.Speed)
}
```

<details>
<summary> Answer </summary>

```sh
200
8.7
```

</details>

---

3 | 다음 구조체의 패딩을 최대한 줄이고 구조체 크기를 적으세요.
:--:|:--

```go
type Padding struct {
    A int8
    B int
    C float64
    D uint16
    E int
    F float32
    G int8
}
```

<details>
<summary> Answer </summary>

```sh
32 byte
```

```go
// q13.03.go
package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A int8
	G int8
	D uint16
	F float32
	B int
	C float64
	E int
}

func main() {
	var padding Padding
	fmt.Println(unsafe.Sizeof(padding))
}
```

[q13.03.go](./q13.03/q13.03.go)

</details>
