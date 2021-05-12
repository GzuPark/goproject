1 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    str1 := "학교종이 "
    str2 := "땡땡땡"
    str1 += str2

    fmt.Println(str1)
}
```

<details>
<summary> Answer </summary>

```sh
학교종이 땡땡땡
```

</details>

---

2 | 다음은 소문자를 대문자로 바꾸는 함수입니다. 공란을 채우세요.
:--:|:--

```go
package main

import (
    "fmt"
    "strings"
)

func ToUpper(str string) string {
    var builder strings.Builder

    for _, v := ________ str {
        if v >= 'a' && v <= 'z' {
            builder.WriteRune('A' + (v - 'a'))
        } else {
            builder.WriteRune(v)
        }
    }

    return builder.String()
}

func main() {
    str := "hello World!"

    fmt.Println(ToUpper(str))
}
```

<details>
<summary> Answer </summary>

```go
// q15.02.go
package main

import (
    "fmt"
    "strings"
)

func ToUpper(str string) string {
    var builder strings.Builder

    for _, v := range str {
        if v >= 'a' && v <= 'z' {
            builder.WriteRune('A' + (v - 'a'))
        } else {
            builder.WriteRune(v)
        }
    }

    return builder.String()
}

func main() {
    str := "hello World!"

    fmt.Println(ToUpper(str))
}
```

[q15.02.go](./q15.02/q15.02.go)

</details>

---

3 | 64비트 컴퓨터에서 다음 구조체의 크기를 적어보세요.
:--:|:--

```go
type User struct {
    FirstName string
    LastName  string
    Age       int
}
```

<details>
<summary> Answer </summary>

- 각 `string`은 메모리 주소와 길이를 가지는 구조체이므로 16 byte
- 64비트의 `int`는 `int64`를 의미하므로 8 byte
- 따라서, 16 * 2 + 8 = 40 byte

</details>
