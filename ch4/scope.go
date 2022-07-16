package main

import "fmt"

var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// 에러 발생
	// m = s + 20
}
