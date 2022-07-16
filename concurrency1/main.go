package main

import (
	"fmt"
	//"math/rand"
	"time"
)

func fib(ch chan string, number float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)

}

func main() {
	start := time.Now()
	size := 15
	ch := make(chan string, size)

	for i := 0; i < size; i++ {
		go fib(ch, float64(i))
	}

	for i := 0; i < size; i++ {
		fmt.Printf(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
