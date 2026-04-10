package main

import "fmt"

func AddTask(tasks *[]Task, title string, nextID *int) Task {
	t := Task{ID: *nextID, Title: title, Done: false}
	*tasks = append(*tasks, t)
	*nextID++
	return t
}

func markDone(tasks *[]Task, id int) error {
	for i := range *tasks {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("задача №%d не найдена", id)
}
