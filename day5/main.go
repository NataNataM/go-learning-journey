package main

import (
	"fmt"
	"time"
)

func printNumbers(id int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Горутина %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond)

	}
}

func main() {
	go printNumbers(1)
	go printNumbers(2)

	fmt.Println("main: запустил горутины")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main: завершаюсь")
}
