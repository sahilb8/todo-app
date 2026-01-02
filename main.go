package main

import (
	"flag"
	"fmt"
	"time"
)

type Task struct {
	Title       string
	Completed   bool
	CreatedTime time.Time
}

type Todos []*Task

func (t *Todos) add(title string) {
	fmt.Println("Adding task:", title)

	task := &Task{
		Title:       title,
		Completed:   false,
		CreatedTime: time.Now(),
	}

	*t = append(*t, task)
}

func (t *Todos) list() {
	for i, task := range *t {
		fmt.Println(i+1, task.Title, task.Completed, task.CreatedTime)
	}
}

func main() {

	add := flag.String("add", "", "Add a new task (e.g., -add 'My task')")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Mark a task as completed (e.g., -complete 1)")
	del := flag.Int("delete", 0, "Delete a task (e.g., -delete 1)")

	flag.Parse()

	fmt.Println("add:", *add)
	fmt.Println("list:", *list)
	fmt.Println("complete:", *complete)
	fmt.Println("del:", *del)
	fmt.Println("tail:", flag.Args())

	myTodos := Todos{}

	if *add != "" {
		myTodos.add(*add)
		myTodos.add("hello world")
	}

	if *list {
		myTodos.list()
	}

}
