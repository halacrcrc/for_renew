package main

import "fmt"

func transf(i string) int {
	trans := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	tempvals := make([]int, len(i)+1)

	for index, data := range i {
		if val, judge := trans[data]; judge {
			tempvals[index] = val
		} else {
			fmt.Printf("Error: The Roman numeral %s has a bad digit: %c\n", i, data)
			return 0
		}
	}

	sum := 0

	for index := 0; index < len(i); index++ {
		if tempvals[index] < tempvals[index+1] {
			tempvals[index] = -tempvals[index]
		}
		sum += tempvals[index]
	}

	return sum
}

func main() {
	var i string
	fmt.Println("plz input: ")
	fmt.Scanln(&i)
	fmt.Println("the ouput: ", transf(i))
}
