package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	// ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func PrintTasks(tasks []Task) {
	var incompleteTasks int

	for i, task := range tasks {

		var complete rune = '✔'
		if task.Complete == false {
			complete = ' '
			incompleteTasks++
		}

		fmt.Printf("  [%d][%c] %s\n", i+1, complete, task.Name)
	}

	fmt.Printf("There's (%d) remaining task", incompleteTasks)
	if incompleteTasks > 1 {
		fmt.Printf("s")
	}
	fmt.Printf("\n")
}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		Name:     name,
		Complete: false,
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
