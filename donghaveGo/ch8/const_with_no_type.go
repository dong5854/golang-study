package main

import "fmt"

const PI = 3.14              // 타입 없는 상수
const FloatPI float64 = 3.14 // float64 타입 상수

// 상수 선언 시 타입을 명시하지 않을 수 있다. 그러면 타입 없는 상수가 된다. 타입없는 상숫값은 타입이 정해지지 않은 상태로 사용된다.
// 타입 없는 상수는 변수에 복사될 때 타입이 정해지기 때문에 여러 타입에 사용되는 상숫값을 사용할 때 편리하다.
func main() {
	var a int = PI * 100 // 오류 발생하지 않습니다.
	// var b int = FloatPI * 100  오류, 타입 오류 발생

	fmt.Println(a)
	// fmt.Println(b)
}
