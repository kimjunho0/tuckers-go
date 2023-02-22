package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		r := rand.Intn(101)
		if r >= 0 && r <= 20 {
			fmt.Println("성공!")
			c++
		} else {
			fmt.Println("실패!")
		}
	}
	fmt.Printf("성공 횟수 : %d회", c)
}
