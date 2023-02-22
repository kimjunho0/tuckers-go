package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200 // 40byte 다 복사
}

func changeSlice(slice2 []int) {
	slice2[2] = 200 // 3개의 필드(Data, Len, Cap) 복사 24byte 고정
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array)
	fmt.Println(slice)
}
