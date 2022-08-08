package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) { // 파라미터로 Data 포인터를 받음
	arg.value = 999 // arg 데이터 변경
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data)                              // 인수로 data의 주소를 넘김, 1608바이트의 구조체 전부가 복사되는 게 아닌 8바이트의 메모리 주소만 복사
	fmt.Printf("value = %d\n", data.value)         // data 필드값 출력, arg 포인터 변수가 data 변수의 메모리 주소를 값으로 가지고 있어서 Data 구조체의 내부 필드값을 변경할 수 있다.
	fmt.Printf("data[100] = %d\n", data.data[100]) // data 필드값 출력, 역시 필드값이 변경되어 있다. 포인터를 이요하면 더 효율적으로 데이터 조작 가능
}
