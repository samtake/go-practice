package main

import "fmt"

func exp() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}
func main() {
	exp()

	a, b := 3, 4
	swap(&a, &b)
	println(a, b)
}
