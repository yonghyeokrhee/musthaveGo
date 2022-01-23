//ch8/ex8.1/ex8.1.go
package main

import "fmt"

func main() {
	const C int = 10 // ❶ 상수 선언

	var b int = C * 20 // ❷ 대입문 우변에 값으로 동작합니다.
	fmt.Println(&C)    //
}
