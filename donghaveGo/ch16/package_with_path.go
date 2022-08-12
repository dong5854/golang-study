package main

import ( // 둘 이상의 패키지는 소괄호로 묶어준다.
	"fmt"
	"math/rand" // 패키지명은 rand이다.
)

func main() {
	fmt.Println(rand.Int()) // 랜덤한 숫자값을 출력한다.
}
