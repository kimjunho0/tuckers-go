package main

import "fmt"

func main() {
	str := "Hello 월드"
	for i := 0; i < len(str); i++ { // len()은 바이트 길이 반환
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
}
