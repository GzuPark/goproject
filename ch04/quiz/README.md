1 | 다음 a, b, c, d, e 변수의 타입을 적어보세요.
:--:|:--

```go
a := 3
var b = 3.1415
c := "hello world"
d := int32(10)
var e float32 = 3.1415
```

<details>
<summary> Answer </summary>

- `a`: int64 (if 64bit OS)
- `b`: float64 (if 64bit OS)
- `c`: string
- `d`: int32
- `e`: float32

</details>

2 | 다음 프로그램의 출력 결과가 `360`이 아닌 `104`인 이유를 설명하세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    var a int32 = 360
    var b int8 = int8(a)

    fmt.Println(b)
}

// 104
```

<details>
<summary> Answer </summary>

`b`는 int8 타입으로 -128 ~ 127 까지 값을 표현할 수 있습니다. 360은 int8의 범위를 벗어나기 때문에 마지막 1바이트 값만 남고 나머지가 사라져서 올바르지 않은 값이 나옵니다.

</details>

3 | 다음 프로그램에서 `f1`과 `f2`가 서로 다른 값을 갖는 이유를 설명하세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    var f1 float32 = 123.546789 * 345.678
    var f2 float32 = float32(123.546789) * 345.678

    fmt.Println(f1)
    fmt.Println(f2)
}

// 42707.406
// 42707.41
```

<details>
<summary> Answer </summary>

`f1`은 123.546789 * 345.678 = 42707.406927942가 float32로 변환되는 과정에서 자릿수 제한으로 값이 42707.406으로 변환됐습니다.
`f2`는 123.546789가 먼저 float32로 변환되면서 자릿수 탈락이 되어서 123.54679 가 됐고 그 값에 345.678을 곱한 결과는 52707.40727362 이지만, float32로 표현 가능한 근사값으로 변환되면서 자릿수 탈락이 발생하여 42707.41이 됐습니다.

</details>
