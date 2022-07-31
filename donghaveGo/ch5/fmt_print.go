package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	// 서식 문자를 전부 외울 것 까지는 없다. %f, %d, %s 정도만 기억,  %v는 데이터 타입의 기본 서식에 맞춰 출력
	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b", b, "f:", f)
	fmt.Printf("a: %d b: %d f:%f\n", a, b, f)

	// 최소 출력 너비 지정
	var c = 123
	var d = 456
	var e = 123456789
	fmt.Printf("%5d, %5d\n", c, d)
	fmt.Printf("%05d, %05d\n", c, d)
	fmt.Printf("%-5d, %-5d\n", c, d)

	fmt.Printf("%5d , %5d\n", e, e)
	fmt.Printf("%05d , %05d\n", e, e)
	fmt.Printf("%-5d , %-5d\n", e, e)

	// 소수점 출력
	f = 324.13455
	pi := 3.14
	fmt.Printf("%08.2f\n", f)
	fmt.Printf("%08.2g\n", f)
	fmt.Printf("%08.5g\n", f)
	fmt.Printf("%f\n", pi)
}
