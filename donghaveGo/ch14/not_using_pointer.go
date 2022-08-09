package main

import "fmt"

type Data struct { // Data형 구조체
	value int
	data  [200]int
}

func ChangeData(arg Data) { // 파라미터로 Data를 받음
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(data)                               // 인수로 data를 넣음, 변수값이 모두 복사되기 때문에 ChangeData()함수의 매개변수 arg와 data는 서로 다른 메모리 공간을 갖는 변수이다.
	fmt.Printf("value = %d\n", data.value)         // data 필드 출력, 다른 메모리 공간에 복사 되었기 때문에 변경되지 않음
	fmt.Printf("data[100] = %d\n", data.data[100]) // data 필드 출력, 다른 메모리 공간에 복사 되었기 때문에 변경되지 않음
}
