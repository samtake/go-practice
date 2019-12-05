package main

import (
	"fmt"
	"time"
)

func creatWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()

	return c
}
func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = creatWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)

}

// func work(id int, c chan int) {
// 	for {
// 		n, ok := <-c
// 		if !ok {
// 			break
// 		}

// 		fmt.Printf("Worker %d received %c\n", id, n)
// 	}
// }

func work(id int, c chan int) {
	for {
		for n := range c {
			fmt.Printf("Worker %d received %c\n", id, n)
		}
	}
}

func bufferedChannel() {
	c := make(chan int, 3) //给个缓存3
	go work(0, c)
	c <- 1
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go work(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
