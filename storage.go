package main

import (
	"encoding/json"
	"os"
)

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

//func LoadTasks() ([]Task, error) {
//	data, err := os.ReadFile("tasks.json")
//	if err != nil {
//		// Si le fichier n'existe pas encore, on retourne une liste vide
//		if errors.Is(err, os.ErrNotExist) {
//			return []Task{}, nil
//		}
//		return nil, err
//	}
//
//	var tasks []Task
//	err = json.Unmarshal(data, &tasks)
//	if err != nil {
//		return nil, err
//	}
//	return tasks, nil
//}
