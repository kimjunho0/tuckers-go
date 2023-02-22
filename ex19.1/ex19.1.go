package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 호출 과정 완전히 똑같음
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100} // *account

	withdrawFunc(a, 30)

	a.withdrawMethod(30) //a를 리시버로 받는 method 호출

	fmt.Printf("%d\n", a.balance)
}
