// 배열 복사
package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100} // 크기 다르면 안됌

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
