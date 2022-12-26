package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	slice[1] = 10 // 슬라이스를 길이를 초과해서 접근하면 런 타임 에러 발생(패닉)
	fmt.Println(slice)
}
