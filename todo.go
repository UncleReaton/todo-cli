package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
)

var tasks []Task

type Task struct {
	ID          int
	Description string
	Priority    int
}

func newTask(desc string) *Task {
	task := Task{Description: desc, Priority: 1, ID: len(tasks) + 1}
	return &task
}

func printTask(task Task) {
	fmt.Println("ID:", task.ID, task.Description, "-", task.Priority)
}

func listTasks() {
	for _, task := range tasks {
		printTask(task)
	}
}

func saveToFile() {
	file, err := os.Create("tasks.json")
	if err != nil {
		log.Error("Error creating file", "err", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		log.Error("Error encoding tasks", "err", err)
	}
}

func loadFromFile() {
	file, err := os.Open("tasks.json")
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		log.Error("Error decoding tasks", "err", err)
	}
}

func deleteTask(id int) {
	var updatedTask []Task
	found := false

	for _, task := range tasks {
		if task.ID != id {
			updatedTask = append(updatedTask, task)
		} else {
			found = true
		}
	}

	if found {
		tasks = updatedTask
		saveToFile()
		log.Info("Task deleted", "id", id)
	} else {
		log.Warn("Task not found", "id", id)
	}
}

func main() {
	loadFromFile()

	if len(os.Args) < 2 {
		log.Info("Usage: todo add <task description>")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			log.Info("Please provide a task description.")
			return
		}
		task := newTask(os.Args[2])
		tasks = append(tasks, *task)
		saveToFile()
		format := "%s %s"
		log.Infof(format, "Task added:", task.Description)

	case "list":
		listTasks()

	case "done":
		if len(os.Args) < 3 {
			log.Error("Please provide the task ID to mark as done")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Error("Invalid task ID", "value", os.Args[2])
			return
		}
		deleteTask(id)

	default:
		format := "%s %s"
		log.Infof(format, "Unknown command:", os.Args[1])
	}
}
