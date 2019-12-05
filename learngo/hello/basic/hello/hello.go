package main

import (
	"fmt"
	"math"
	"runtime"
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
	var a, b, c, d = 3, 4, true, "def"
	var s string = "abc"
	fmt.Println(a, b, c, d, s)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	b = 6
	var s string = "abc"
	fmt.Println(a, b, c, d, s)
}

var aa = 55
var ss = "llll"
var bb = true

var (
	bbb = 666
	sss = "jjj"
	ttt = false
)

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b + b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		golang
		swift
		_
		C
	)

	const (
		b = 1 << (10 * iota)
		bb
		bbb
		bbbb
	)

	fmt.Println(cpp, java, golang, C)
	fmt.Println(b, bb, bbb, bbbb)
}

func main() {

	fmt.Println(runtime.GOARCH)

	variableZeroValue()

	variableInitialValue()

	variableTypeDeduction()

	variableShorter()

	fmt.Println(aa, bbb)

	consts()

	enums()

}
