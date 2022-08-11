package main

import "fmt"

func main() {
	// range를 사용하면 추가 메모리 할당 없이 문자열 순회가 가능하다.
	str := "Hello 월드!"      // 한영 문자가 섞인 문자열
	for _, v := range str { // range를 이용한 순회
		fmt.Printf("타입:%T 값:%d 문자:%c\n", v, v, v) // 출력
	}
}
