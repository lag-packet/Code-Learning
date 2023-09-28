package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Item represents a single todo item.
type Item struct {
	ID   int
	Task string
}

// TodoList represents a list of todo items.
type TodoList struct {
	items []Item
	mux   sync.RWMutex
}

// NewTodoList creates a new todo list.
func NewTodoList() *TodoList {
	return &TodoList{}
}

// Add adds a new item to the todo list.
func (t *TodoList) Add(item Item) {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.items = append(t.items, item)
}

// GetAll returns all items in the todo list.
func (t *TodoList) GetAll() []Item {
	t.mux.RLock()
	defer t.mux.RUnlock()

	return t.items
}

func main() {
	todoList := NewTodoList()

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		task := r.FormValue("task")

		if task == "" {
			http.Error(w, "Task cannot be empty", http.StatusBadRequest)
			return
		}

		item := Item{
			ID:   len(todoList.GetAll()) + 1,
			Task: task,
		}

		todoList.Add(item)

		fmt.Fprintf(w, "Task added successfully")
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		items := todoList.GetAll()

		for _, item := range items {
			fmt.Fprintf(w, "ID: %d, Task: %s\n", item.ID, item.Task)
		}
	})

	http.ListenAndServe(":8080", nil)
}
