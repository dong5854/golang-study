package main

import "fmt"

// iota를 통해 간편하게 열거값을 사용할 수 있다.

// iota는 0부터 시작해 서서히 1씩 증가
const (
	A1 int = iota // 0
	A2 int = iota // 1
	A3 int = iota // 2
)

// iota는 소괄호를 벗어나면 초기화된다.
// 첫 번째 값과 똑같은 규칙이 계속 적용된다면 타임과 iota를 생략 가능
const (
	C1 uint = iota + 1 // 1 = 0 + 1
	C2                 // 2 = 1 + 1
	C3                 // 3 = 2 + 1
)

const (
	BitFlag1 uint = 1 << iota // 1 = 1 << 0
	BitFlag2                  // 2 = 1 << 1
	BitFlag3                  // 4 = 1 << 2
	BitFlag4                  // 8 = 1 << 3
)

func main() {
	fmt.Println(A1, A2, A3)
	fmt.Println(C1, C2, C3)
	fmt.Println(BitFlag1, BitFlag2, BitFlag3, BitFlag4)
}
