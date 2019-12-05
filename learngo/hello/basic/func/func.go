package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func div(a, b int) (q, r int) {
	// return a / b, a % b
	q = a / b
	r = a % b
	return q, r
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d,%d)", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	q, r := div(8, 6)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 1, 5))

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 1, 5))

	println(sum(1, 2, 5, 56, 1))
}
