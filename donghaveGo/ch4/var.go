package main

import "fmt"

func main() {
	var a int = 3
	var b int
	var c = 4
	d := 5
	fmt.Println(a, b, c, d)

	var e float64 = 3.6
	var f int = int(e)
	fmt.Println(e, f)

	var g int = int(e * 3)
	var h int = int(e) * 3
	fmt.Println(g, h)

	var i int16 = 3456
	var j int8 = int8(i)
	fmt.Println(i, j)
}
