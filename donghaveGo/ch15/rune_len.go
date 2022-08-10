package main

import "fmt"

func main() {
	/*
		string과 rune 슬라이스 타입은 상호 타입 변환이 가능하다.
		string 타입을 []rune 타입 변환을 하면 각 글자들로 이뤄진 배열로 변환된다.
		따라서 len(runes)와 같이 len()의 인수로 배열을 넣으면 배열의 요소 개수가 반환된다.
		string 타입은 연속된 바이트 메모리라면 []rune 타입은 글자들의 배열로 이루어져 있다.
		그래서 이 둘은 완전히 다른 타입이지만 편의를 위해서 Go언어는 둘의 상호 타입 변환을 지원한다.
	*/
	str := "hello 월드"    // 한글과 영문자가 섞인 문자열
	runes := []rune(str) // []rune 타입으로 타입 변환

	fmt.Printf("len(str) = %d\n", len(str))     // string 타입 길이, 12
	fmt.Printf("len(runes) = %d\n", len(runes)) // []rune 타입 길이, 8
}
