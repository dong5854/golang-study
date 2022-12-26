package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3} // 대괄호를 사용해 slice 초기화

	for i := 0; i < len(slice); i++ { // 각 요소에 10 더하기
		slice[i] += 10
	}

	for i, v := range slice { // 각 요소에 2 곱하기
		slice[i] = v * 2
	}

	slice2 := append(slice, 4) // 요소 추가

	fmt.Println(slice)
	fmt.Println(slice2)

	slice2 = append(slice2, 11, 12, 13, 14, 15) // 한번에 여러 요소 추가
	fmt.Println(slice2)
}
