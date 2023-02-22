package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		slice[i] += 10
	}

	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Println(slice)
}
