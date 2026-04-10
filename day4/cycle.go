package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type App struct {
	store TaskStorage
}

func NewApp(store TaskStorage) *App {
	return &App{store: store}
}

func (a *App) Run() {
	tasks, err := a.store.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки данных:", err)
		tasks = []Task{}
	} else {
		fmt.Printf("Загружено задач: %d\n", len(tasks))
	}
	nextID := 1
	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Консольный To-Do List")
	fmt.Println("Команды: add, list, done <id>, del <id>, quit")

	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		cmd := parts[0]

		switch cmd {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Использование: add <название задачи>")
				continue
			}
			newTask := AddTask(&tasks, parts[1], &nextID)

			fmt.Printf("Задача №%d '%s'добавлена", newTask.ID, newTask.Title)

			if err := a.store.Save(tasks); err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}

		case "list":
			if len(tasks) == 0 {
				fmt.Println("Список пуст")
				continue
			}
			for _, t := range tasks {
				status := "no OK"
				if t.Done {
					status = "OK"
				}
				fmt.Printf("%s №%d: %s\n", status, t.ID, t.Title)
			}

		case "done":
			id := parseID(parts[1])
			if id == 0 {
				continue
			}
			err := markDone(&tasks, id)
			if err != nil {
				fmt.Println("Ошибка", err)
				continue
			}

			fmt.Println("Задача выполнена!")

			if err := a.store.Save(tasks); err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}

		case "del":
			id := parseID(parts[1])
			if id == 0 {
				continue
			}
			for i, t := range tasks {
				if t.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("Задача удалена")
					break
				}
			}

		case "quit":
			fmt.Println("До встречи")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func parseID(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil || id <= 0 {
		fmt.Printf("Введите корректный ID задачи")
		return 0
	}
	return id
}

func main() {
	store := &FileStorage{Path: "data.json"}
	app := NewApp(store)
	app.Run()
}
