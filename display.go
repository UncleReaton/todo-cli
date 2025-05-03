package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	purple = lipgloss.Color("99")
	white  = lipgloss.Color("255")

	headerStyle = lipgloss.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
	cellStyle   = lipgloss.NewStyle().Padding(0, 1).Width(20)
	RowStyle    = cellStyle.Foreground(white)
)

func tasksToRows(tasks []Task) [][]string {
	var rows [][]string

	for _, task := range tasks {
		status := "TO DO"
		if task.Done {
			status = "DONE"
		}
		row := []string{
			fmt.Sprintf("%d", task.ID),
			task.Description,
			status,
		}
		rows = append(rows, row)
	}
	return rows
}

func printTable(tasks []Task) {
	rows := tasksToRows(tasks)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			default:
				return RowStyle
			}
		}).
		Headers("ID", "DESCRIPTION", "STATUS").
		Rows(rows...)

	fmt.Println(t.Render())
}
