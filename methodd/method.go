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
		if c == "y" {
			i.Name = Name
			break
		} else {
			return
		}
	}
}
func (i *information) WhatAge(Age int) {
	for {
		var c string
		fmt.Printf("나이가 %d살맞나요?\ny/n : ", Age)

		if c == "y" {
			i.Age = Age
			break
		} else {
			return
		}
	}
}

func main() {
	var a string
	d := information{Name: ""}
	d.WhatName(a)
}
