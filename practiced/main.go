package main

import "fmt"

type Product struct {
	name  string
	price int
}

func main() {
	m := make(map[int]Product)

	m[24] = Product{"볼펜", 500}
	m[50] = Product{"샤프", 1000}
	m[60] = Product{"샤프심", 2000}
	m[82] = Product{"필통", 7000}
	m[100] = Product{"물감", 3000}

	for key, value := range m {
		fmt.Println(key, value)
	}
}
