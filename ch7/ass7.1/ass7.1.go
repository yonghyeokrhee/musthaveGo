package main

import "fmt"

func Multiple(a, b int) (result int) {
	result = a * b

	return
}

func main() {
	c := Multiple(3, 4)
	fmt.Println(c)
}
