package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	//var slice []int = []int{1, 2, 3}
	//slice := make([]int, 3)
	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}
	slice[1] = 10
	fmt.Println(slice)
}
