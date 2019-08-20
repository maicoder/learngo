package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = true
	ss = "kkk"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 	欧拉
func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b float64 = 3.0000, 4.0000
	var c float64
	c = math.Sqrt(a*a + b*b)
	fmt.Println(a*a)
	fmt.Println(c)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)

	euler()

	triangle()
}
