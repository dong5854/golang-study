package main

import "fmt"

func main() {
	str1 := "가나다라마" // 한글 문자열
	str2 := "abcde" // 영문 문자열

	fmt.Printf("len(str1) = %d\n", len(str1)) // 한글 문자열 크기, 글자 수인 5가 아니라 문자열 크기인 15가 나온다.
	fmt.Printf("len(str2) = %d\n", len(str2)) // 영문 문자열 크기, 글자 수인 5가 나온다. UTF-8에서 영문자는 1바이트기 때문
}
