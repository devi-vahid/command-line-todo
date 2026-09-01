package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/go-task/task"
)

const taskFile = "task.json"

func main() {
	var list task.List

	if err := list.Load(taskFile); err != nil {
		printError(err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go-task <add|list|delete|done|find>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		runAdd(&list, os.Args[2:])
	case "list":
		runList(list)
	case "done":
		runDone(&list, os.Args[2:])
	case "delete":
		runDelete(&list, os.Args[2:])
	case "find":
		runFind(list, os.Args[2:])
	case "update":
		runUpdate(&list, os.Args[2:])
		return
	default:
		fmt.Println("Unknown command:", command)
	}
}

func printError(err error) {
	fmt.Println("Error:", err)
}

func runAdd(list *task.List, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: go-task add <task>")
		return
	}
	title := strings.Join(args, " ")
	if err := list.Add(title); err != nil {
		printError(err)
		return
	}
	if err := list.Save(taskFile); err != nil {
		printError(err)
		return
	}
}

func runList(list task.List) {
	tasks := list.All()
	if len(tasks) == 0 {
		fmt.Println("No tasks")
		return
	}
	for _, item := range tasks {
		fmt.Printf("%d. %s - %s\n", item.ID, item.Title, item.Status())
	}
}

func runDone(list *task.List, args []string) {
	id, ok := parseTaskID(args)
	if !ok {
		fmt.Println("Usage go-task update <task-id> <task-new-title>")
		return
	}
	if !list.Complete(id) {
		fmt.Println("Error: task not found")
		return
	}
	if err := list.Save(taskFile); err != nil {
		printError(err)
		return
	}

	fmt.Printf("task %d marked as done\n", id)
}

func runDelete(list *task.List, args []string) {
	id, ok := parseTaskID(args)
	if !ok {
		fmt.Println("Usage go-task update <task-id> <task-new-title>")
		return
	}

	if !list.Delete(id) {
		fmt.Println("Error: task not found")
		return
	}

	if err := list.Save(taskFile); err != nil {
		printError(err)
		return
	}

	fmt.Printf("task %d removed successfully\n", id)
}

func parseTaskID(args []string) (int, bool) {
	if len(args) == 0 {
		return 0, false
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: task id must be a number")
		return 0, false
	}
	return id, true
}

func runFind(list task.List, args []string) {
	id, ok := parseTaskID(args)
	if !ok {
		fmt.Println("Usage: go-task find <task-id>")
		return
	}
	task, ok := list.Find(id)
	if !ok {
		fmt.Println("Error: task not found")
		return
	}

	fmt.Printf("%d. %s - %s\n", task.ID, task.Title, task.Status())
}

func runUpdate(list *task.List, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage go-task update <task-id> <task-new-title>")
		return
	}
	id, ok := parseTaskID(args)
	if !ok {
		fmt.Println("Usage: go-task update <task-id> <task-new-title>")
		return
	}
	title := strings.Join(args[1:], " ")
	task, err := list.UpdateTitle(id, title)
	if err != nil {
		printError(err)
		return
	}
	fmt.Printf("%d. %s - %s\n", task.ID, task.Title, task.Status())
}
