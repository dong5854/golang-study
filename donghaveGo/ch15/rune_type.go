package main

import "fmt"

func main() {
	/*문자하나를 표현하는데 rune 타입을 사용, golang은 UTF-8이 표준 문자코드인데 UTF-8은한 글자가 1~3바이트이기 때문에 3바이트가 필요
	go 언어 기본 타입은 3바이트 정수 타입은 제공하지 않기 때문에 rune 타입은 4바이트 정수 타입인 int32 타입의 별칭 타입이다.*/
	var char rune = '한'

	fmt.Printf("%T\n", char) // char 타입 출력, rune 타입은 int32 타입과 같기 때움에 int32 출력
	fmt.Println(char)        // char값 출력, char 값을 출력, char 타입이 int32라서 숫자로 출력
	fmt.Printf("%c\n", char) // 문자 출력, %c 포멧을 사용해 문자 하나를 출력
}
