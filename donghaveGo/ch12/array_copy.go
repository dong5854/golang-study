package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a { // 배열 a 원소 출력
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()         // 개행
	for i, v := range b { // 배열 b 원소 출력
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a // a 배열을 b변수에 복사

	fmt.Println()         // 개행
	for i, v := range b { // 배열 b 원소 출력
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b[1] = 100        // 깊은 복사인지 얕은 복사인지 확인
	fmt.Println(a[1]) // 2
	fmt.Println(b[1]) // 100, 즉 깊은 복사
}
