package main

import "fmt"

func main() {
	for i := 2; i < 10; i++ {

		if i >= 3 && i <= 6 {
			continue
		} else {
			for j := 1; j < 10; j++ {
				fmt.Println(i, "*", j, "=", i*j)
			}
			fmt.Println()

		}
	}
}
