package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	// 실제 배열의 포인터 주소를 복사하므로 값을 변경하면 같은 포인터 주소를 바라보고 있으므로 변경됨
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}
