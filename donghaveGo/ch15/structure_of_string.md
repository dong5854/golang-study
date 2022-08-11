# 문자열 구조

string 타입은 Go 언어에서 제공하는 내장 타입으로 내부 구현이 감추어져 있다.
하지만 reflect 패키지 안의 StringHeader 구조체를 통해서 내부 구현을 엿볼 수 있다.
``` go
type StringHeader struct {
    Data uintptr
    Len int
}
```
string은 필드가 2개인 구조체이다. 첫 번째 필드 Data는 uintptr 타입으로 문자열의 데이터가 있는 메모리 주소를 나타내는 포인터이다. 두 번째 필드 Len은 int 타입으로 문자열의 길이를 나타낸다.

string 변수가 대입되면 어떤 일이 일어날까
```go
package main

import "fmt"

func main() {
	str1 := "안녕하세요. 한글 문자열입니다."
	str2 := str1

	fmt.Printf(str1)
	fmt.Printf("\n")
	fmt.Printf(str2)
}
```
str2 변수에 str1 변수를 대입하니 문자열이 잘 출력된다.
여기서 str1의 Data 포인터값와 Len값만을 str에 복사된다.
여기서 문자열의 주솟값인 Data가 복사되었기 때문에 문자열 자체가 복사되지는 않는다. 
따라서 문자열이 아무리 길어도 문자열 전체가 복사되어 메모리나 성능문제가 생실 걱정은 할 필요없다.

정말로 위의 설명대로 복사가 되는지 아래의 코드로 확인할 수 있다.


``` go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello World!"
	str2 := str1 // str1 변수값을 str2에 복사

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)) // Data값 추출
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)) // Data값 추출

	fmt.Println(stringHeader1) // 각 필드 값을 출력
	fmt.Println(stringHeader2)
}
```
Go는 string 타입에서 `*reflect.StringHeader` 타입으로 변환을 막고 있어 string 타입을 `*reflect.StringHeader`타입으로 강제 변환을 시키기 위해 `unsafe.Pointer(&str1)`을 사용해 unsafe.Pointer 타입으로 변환한 다음에 `*reflect.StringHeader` 타입으로 변환한다. 내부 필드값을 접근하고자 강제로 타입 변환을 한다고 이해하면 된다.

위 코드를 통해 두 string의 Data 값이 같은 것을 확인할 수 있다.