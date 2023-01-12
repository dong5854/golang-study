# 슬라이스의 동작 원리

슬라이스는 내장 타입으로 내부 구현이 감춰져 있지만 reflect 패키지의 Slice4Header 구조체를 사용해 내부 구현을 살펴볼 수 있다.
```go
type SliceHEader struct {
    Data uintptr // 실제 배열을 가리키는 포인터
    Len int // 요소 개수
    Cap int // 실제 배열의 길이
}
```
슬라이스는 배열을 가리키는 포인터와 요소 개수ㅠ를 나타내는 len, 전체 배열 길이를 나타내는 cap 필드로 구성된 구조체입니다. 슬라이스가 실제 배열을 가리키는 포인터를 가지고 있어서 쉽게 크기가 다른 배열을 가리키도록 변경할 수 있고, 슬라이스 변수 대입 시 배열에 비해 사용되는 메모리나 속도에 이점이 있다.

## make() 함수를 이용한 선언

make() 함수를 사용해 슬라이스를 만들 때, 파라미터를 2개 혹은 3개를 넣는다.
```go
var slice = make([]int, 3)
```
slice는 len이 3 이고, cap이 3이다. 즉 요소 개수(len)도 3, 총 배열의 길이(cap)도 3 이다.
```go
var slice2 = make([]int, 3, 5)
```
slice2는 len이 3 이고 cap이 5이다. 즉 요소 개수(len)는 3, 총 배열 길이(cap)는 5이다. 총 5개 중 3개만 사용하고 나머지 2개는 나중에 추가될 요소를 위해서 비워뒀다고 보면 된다.

slice에서 사용할 공간을 적절하게 reserve 해주지 않으면 slice의 크기가 capacity를 넘을 때마다 2 2배씩 늘려주는데 이렇게 capacity를 늘려주는 과정에서 메모리 할당이 일어나게 되는데 이는 성능에 영향을 미친다.

## 슬라이스와 배열 동작 차이

슬라이스의 내부 구현은 배열과 다르기 때문에 동작도 배열과 매우 달라, 슬라이스와 배열이 사용법이 비슷하다고 똑같이 사용하면 예기치 못한 버그를 만날 수 있다.

```go
package main

import "fmt"

func changeArray(array2 [5]int) { // 배열을 받아서 세 번째 값 변경
	array2[2] = 200
}

func changeSlice(slice2 []int) { // 슬라이스를 받아서 세 번째 값 변경
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}
```
위 코드의 결과는 아래와 같다.
```
array: [1 2 3 4 5]
slice: [1 2 200 4 5]
```
결과를 보면 slice의 3 번째 값은 200으로 바뀌었는데 changeArray()의 인수로 사용한 array 배열의 3번째 값이 바뀌지 않는다. 이런 현상이 일어가는 이유는 배열과 슬라이스의 구조가 서로 다르기 때문이다.

## 동작 차이의 원인

Go 언어에서 모든 값의 대입은 복사로 일어난다. 함수에 인수로 전달될 때와 다른 변수에 대입할 때 값의 이동은 복사로 일어난다. 복사는 타입의 값이 복사되는데, 포인터는 포인터의 값인 메모리 주소가 복사되고, 구조체가 복사될 때는 구조체의 모든 필드가 복사된다. 배열은 배열의 모든 값이 복사된다.

따라서 위의 코드에서 배열이 함수에 인수로 전달 될 때 메모리 공간이 다른 완전히 다른 배열로 복사가 되어 넘어가기 때문에 array 가 복사된 array2의 세번째 값을 200으로 변경해도 array 배열은 변경되지 않는 것이다.

이에 반에 slice 타입인 []int 는 포인터, len, cap 세 개의 필드를 갖는 구조체이다. 포인터는 메모리 주소로 8 파이트의 길이이고, len과 cap은 각각 int 타입으로 8바이트의 길이이기 때문에 슬라이스 크기는 24바이트이다. changeSlice() 함수의 인수로 slice 가 입력되어 호출되면 slice 내부의 포인터가 가리키는 배열 크기에 상관없이 항상 총 24바이트 값이 복사된다.
위의 예시에서 slice 깞이 slice2로 복사되면 구조체의 각 필드값이 복사되기 때문에 포인터의 메모리 주소값, len, cap값이 복사된다.
복사된 것은 메모리 주소값이기 떄문에 slice와 slice2는 같은 배열 데이터를 가리키기 때문에, slice2의 세 번째 요솟값을 바꿔도 slice 또한 같은 배열을 가리키기 때문에 slice[2] 값 또한 바뀌게 된다.

## append()를 사용할 때 발생하는 예기치 못한 문제1
append 함수의 동작을 알아보자, append() 함수가 호출되면 먼저 슬라이스에 값을 추가할 수 있는 빈 공간이 있는지 확인한다. 남은 빈 공간은 실제 배열의 cap 에서 슬라이스 요소 개수 len 을 뺀 값이다.
```
남은 빈 공간 = cap - len
```
남은 빈 공간의 개수가 추가하는 값의 개수보다 크거나 같은 경우 배열 뒷부분에 값을 추가한 뒤 len 값을 증가시킨다.
```go
package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5) // len:3 cap:5 슬라이스를 만든다

	slice2 := append(slice1, 4, 5)
	// cap() 함수를 이용해 슬라이스 capacity 값을 알 수 있다.
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 // slice2까지 바뀝니다.

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500) // 역시 slice2까지 바뀝니다.

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
```
위의 코드를 통해 slice1과 slice2는 같은 배열을 가리키는 것을 알 수 있다.