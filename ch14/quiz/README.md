1 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func add(p1, p2, p3 *int) {
    *p = *p1 + *p2
}

func main() {
    a := 3
    b := 5
    c := 0

    add(&a, &b, &c)
    fmt.Println(c)
}
```

<details>
<summary> Answer </summary>

```sh
8
```

</details>

---

2 | `Actor` 구조체를 생성하고 주어진 인수를 이용해서 값을 초기화해서 반환하는 `NewActor()` 함수를 완성하세요.
:--:|:--

```go
package main

import "fmt"

type Actor struct {
    Name  string
    HP    int
    Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
    // 여기를 채우세요.
}

func main() {
    var actor = NewActor("금토끼", 99, 100)

    fmt.Println(actor.Speed)
    fmt.Println(actor.Name)
}
```

<details>
<summary> Answer </summary>

```go
// q14.02.go
package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	return &Actor{name, hp, speed}
}

func main() {
	var actor = NewActor("금토끼", 99, 100)

	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}
```

[q14.02.go](./q14.02/q14.02.go)

</details>

---

3 | 다음 예제에서 총 몇 개의 `User` 인스턴스가 존재하는지 쓰세요.
:--:|:--

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func NewUser(name string, age int) *User {
    var u = User{name, age}
    return &u
}

func main() {
    newUser := NewUser("AAA", 23)
    var p *User = NewUser
    p.Age += 10

    fmt.Println(p.Age)
}
```

<details>
<summary> Answer </summary>

```sh
1
```

</details>
