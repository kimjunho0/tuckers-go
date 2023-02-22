package main

import (
	"fmt"
)

type information struct {
	Name   string
	Age    int
	Gender string
}

func (i *information) WhatName(Name string) {
	for {
		var c string
		fmt.Printf("이름이 %s맞나요?\ny/n : ", Name)
		fmt.Scanln(&c)
		if c == "y" {
			i.Name = Name
			break
		} else {
			main()
			break
		}
	}
}
func (i *information) WhatAge(Age int) {
	for {
		var c string
		fmt.Printf("나이가 %d살맞나요?\ny/n : ", Age)
		fmt.Scanln(&c)

		if c == "y" {
			i.Age = Age
			break
		} else {
			main()
			break
		}
	}
}

func main() {
	var a, c string
	var b int

	d := information{Name: ""}
	fmt.Println("이름 입력 : ")
	fmt.Scanln(&a)
	d.WhatName(a)
	e := information{Age: 0}
	fmt.Println("나이 입력 : ")
	fmt.Scanln(&b)
	e.WhatAge(b)
	f := information{Gender: ""}
	fmt.Println("성별 입력 : ")
	fmt.Scanln(&b)
	f.WhatAge(b)
}
