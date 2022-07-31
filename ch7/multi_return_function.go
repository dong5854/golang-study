package main

import "fmt"

// Go는 함수의 반환값이 여러개가 될 수 있다는 특징이 있다.
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false //제수가 0 일때 반환
	}
	return a / b, true //제수가 0 이 아닐 때 반환
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
