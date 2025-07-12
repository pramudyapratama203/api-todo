package services

import (
	"go-tour/models"
	"sync"
)

var todosData = make(map[int]models.Todo)
var todoMuted sync.Mutex
var nextId int

// array or slice input
func IntTodos() {
	todoMuted.Lock()
	defer todoMuted.Unlock()

	todosData[1] = models.Todo{ID: 1, Title: "Belajar Go Lang", IsCompleted: false}
	todosData[2] = models.Todo{ID: 2, Title: "Makan Siang", IsCompleted: true}
	nextId = 3
}

// Get all todos
func GetAllTodos() ([]models.Todo, bool){
	todoMuted.Lock()
	defer todoMuted.Unlock()

	var allTodos []models.Todo
	for _, todo := range todosData {
		allTodos = append(allTodos, todo)
	}

	return allTodos, true
}

// Get todo by id
func GetTodoById(id int) (models.Todo, bool) {
	todoMuted.Lock()
	defer todoMuted.Unlock()

	findTodo, found := todosData[id]
	return findTodo, found
}

// Create todo
func CreateTodo(newTodo models.Todo) models.Todo {
	todoMuted.Lock()
	defer todoMuted.Unlock()

	newTodo.ID = nextId
	newTodo.IsCompleted = false
	nextId++
	todosData[newTodo.ID] = newTodo
	return newTodo
}

// Update Todo
func UpdateTodo(id int, updateTodo models.Todo) (models.Todo, bool) {
	todoMuted.Lock()
	defer todoMuted.Unlock()

	update, found := todosData[id]
	if !found {
		return models.Todo{}, false
	}

	update.Title = updateTodo.Title
	update.IsCompleted = updateTodo.IsCompleted
	todosData[id] = update
	return update, found
}

// Delete todo
func DeleteTodo(id int) bool {
	todoMuted.Lock()
	defer todoMuted.Unlock()

	_, found := todosData[id]
	if !found {
		return false
	}

	delete(todosData, id)
	return true
}
