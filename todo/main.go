package main

import (
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

	if len(os.Args) <= 1 {
		fmt.Println("no arguments detected")
	}

	command := os.Args[1]

	newtask := task{
		name:     strings.Join(os.Args[2:], " "),
		complete: false,
	}

	switch command {
	case "add":
		add(newtask)
	case "complete":
		complete(newtask)
	case "delete":
		delete(newtask)
	case "list":
		list()
	case "quit":
		os.Exit(0)
	case "help":
		help()
	}

}

func help() {
	// TODO: list all commands
}
func list() {
	fmt.Println(todolist)
}
func add(t task) {
	todolist = append(todolist, t)
	list()
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
func delete(t task) {
	for _, x := range todolist {
		if x.name == t.name {
			// TODO: delete element
			return
		}
	}
	list()
}
