1 | 숫자를 한 번에 하나씩 출력하는 for문을 이용해서 아래 결과를 출력하는 프로그램을 작성하세요.
:--:|:--

```sh
10 9 8 7 6 5 4 3 2 1
```

<details>
<summary> Answer </summary>

```go
// q11.01.go
package main

import "fmt"

func main() {
	for i := 10; i > 0; i-- {
		fmt.Printf("%d ", i)
	}
}
```

[q11.01.go](./q11.01/q11.01.go)

</details>

---

2 | 다음은 구구단을 출력하는 프로그램입니다. 주석 부분을 채워서 3단부터 6단까지는 출력하지 않도록 바꿔보세요.
:--:|:--

```go
package main

import "fmt"

func main() {
    for i := 2; i < 10; i++ {
        // 여기를 채워보세요

        for j := 1; j < 10; j++ {
            fmt.Println(i, "*", j, "=", i*j)
        }

        fmt.Println()
    }
}
```

<details>
<summary> Answer </summary>

```go
// q11.02.go
package main

import "fmt"

func main() {
	for i := 2; i < 10; i++ {
		if i >= 3 && i <= 6 {
			continue
		}

		for j := 1; j < 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}

		fmt.Println()
	}
}
```

[q11.02.go](./q11.02/q11.02.go)

</details>

---

3 | 다음 출력 결과가 나오도록 9까지 홀수의 제곱값을 출력하는 프로그램을 작성하세요.
:--:|:--
Hint | 후처리는 꼭 ++나 --를 쓰지 않고 다른 연산자를 쓰셔도 됩니다.

```sh
1 * 1 = 1
3 * 3 = 9
5 * 5 = 25
7 * 7 = 49
9 * 9 = 81
```

<details>
<summary> Answer </summary>

```go
// q11.03.go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i += 2 {
		fmt.Printf("%d * %d = %d\n", i, i, i*i)
	}
}
```

[q11.03.go](./q11.03/q11.03.go)

</details>

---

4 | 이중 for문을 사용해서 다음 모양의 별을 출력하는 프로그램을 작성하세요.
:--:|:--

```sh
*****
****
***
**
*
```

<details>
<summary> Answer </summary>

```go
// q11.04.go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Printf("*")
		}

		fmt.Println()
	}
}
```

[q11.04.go](./q11.04/q11.04.go)

</details>
