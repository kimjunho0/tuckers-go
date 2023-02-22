package main

import "fmt"

func main() {
	var price, friends int
	var richhave bool
	fmt.Println("가격은? : ")
	fmt.Scanln(&price)

	if price > 50000 {
		fmt.Println("부자가 있는가?")
		fmt.Scanln(&richhave)
		if richhave == true {
			fmt.Println("신발끈을 묶는다.")
		} else {
			fmt.Println("돈을 나눠 낸다")
		}
	} else if price >= 30000 {
		fmt.Println("친구 수는?")
		fmt.Scanln(&friends)
		if friends > 3 {
			fmt.Println("신발끈을 묶는다")
		} else {
			fmt.Println("나눠 낸다")
		}
	} else {
		fmt.Println("내가 낸다")
	}
}
