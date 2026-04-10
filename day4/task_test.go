package main

import "testing"

func TestMarkDone(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []Task
		targetID int
		wantErr  bool
	}{
		{"found", []Task{{ID: 1, Done: false}}, 1, false},
		{"not_found", []Task{{ID: 1}}, 5, true},
		{"multiple_tasks_found", []Task{
			{ID: 1, Done: false},
			{ID: 2, Done: false},
			{ID: 3, Done: false},
		}, 2, false},
		{"already_done", []Task{{ID: 1, Done: true}}, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := markDone(&tt.tasks, tt.targetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("markDone() error = %v, want %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				found := false
				for _, task := range tt.tasks {
					if task.ID == tt.targetID {
						if !task.Done {
							t.Errorf("expected task ot be done")
						}
						found = true
						break
					}
				}
				if !found && !tt.wantErr {
					t.Errorf("task №%d not found in list after markDone", tt.targetID)
				}
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	tests := []struct {
		name          string
		initialTasks  []Task
		title         string
		initialNextID int
		expectedID    int
		expectedLen   int
	}{
		{
			name:          "add_to_empty_list",
			initialTasks:  []Task{},
			title:         "Первая задача",
			initialNextID: 1,
			expectedID:    1,
			expectedLen:   1,
		},
		{
			name:          "add_to_non_empty_list",
			initialTasks:  []Task{{ID: 1, Title: "Старая", Done: false}},
			title:         "Новая задача",
			initialNextID: 2,
			expectedID:    2,
			expectedLen:   2,
		},
		{
			name:          "add_multiple_times",
			initialTasks:  []Task{},
			title:         "Задача A",
			initialNextID: 10,
			expectedID:    10,
			expectedLen:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tasks := make([]Task, len(tt.initialTasks))
			copy(tasks, tt.initialTasks)
			nextID := tt.initialNextID

			newTask := AddTask(&tasks, tt.title, &nextID)

			if newTask.ID != tt.expectedID {
				t.Errorf("AddTask() returned task with ID = %d, want %d", newTask.ID, tt.expectedID)
			}
			if newTask.Title != tt.title {
				t.Errorf("AddTask() returned task with Title = %s, want %s", newTask.Title, tt.title)
			}
			if newTask.Done != false {
				t.Errorf("AddTask() returned task with Done = %v, want false", newTask.Done)
			}
			if len(tasks) != tt.expectedLen {
				t.Errorf("After AddTask(), len(tasks) = %d, want %d", len(tasks), tt.expectedLen)
			}
			if nextID != tt.expectedID+1 {
				t.Errorf("After AddTask(), nextID = %d, want %d", nextID, tt.expectedID+1)
			}
		})
	}
}
