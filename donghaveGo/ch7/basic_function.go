package main

import "fmt"

// argument(인수)는 함수를 호출 할 때 입력하는 값이고, parameter(매개변수)는 함수가 외부로부터 입력받는 변수
// 밑의 함수에서 parameter는 a, b이다.
func Add(a int, b int) int {
	return a + b
}

// 밑의 함수를 호출할 때 입력한 3, 5 argument이다.
func main() {
	// Go에서 함수의 핵심 포인트는 인수는 매겨변수로 복사된다는 점이다.
	// Go는 이런 값 전달(call by value)만 지원하기 때문에 레퍼런스 전달(call by reference)를 위해서는 포인터를 사용해야한다.
	c := Add(3, 5)
	fmt.Println(c)
}
