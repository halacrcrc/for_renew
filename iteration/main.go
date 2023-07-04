package main

import "fmt"

func fibonacci(i int) {

	var m, n, l = 1, 1, 2
	switch {
	case i < 3:
		fmt.Println(make([]int, 0))
	default:
		Fibonacci := make([]int, i)
		Fibonacci[0] = m
		Fibonacci[1] = n
		Fibonacci[2] = l
		for count := 2; count <= i-1; count++ {
			Fibonacci[count] = Fibonacci[count-1] + Fibonacci[count-2]
		}
		fmt.Println(Fibonacci)
	}
}

func main() {
	var i int

	fmt.Println("input:")
	fmt.Scanln(&i)
	fibonacci(i)
}
