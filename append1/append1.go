package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 5)
	slice1 = append(slice1, 1, 2, 3)
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice1 [1] = 100")
	slice1[1] = 100
	fmt.Println("append slice1,500")
	slice1 = append(slice1, 500)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
}
