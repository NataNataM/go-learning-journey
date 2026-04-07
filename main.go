package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "2k"
	age, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Ошибка: не удалось преобразовать %s в число: %v/n", input, err)
		return
	}
	fmt.Printf("✅ Твой возраст: %d\n", age)
}
