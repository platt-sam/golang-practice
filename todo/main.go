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

var todolist []task

func main() {

	validcommands := []string{
		"add", "complete", "delete", "help", "list", "quit",
	}

	var command, input, taskname string

	for command != "quit" {
		fmt.Printf("\nEnter a command: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
			tmp := strings.Split(input, " ")
			command = tmp[0]
			taskname = ""
			if len(tmp) > 1 {
				taskname = strings.Join(tmp[1:], " ")
			}
		}

		if command == "add" {
			for taskname == "" {
				fmt.Printf("\nEnter name of task")
				if scanner.Scan() {
					taskname = scanner.Text()
				}
			}

			index := search(todolist, taskname)
			if index != -1 {
				fmt.Printf("\nThere is already a task with the name %s in your todo list. Consider using the command complete or delete instead", taskname)
			} else {
				todolist = add(task{name: taskname, complete: false})
			}

			list()
		} else if command == "complete" {
			for taskname == "" {
				fmt.Printf("\nEnter name of task")
				if scanner.Scan() {
					taskname = scanner.Text()
				}
			}

			index := search(todolist, taskname)
			if index == -1 {
				fmt.Printf("\nA task with the name %s was not found. Consider adding this task to your todo list first or confirming the correct taskname", taskname)
			} else {
				todolist[index].complete = true
			}
			list()
		} else if command == "delete" {
			for taskname == "" {
				fmt.Printf("\nEnter name of task")
				if scanner.Scan() {
					taskname = scanner.Text()
				}
			}

			index := search(todolist, taskname)
			if index == -1 {
				fmt.Printf("\nA task with the name %s was not found. Consider adding this task to your todo list or confirming the correct taskname", taskname)
			} else {
				todolist = delete(index)
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
		fmt.Printf("\n[%s] %s", status, x.name)
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
