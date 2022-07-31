package main

import "fmt"

func main() {
	const C int = 10

	var b int = C * 20
	// C = 20 에러, 상수는 값을 바꿀 수 없음
	fmt.Println(b)
	// fmt.Fprintln(&C) 에러, 상수는 값으로만 동작하기 때문에 메모리 주솟값으로 접근할 수 없다.
}
