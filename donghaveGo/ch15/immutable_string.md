# 문자열은 불변

go에서 문자열은 불변이기 때문에 string 타입이 가리키는 문자열의 일부만 변경할 수 없다. 따라서 아래 코드는 컴파일 에러가 발생한다.
``` go
var str string = "Hello World"
str = "How are you?" //전체를 바꾸는건 가능하지만,
str[2] = 'a' //일부를 바꾸는 것은 불가능하다.
```

만약에 문자열의 일부를 바꾸려면 아래와 같이 가능하다.
``` go
package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str) // 슬라이스로 타입 변환

	slice[2] = 'a' // 3번째 문자 변경

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
```
byte 슬라이스로 타입 변환을 한 후 문자 일부 변경이 가능한 이유는 str이 가리키는 메모리 공간과 slice가 가리키는 메모리 공간이 서로 다르기 때문이다. 서로 메모리 공간이 다르기 때문에 slice는 `Healo World`로 바뀌었지만, str은 여전히 `Hello World`이다.

아래의 코드로 실제 주솟값을 확인하면 더 명확하게 알 수 있다.
``` go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // ❶
	sliceheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringheader.Data) // ❷
	fmt.Printf("slice:\t%x\n", sliceheader.Data)
}
```

## 문자열 합산

go에서 string의 탑 연산을 하면 기존 문자열의 메모리 공간을 건드리지 않고, 새로운 메모리 공간을 만들어 두 문자열을 합치기 때문에 string 합 연산 이후 주솟값이 변경된다. 따라서 문자열 불변 원칙이 준수된다.

``` go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	addr1 :=stringheaderData                                    

	str += " World"            
	addr2 := stringheader.Data

	str += " Welcome!"         
	addr3 := stringheader.Data

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)
}
```
위의 코드를 통해 합연산을 진행할 때마다 주소가 다르게 출력되는 것을 볼 수 있다. 따라서 string 합 연산을 빈번하게 하면 메모리가 낭비된다.
그래서 string 합 연산을 빈번하게 사용하는 경우에는 strings 패키지의 Builder를 이용해서 메모리 낭비를 줄일 수 있다.

```go
package main

import (
	"fmt"
	"strings"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a')) // 합연산 사용
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a')) // strings.Builder 사용
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
```
`ToUpper1()` 함수와 `ToUpper2()` 함수 모두 소문자를 대문자로 바꾼다. `ToUpper1()` 함수는 합 연산을 사용해서 문자를 더한다. Go 언어 내부에서는 합 연산을 사용할 때마다 새로운 메모리 공간을 할당해 두 문자열을 더해 합 연산을 할 때마다 메모리 공간이 버려지기 때문에 메모리 공간 낭비와 성능 문제가 발생한다.

ToUpper2()는 `strings.Builder` 객체를 이용해 문자를 더한다. `strings.Builder`는 내부에 슬라이스를 가지고 있기 때문에 `WriteRune()` 메서드를 통해 문자를 더할 때 매번 메모리를 새로 새엉하지 않고 기존 메모리 공간에 빈자리가 있으면 그냥 더하기 때문에 메모리 공간 낭비를 없앨 수 있다.

## 왜 문자열 불변 원칙을 지키려할까?

왜 Go는 빈번한 합 연산 시 메모리가 낭비되는 데도 문자열 불변 원칙을 지키려 할까? 가장 큰 이유는 예기치 못한 버그를 방지하기 위해서이다.

만약 문자열 불변 윈칙이 없다면 문자열이 언제라도 변화할 수 있게 되어 string 타입 변수를 안심하고 사용할 수 없다.
``` go
func CahngeString(str3 string){
    str3[4] = 'T'
}

func main(){
    str := "Hello World"
    str2 := str

    ChangeString(str)
}
```
string 타입은 복사될 때 문자열 전체가 복사되는 것이 아니라 Data, Len 필드값만 복사된다. 따라서 str, str2, str3는 모두 같은 문자열을 가리키게 된다. 만약 문자열 불변 원칙이 없어서 `str3[4] = "T"` 와 같이 문자열 일부 값을 변경하면 str, str2, str3 모두 변경된 문자열을 가리키게 되는데 이는 프로그래머가 의도하지 않은 상황일 수 있고 string 변수값이 코드 전반에 걸쳐서 복사됐다면 언제 어디에서 문자열이 변경되었는지 알 수 없어서 많은 버그를 양산할 수 있다.

위와 같은 이유 때문에 string 타입을 안심하고 사용할 수 있게 하기 위해 go는 문자열의 불변 원칙을 지킨다. 

결국 핵심은 string 합 연산이 빈번하면 `strings.Builder`를 사용하라는 점이다.


