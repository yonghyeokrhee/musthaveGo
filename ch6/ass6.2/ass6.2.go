//ch6/ex6.2/ex6.2.go
package main

import "fmt"

func main() {
	var a uint8
	a |= 2
	a |= 4
	a |= 8

	var b uint8
	b = 4

	a &^= b
	fmt.Println(a)
}
