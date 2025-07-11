package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"      // Key
	"strconv"

	"github.com/gin-gonic/gin" 
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	IsCompleted bool `json:"isCompleted"`
}

var todosData = make(map[int]Todo)
var todosMuted sync.Mutex // Karena ga pakai database
var nextId int = 1

func sendErrorResponse(c *gin.Context, statusCode int, message string){
	c.JSON(statusCode, gin.H{"error" : message})
}

func main(){
	router := gin.Default()

	todosMuted.Lock()
	todosData[1] = Todo{ID: 1, Title: "Belajar Go Lang", IsCompleted: false}
	todosData[2] = Todo{ID: 2, Title: "Buat Database", IsCompleted: true}
	nextId = 3
	todosMuted.Unlock()

	router.GET("/todos", getAllTodos)
	router.GET("/todos/:id", getTodoById)
	router.POST("/todos", createNewTodo)
	router.PUT("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	fmt.Println("http://localhost:8080")
	log.Fatal(router.Run(":8080")) // Untuk mencatat log jika ada error
}

func getAllTodos(c *gin.Context){
	todosMuted.Lock()
	defer todosMuted.Unlock()

	var allTodos []Todo
	for _, todo := range todosData {
		allTodos = append(allTodos, todo)
	}

	c.JSON(http.StatusOK, allTodos)
}

func getTodoById(c *gin.Context){
	todoID := c.Param("id")

	todosMuted.Lock()
	defer todosMuted.Unlock()

	// Convert todoID to int
	todoIdInt, err := strconv.Atoi(todoID)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "ID tugas harus berupa angka")
		return
	}

	todoItem, found := todosData[todoIdInt]
	if !found {
		sendErrorResponse(c, http.StatusNotFound, "Tugas tidak tidak ditemukan")
		return
	}

	c.JSON(http.StatusOK, todoItem)
}

func createNewTodo(c *gin.Context){
	var newTodo Todo

	if err := c.BindJSON(&newTodo); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Format tugas tidak valid")
		return
	}

	todosMuted.Lock()
	defer todosMuted.Unlock()

	if newTodo.Title == "" {
		sendErrorResponse(c, http.StatusBadRequest, "Judul todolist tidak boleh kosong")
		return
	}

	newTodo.ID = nextId
	newTodo.IsCompleted = false
	todosData[newTodo.ID] = newTodo
	nextId++

	c.JSON(http.StatusOK, newTodo)
}

func updateTodo(c *gin.Context){
	todoID := c.Param("id")

	var updatedData Todo
	if err := c.BindJSON(&updatedData); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Format data update tidak valid")
		return
	}

	todosMuted.Lock()
	defer todosMuted.Unlock()

	todoIdInt, err := strconv.Atoi(todoID)
	if err != nil{
		sendErrorResponse(c, http.StatusBadRequest, "ID tugas harus berupa angka")
	}

	existingTodo, found := todosData[todoIdInt]
	if !found {
		sendErrorResponse(c, http.StatusNotFound, "Tugas tidak ditemukan!")
		return
	}

	existingTodo.Title = updatedData.Title
	existingTodo.IsCompleted = updatedData.IsCompleted

	todosData[todoIdInt] = existingTodo
	c.JSON(http.StatusOK, existingTodo)
}

func deleteTodo(c *gin.Context){
	todoID := c.Param("id")

	todosMuted.Lock()
	defer todosMuted.Unlock()

	todoIdInt, err := strconv.Atoi(todoID)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "ID Tugas harus berupa angka")
		return
	}

	_, found := todosData[todoIdInt]
	if !found {
		sendErrorResponse(c, http.StatusNotFound, "Tugas tidak ditemukan!")
		return
	}

	delete(todosData, todoIdInt)
	c.JSON(http.StatusOK, gin.H{"message": "Tugas berhasil dihapus!"})
}