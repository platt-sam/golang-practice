package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task struct {
	name     string
	complete bool
}

func main() {

	validcommands := []string{
		"add", "complete", "delete", "help", "list", "quit",
	}

	var todolist []task

	var command, taskname string

	for command != "quit" {
		fmt.Printf("\nEnter a command: ")
		fmt.Scanln(&command)

		if command == "add" {
			fmt.Printf("\nEnter the name of the task to add it to your to do list: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				taskname = scanner.Text()
			}

			index := search(todolist, taskname)
			if index != -1 {
				fmt.Printf("\nThere is already a task with the name %s in your todo list. Consider using the command complete or delete instead", taskname)
			} else {
				todolist = add(task{name: taskname, complete: false})
			}

			list()
		} else if command == "complete" {
			fmt.Printf("\nEnter the name of the task to complete it")
			fmt.Scanln(&taskname)

			index := search(todolist, taskname)
			if index == -1 {
				fmt.Printf("\nA task with the name %s was not found. Consider adding this task to your todo list or confirming the correct taskname", taskname)
			} else {
				todolist[index].complete = true
			}
			list()
		} else if command == "delete" {
			fmt.Printf("\nEnter the name of the task to delete it")
			fmt.Scanln(&taskname)

			index := search(todolist, taskname)
			if index == -1 {
				fmt.Printf("\nA task with the name %s was not found. Consider adding this task to your todo list or confirming the correct taskname", taskname)
			} else {
				delete(index)
			}
			list()
		} else if command == "help" {
			fmt.Printf("\nValid commands are: %s", strings.Join(validcommands, ", "))
		} else if command == "list" {
			list()
		} else if command == "quit" {
			return
		} else {
			fmt.Printf("\nCommand not found. Consider using the help command to list all valid commands\n")
			return
		}

	}
}

func list() {
	fmt.Println("TODO LIST:")
	for _, x := range todolist {
		var status string
		if x.complete {
			status = "x"
		} else {
			status = " "
		}
		fmt.Printf("[%s] %s", status, x.name)
	}
}
func add(t task) []task {
	return append(todolist, t)
}
func complete(t task) {
	for i, x := range todolist {
		if x.name == t.name {
			todolist[i].complete = true
			return
		}
	}
	list()
}
func delete(i int) []task {
	return append(todolist[:i], todolist[i+1:]...)
}

/* Returns index of task if found, returns -1 if not found */
func search(tasks []task, name string) int {
	for i, x := range todolist {
		if x.name == name {
			return i
		}
	}
	return -1
}
