package main

import (
	"fmt"
	"time"
)

var quit = make(chan bool)

func fib(ch chan int) {
	x, y := 1, 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	var command string

	start := time.Now()

	ch := make(chan int)

	go fib(ch)

	for {
		num := <-ch
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	//time.Sleep(1 * time.Second)
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

}
