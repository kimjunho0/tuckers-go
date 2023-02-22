package main

import "fmt"

type inform struct {
	height int
	weight int
}

func (infor inform) hw() int {

	i := infor.height * infor.weight
	return i
}

func main() {
	i := inform{180, 70}
	fmt.Println(i.hw())
	fmt.Println(i)
}
