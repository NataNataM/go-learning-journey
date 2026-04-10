package main

import (
	"encoding/json"
	"os"
)

type TaskStorage interface {
	Load() ([]Task, error)
	Save([]Task) error
}

type FileStorage struct{ Path string }

func (f *FileStorage) Load() ([]Task, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	return tasks, json.Unmarshal(data, &tasks)
}

func (f *FileStorage) Save(tasks []Task) error {
	data, _ := json.MarshalIndent(tasks, "", " ")
	return os.WriteFile(f.Path, data, 0644)
}

type InMemoryStorage struct{ data []Task }

func (m *InMemoryStorage) Load() ([]Task, error) {
	return m.data, nil
}

func (m *InMemoryStorage) Save(tasks []Task) error {
	m.data = tasks
	return nil
}
