package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Task struct {
	Name     string    `json:"name"`
	Complete bool      `json:"complete"`
	Date     time.Time `json:"date"`
}

func PrintTasks(tasks []Task) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Done", "Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	var incompleteTasks int
	var date string

	for i, task := range tasks {
		complete := "✔"
		if task.Complete == false {
			complete = " "
			incompleteTasks++
		}
		date = fmt.Sprintf(
			"%d/%02d/%02d",
			task.Date.Year(),
			task.Date.Month(),
			task.Date.Day(),
		)
		tbl.AddRow(i+1, task.Name, complete, date)
	}

	tbl.Print()
	fmt.Printf("\nThere's (%d) remaining tasks\n", incompleteTasks)
}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		Name:     name,
		Complete: false,
		Date:     time.Now(),
	}

	return append(tasks, newTask)
}

func CompleteTask(tasks []Task, index int) []Task {
	tasks[index-1].Complete = true
	return tasks
}

func UndoTask(tasks []Task, index int) []Task {
	tasks[index-1].Complete = false
	return tasks
}

func RemoveTask(tasks []Task, i int) []Task {
	return append(tasks[:i-1], tasks[i:]...)
}

func SaveTasks(tasks []Task, file *os.File) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	// Clean the file before overwriting
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// Overwrite the file
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
