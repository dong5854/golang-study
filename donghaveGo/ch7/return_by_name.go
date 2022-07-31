package main

import "fmt"

// Go는 반환 타입을 지정할 때 변수명을 지정하면 return 문으로 명시적으로 해당 변수를 반환하지 않아도 값을 반환할 수 있다.
func Divide(a, b int) (result int, success bool) { // 반환할 변수명을 명시
	if b == 0 {
		result = 0
		success = false
		return // 명시적으로 반환할 값을 지정할 필요가 없음
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
