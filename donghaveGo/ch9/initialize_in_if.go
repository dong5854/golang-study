package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

//if 문을 검사하기 전에 초기문을 넣을 수 있다. 초기문은 검사에 사용할 변수를 초기화할 때 주로 사용
func main() {
	// 초기문에서 선언한 변수의 범위는 if문 안으로 한정된다.
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age) // age 값에 접근가능
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error")
	}

	// fmt.Println("Your age is", age) // 에러, age 소멸
}
