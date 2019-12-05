package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// for i := 0; i < 1000; i++ {
	// 	go func(i int) {
	// 		for {
	// 			fmt.Printf("hello form goroutine %d\n", i)
	// 		}
	// 	}(i)
	// }

	// time.Sleep(time.Millisecond)

	tryGoroutine()
}

func tryGoroutine() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
