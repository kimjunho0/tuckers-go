package main

import "fmt"

func main() {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}

	var array = [...]int{1, 2, 3} // 배열 선언
	var slice = []int{1, 2, 3}    //	 슬라이스 선언

	fmt.Println(slice1, slice2, slice, array)

}
