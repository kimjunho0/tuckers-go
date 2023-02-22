package main

import "fmt"

func addNum(slice []int) []int {
	slice = append(slice, 4) //포인터로 넘기는 방법도 있음
	fmt.Println("addNum slice :", slice)
	return slice
}

func main() {
	slice := []int{1, 2, 3}
	slice = addNum(slice)

	fmt.Println("main slice :", slice)
}
