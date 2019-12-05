package main

import "fmt"

//返回一个函数
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println((a(i)))
	}

	f := fibonacci()
	fmt.Println("fibonacci")

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
