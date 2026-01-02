package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"time"
)

type Task struct {
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	CreatedTime time.Time `json:"created_time"`
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
		status := "[ ]"
		if task.Completed {
			status = "[✓]"
		}
		fmt.Printf("%d %s %s\n", i+1, task.Title, status)
	}
}

func (t *Todos) complete(index int) error {
	if index < 0 || index >= len(*t)+1 {
		return fmt.Errorf("invalid task index: %d", index)
	}
	(*t)[index-1].Completed = true
	return nil
}

func (t *Todos) delete(index int) error {
	if index < 0 || index >= len(*t)+1 {
		return fmt.Errorf("invalid task index: %d", index)
	}
	*t = slices.Delete(*t, index-1, index)
	return nil
}

func (t *Todos) save(fileName string) error {
	data, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}
	os.WriteFile(fileName, data, 0644)
	return nil
}

func (t *Todos) load(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	add := flag.String("add", "", "Add a new task (e.g., -add 'My task')")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Mark a task as completed (e.g., -complete 1)")
	del := flag.Int("delete", 0, "Delete a task (e.g., -delete 1)")

	flag.Parse()

	todoFileName := "todos.json"
	todos := &Todos{}

	err := todos.load(todoFileName)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		todos.add(*add)
	case *list:
		todos.list()
	case *complete > 0:
		err := todos.complete(*complete)
		if err != nil {
			fmt.Println("Error:", err)
		}
	case *del > 0:
		err := todos.delete(*del)
		if err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Invalid command")
	}

	err = todos.save(todoFileName)
	if err != nil {
		fmt.Println("Error saving todos:", err)
		os.Exit(1)
	}
}
