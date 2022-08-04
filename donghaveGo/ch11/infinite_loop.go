package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for { // go는 while문이 없어 이렇게 for를 사용해 무한루프를 구현한다.
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}
