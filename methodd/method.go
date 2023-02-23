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
func (i *information) WhatGender(Gender string) {
	for {
		var c string
		fmt.Printf("성별이 %s맞나요?\ny/n : ", Gender)
		fmt.Scanln(&c)
		if c == "y" {
			i.Gender = Gender
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

	for {
		fmt.Println("성별 입력 : male or female")
		fmt.Scanln(&c)
		if c == "male" {
			c = "남자"
			f.WhatGender(c)
			break
		} else if c == "female" {
			c = "여자"
			f.WhatGender(c)
			break
		} else {
			continue
		}

	}
	fmt.Printf("이름 : %s\n나이 : %d\n성별 : %s", d.Name, e.Age, f.Gender)
}
