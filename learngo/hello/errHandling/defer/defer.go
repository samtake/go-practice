package main

import (
	"bufio"
	"fmt"
	"learngo/function/fib"
	"os"
)

// func tryDefer() {
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)

// 	for i := 0; i < 100; i++ {
// 		defer fmt.Println(i)

// 		if i == 30 {
// 			panic("printed too many")
// 		}
// 	}
// }

// func writeFle(filename string) {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)
// 	defer writer.Flush()

// 	f := fib.Fibonacci()
// 	for i := 0; i < 20; i++ {
// 		fmt.Fprintln(writer, f())
// 	}
// }

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			// Uncomment panic to see
			// how it works with defer
			// panic("printed too many")
		}
	}
}

func writeFle(filename string) {
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFle("fib.txt")

	tryDefer()
}
