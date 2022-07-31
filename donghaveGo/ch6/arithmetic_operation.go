package main

import "fmt"

func main() {
	//go는 연산자의 타입이 같아야 하기 때문에 타입이 같도록 맞춰준 다음에 연산을 해야한다.
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)

	// fmt.Println("x + t = ", x+t) 타입이 맞지 않아 에러
}
