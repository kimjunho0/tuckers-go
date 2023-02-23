package main

import "fmt"

//함수타입 변수
//별칭타입 선언

type OperateFunc func(int, int) int

func getOperator(op string) OperateFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	var op OperateFunc

	op = getOperator("+")
	result := op(20, 30)
	fmt.Println(result)

	op = getOperator("*")
	result = op(10, 20)
	fmt.Println(result)
}
