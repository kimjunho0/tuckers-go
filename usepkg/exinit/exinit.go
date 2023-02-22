package exinit

import (
	"fmt"
	"go_practice/usepkg/custompkg"
)

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("exinit.init function:", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	custompkg.PrintCustom()
	fmt.Println("d:", d)
}
