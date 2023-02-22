package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(101)
	var s, cnt int

	for {
		fmt.Print("숫자 입력")
		fmt.Scanln(&s)

		if s < r {
			fmt.Println("큽니다")
			continue
		} else if s > r {
			fmt.Println("작습니다")
			continue
		} else {

			fmt.Println("축하합니다 시도한 횟수:", cnt)
			break
		}
		cnt++
	}
}
