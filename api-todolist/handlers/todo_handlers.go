package handlers

import (
	"net/http"
	"strconv"
	"go-tour/models"
	"go-tour/services"
	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error" : message})
}

func GetAllTodos(c *gin.Context) {
	getAll, found := services.GetAllTodos()
	if !found {
		sendError(c, http.StatusBadRequest, "Data todolist kosong")
		return
	}

	c.JSON(http.StatusOK, getAll)
}

func GetTodoById(c *gin.Context) {
	idStr := c.Param("id")
	
	// Convert string to int
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(c, http.StatusBadRequest, "ID todo harus berupa angka")
		return
	}

	getTodo, success := services.GetTodoById(idInt)
	if !success {
		sendError(c, http.StatusBadRequest, "Data todolist kosong")
		return
	}

	c.JSON(http.StatusOK, getTodo)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	createTodo := services.CreateTodo(newTodo)
	if err := c.BindJSON(&newTodo); err != nil {
		sendError(c, http.StatusBadRequest, "Nama todo harus ada!")
		return 
	}

	c.JSON(http.StatusOK, createTodo)
}

func UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")

	var updateTodo models.Todo
	// Convert string to int
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(c, http.StatusBadRequest, "ID todo harus berupa angka")
		return
	}

	update, success := services.UpdateTodo(idInt, updateTodo)
	if !success {
		sendError(c, http.StatusBadRequest, "Gagal update todo!")
		return
	}

	c.JSON(http.StatusOK, update)
}

func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")

	// Convert string to int
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(c, http.StatusBadRequest, "ID todo harus berupa angka")
		return
	}

	success := services.DeleteTodo(idInt)
	if !success {
		sendError(c, http.StatusBadRequest, "Gagal menghapus todo")
		return
	}

	c.JSON(http.StatusOK, success)
}