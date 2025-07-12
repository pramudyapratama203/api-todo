package main

import (
	"fmt"
	"go-tour/handlers"
	"go-tour/services"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	services.IntTodos()

	// Router
	router.GET("/todo", handlers.GetAllTodos)
	router.GET("/todo/:id", handlers.GetTodoById)
	router.POST("/todo", handlers.CreateTodo)
	router.PUT("/todo/:id", handlers.UpdateTodo)
	router.DELETE("/todo/:id", handlers.DeleteTodo)

	fmt.Println("Success! http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
