package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1byte
	C int8 // 1byte
	E int8 // 1byte
	B int  // 8byte
	D int  // 8byte
	// 19byte

}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println((unsafe.Sizeof(user)))
}
