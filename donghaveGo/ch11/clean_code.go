package main

import "fmt"

// 클린 코드를 지향하려면 중첩된 내부 로직을 함수로 묶어 중첩을 줄이고 플래그 변수나 레이블 사용을 최소화한다.
func find45(a int) (int, bool) { // 곱해서 45가 되는 값을 찾는 함수
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0

	for ; a <= 9; a++ {
		var found bool
		if b, found = find45(a); found { // 함수 호출
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
