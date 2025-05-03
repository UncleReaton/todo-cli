package main

type Task struct {
	ID          int
	Description string
	Done        bool
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
	task := Task{Description: desc, Done: false, ID: generateID()}
	return &task
}

func markTaskDone(tasks []Task, id int) []Task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			break
		}
	}
	return tasks
}
