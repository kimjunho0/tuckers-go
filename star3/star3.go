package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		if i < 5 {
			for j := 0; j < 7; j++ {
				if j > 4-i || j < 4+i {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			}
			fmt.Println()
		} else {
			for ; i > 0; i-- {

			}
		}
	}
}
