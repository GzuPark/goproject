1 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    var a int8 = 30

    a <<= 2
    a += 8
    fmt.Println(a)
}
```

<details>
<summary> Answer </summary>

```go
30 -> b 0001 1110
b 0001 1110 << 2 = b 0111 1000
b 0111 1000 + 8 -> b 0111 1000 + b 0000 1000 = b 1000 0000
b 1000 0000 -> -128
```

</details>

---

2 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    var a uint8
    a |= 2
    a |= 4
    a |= 8

    var b uint8
    b = 4

    a &^= b
    fmt.Println(a)
}
```

<details>
<summary> Answer </summary>

```go
a = 0
b 0000 0000 | b 0000 0010 = b 0000 0010
b 0000 0010 | b 0000 0100 = b 0000 0110
b 0000 0110 | b 0000 1000 = b 0000 1110

// &^ 연산은 우측 값에서 1인 비트를 0으로 변경
b 0000 1110 &^ b 0000 0100 = b 0000 1010 = 10
```

</details>

---

3 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    var x int8 = 1
    x <<= 7
    x >>= 7
    fmt.Printf("%d\n", x)
}
```

<details>
<summary> Answer </summary>

```go
b 0000 0001 << 7 = b 1000 0000
b 1000 0000 >> 7 = b 1111 1111 // 첫 bit는 부호를 나타내므로 음수일 경우 시프트는 1을 채우게 됨
b 1111 1111 = -1
```

</details>
