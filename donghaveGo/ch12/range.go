package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} // float64 타입 요소 5개를 갖는 배열 t를 할당

	for i, v := range t { // range를 사용해 모든 배열 요소를 순회
		fmt.Println(i, v) //첫번째 i는 인덱스, 두번째 v는 원소값
	}
}
