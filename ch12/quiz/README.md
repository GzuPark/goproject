1 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    a := [5]int{1, 2, 3, 4, 5}

    for i, v := range a {
        a[i] = v*2
    }

    fmt.Println(a[2])
}
```

<details>
<summary> Answer </summary>

```sh
6
```

</details>

---

2 | 다음 배열의 크기를 쓰세요.
:--:|:--

```go
[3][2][5]float64
```

<details>
<summary> Answer </summary>

```sh
3 * 2 * 5 * 8 byte = 240 byte
```

</details>

---

3 | 다음 예제의 결과를 쓰세요.
:--:|:--

```go
package main

import "fmt"

func ChangeArray(arr [5]int) {
    arr[3] = 3000
}

func main() {
    a := [5]int{1, 2, 3, 4, 5}

    ChangeArray(a)

    fmt.Println(a[3])
}
```

<details>
<summary> Answer </summary>

```sh
4
```

</details>
