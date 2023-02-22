package main

import "fmt"

func main() {
	var a int16 = 3456   // a는 int16타입 - 2바이트 정수
	var b int8 = int8(a) // int 16 -> int 8
	var c int8 = 127
	var d = int16(c)

	fmt.Println(a, b, c, d)
}
