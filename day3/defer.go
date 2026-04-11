package main

import (
	"encoding/json"
"fmt"
	"os"
)

func readFile() error {
	f, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer f.Close()

	data := make([]byte, 100)
	n, err := f.Read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:n]))
	return nil
}

func safeParseJSON(data []byte) (result map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Критическая ошибка при парсинге: %v", err)
		}
	}()
	return nil, json.Unmarshal(data, &result)
}

func saveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

func loadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("read file: %w", err)
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return tasks, nil
}
