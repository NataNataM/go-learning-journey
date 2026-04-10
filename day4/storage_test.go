package main

import "testing"

type MockStorage struct {
	LoadFunc  func() ([]Task, error)
	SaveFunc  func([]Task) error
	SaveCalls int
}

func (m *MockStorage) Load() ([]Task, error) {
	return m.LoadFunc()
}

func (m *MockStorage) Save(t []Task) error {
	m.SaveCalls++
	return m.SaveFunc(t)
}

func TestAppWorkflow(t *testing.T) {
	mock := &MockStorage{
		LoadFunc: func() ([]Task, error) {
			return []Task{{ID: 1, Title: "Old"}}, nil
		},
		SaveFunc: func([]Task) error {
			return nil
		},
	}
	tasks, _ := mock.Load()
	tasks = append(tasks, Task{ID: 2, Title: "New"})
	_ = mock.Save(tasks)
	if mock.SaveCalls != 1 {
		t.Errorf("Save должен быть вызван 1 раз, вызван %d", mock.SaveCalls)
	}
}
