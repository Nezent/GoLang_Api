package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        int     `json:"id"`
	Item      string  `json:"item"`
	Completed bool    `json:"completed"`
}

var todos = []todo{
	{ID: 1, Item: "Do BreakFast", Completed: false},
	{ID: 2, Item: "Go Practice", Completed: false},
	{ID: 3, Item: "Do Assignment", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context){
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated,newTodo)
}

func main(){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}