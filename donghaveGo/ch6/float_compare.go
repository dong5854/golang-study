package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c) // false가 나옴
	fmt.Println(a + b)                                    // 더하면 값이 0.30000000000000004임을 볼 수 있음
}
