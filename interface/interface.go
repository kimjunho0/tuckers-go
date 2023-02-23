package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p person) hisName() {
	fmt.Printf("his name is %s.\n", p.Name)
}
func (p person) hisAge() {
	fmt.Printf("his age is %d.\n", p.Age)
}

type GetName interface {
	hisName()
}

type GetAge interface {
	hisAge()
}

type InfoPerson interface {
	GetName
	GetAge
}

func main() {
	var i InfoPerson
	i = person{"kane", 30}
	i.hisName()
	i.hisAge()
}
