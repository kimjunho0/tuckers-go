package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := append([]int{}, slice1...)

	/*slice2 := make([]int, len(slice1))
	for i, v := range slice1 {
		slice2[i] = v
	}*/

	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
}
