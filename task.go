package main

type Task struct {
	ID          int
	Description string
	Priority    int
}

func generateID() int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func newTask(desc string) *Task {
	task := Task{Description: desc, Priority: 1, ID: generateID()}
	return &task
}
