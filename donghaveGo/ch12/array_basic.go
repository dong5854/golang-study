package main

import "fmt"

func main() {
	nums := [...]int{10, 20, 30, 40, 50} // ...로 배열의 개수 생략 가능

	nums[2] = 300 // 값 대입

	for i := 0; i < len(nums); i++ { // for문으로 순회, len으로 배열의 길이 반환
		fmt.Println(nums[i]) // 출력
	}
}
