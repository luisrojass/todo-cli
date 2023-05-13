package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/luisrojass/cli-crud/tasks"
)

func main() {
	// int 0666 means that has read and write permissions for the file
	file, err := os.OpenFile("db.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Start the program --------------------------------------------------

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// Check if the file has data
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		// Parse the encoded bytes to json and stores in tasks
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	// Make sure there are at least 2 arguments
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "ls":
		checkArguments(2, 2)
		task.PrintTasks(tasks)

	case "add":
		tmp := os.Args[2:]
		taskName := strings.Join(tmp, " ")
		tasks = task.AddTask(tasks, taskName)
		task.SaveTasks(tasks, file)

	case "do":
		checkArguments(3, 3)
		i := parseIndex(os.Args[2], len(tasks))
		tasks = task.CompleteTask(tasks, i)
		task.SaveTasks(tasks, file)

	case "undo":
		checkArguments(3, 3)
		i := parseIndex(os.Args[2], len(tasks))
		tasks = task.UndoTask(tasks, i)
		task.SaveTasks(tasks, file)

	case "rm":
		checkArguments(3, 3)
		i := parseIndex(os.Args[2], len(tasks))
		tasks = task.RemoveTask(tasks, i)
		task.SaveTasks(tasks, file)

	case "help":
		checkArguments(2, 2)
		printUsage()
		fmt.Println("\nFollow this project on github:\thttps://github.com/luisrojass/todo-cli")
		fmt.Println("Author: Luis Rojas S.\t\thttps://luisrojass.netlify.app")

	default:
		fmt.Printf("Unknown command: \"%s\"\n", strings.Join(os.Args[1:], " "))
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  ls\t\t\tShow all tasks")
	fmt.Println("  add <args> [string]\tAdd a task, no limit of words")
	fmt.Println("  do <args> [int] \tMark selected task as completed")
	fmt.Println("  undo <args> [int] \tMark selected task as incompleted")
	fmt.Println("  rm <args> [int]\tDelete selected task")
	fmt.Println("  help\t\t\tGet more information")
}

func checkArguments(min int, max int) {
	if len(os.Args) < min {
		fmt.Println("todo: More arguments are required")
		printUsage()
		os.Exit(0)

	} else if len(os.Args) > max {
		fmt.Println("todo: Too many arguments")
		printUsage()
		os.Exit(0)
	}
}

func parseIndex(index string, max int) int {
	i, err := strconv.Atoi(index)
	if err != nil || i < 0 || i > max {
		fmt.Println("todo: Invalid index")
		os.Exit(0)
	}
	return i
}
