package main

import "fmt"

var sroot, x float32 = 1, 0

func main() {
	fmt.Scanln(&x)
	for i := 0; i < 10; i++ {
		temp := sroot
		sroot = sroot - (sroot*sroot-x)/(2*sroot)
		if temp == sroot {
			fmt.Print("Square root is: ", sroot)
			break
		}
		fmt.Println("A guess for square root is ", sroot)
	}
}
