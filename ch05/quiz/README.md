1 | 다음 예제 결과를 적어보세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    var a = 345
    var b = 3.1415

    fmt.Printf("%05d\n", a)
    fmt.Printf("%5.2f\n", b)
}
```

<details>
<summary> Answer </summary>

`00345 3.14`

</details>

2 | 다음 예제가 제대로 동작하지 않는 이유를 적으세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    var a int
    var b int

    fmt.Scanln(a, b)
    fmt.Println(a, b)
}
```

<details>
<summary> Answer </summary>

`Scanln()`을 이용해서 값을 입력 받을 경우 변수의 주소에 해당 값을 받도록 되어 있기 때문에 동작하지 않습니다.

</details>


3 | 다음과 같이 출력되도록 `fmt.Printf`와 서식 문자를 이용한 코드를 채우세요. 이 때 출력 결과의 최소 너비를 6으로 지정합니다.
:--:|:--

```go
//    123
// 004567
//   3.14

package main

import "fmt"

func main() {
    var a = 123
    var b int = 4567
    f := 3.14159269

    // a를 이용해 출력하세요.
    // b를 이용해 출력하세요.
    // f를 이용해 출력하세요.
}
```

<details>
<summary> Answer </summary>

```go
// q05.03.go
package main

import "fmt"

func main() {
    var a = 123
    var b int = 4567
    f := 3.14159269

    fmt.Printf("%6d\n", a)
    fmt.Printf("%06d\n", b)
    fmt.Printf("%6.2f\n", f)
}
```

[q05.03.go](./q05.03/q05.03.go)

</details>
