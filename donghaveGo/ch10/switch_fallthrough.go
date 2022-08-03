package main

import (
	"fmt"
)

func main() {

	var a int
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(err)
	}

	switch a {
	case 1:
		fmt.Println("a == 1")
		break //go는 break를 꼭 넣어주지 않아도 자동으로 break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough // fallthrough를 사용하면 다음 case문까지 실행한다.
	case 4:
		fmt.Println("a == 4")
	case 5:
		fmt.Println("a == 5")
	default:
		fmt.Println("a > 4")
	}
}
