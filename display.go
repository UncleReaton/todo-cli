package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	idStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	descStyle = lipgloss.NewStyle().Bold(true)
	prioStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Italic(true)
)

func printTask(task Task) {
	id := idStyle.Render(fmt.Sprintf("[%2d]", task.ID))
	desc := descStyle.Render(task.Description)
	prio := prioStyle.Render(fmt.Sprintf("(prio: %d)", task.Priority))

	fmt.Println(id, desc, prio)
}

func printList(tasks []Task) {
	for _, task := range tasks {
		printTask(task)
	}
}
