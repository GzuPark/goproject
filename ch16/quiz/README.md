1 | 다음 패키지에서 외부로 공개되는 항목을 모두 쓰세요.
:--:|:--

```go
package screen

var ScreenX = 680
var ScreenY = 720
var width = 1080

const ColorDepth = 16
const velocity = 8

func ResizeScreen(x, y int) {
    ScreenX, ScrrenY = x, y
}

func changeWidth(w int) {
    width = w
}
```

<details>
<summary> Answer </summary>

- variable: ScreenX, ScreenY
- constant: ColorDepth
- function: ResizeScreen()

</details>

---

2 | 다음 패키지를 임포트했을 때 출력 결과를 쓰세요.
:--:|:--

```go
package custom

import "fmt"

var (
    a = b + f()
    b = c
    c = f()
    d = 3
)

func init() {
    fmt.Printf("init function a:%d b:%d c:%d d:%d\n", a, b, c, d)
}

func f() int {
    d++
    fmt.Println("f() d:", d)
    return d
}
```

<details>
<summary> Answer </summary>

```
f() d: 4
f() d: 5
init function a:9 b:4 c:4 d:5
```

</details>
