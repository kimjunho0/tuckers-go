package main

import "fmt"

func main() {
	a := 1
	b := 1
	/* OuterFor: //Label 레이블 */
	flag := false
	for b = 1; b <= 9; b++ {
		for ; a <= 9; a++ {
			if a*b == 45 {
				/*break OuterFor*/
				flag = true
			}
		}
		if flag {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
