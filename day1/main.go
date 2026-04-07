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

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     name := "Анна"
//     birthYear := 2001
//     currentYear := time.Now().Year()
//     age := currentYear - birthYear

//     fmt.Printf("Привет, %s! В %d году тебе исполнится %d лет.\n",
//                name, currentYear, age)

//     // Простая логика
//     if age < 30 {
//         fmt.Println("🎉 Ты принадлежишь к поколению миллениалов!")
//     } else {
//         fmt.Println("✨ Опыт — твоя суперсила!")
//     }
// }
