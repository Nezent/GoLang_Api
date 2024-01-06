package main

import (
	"errors"
	"net/http"
	"strconv"

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

func getTodoById(id int) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i],nil

		}
	}
	return nil,errors.New("todo not found")
}

func getTodo(context *gin.Context){
	id,_ := strconv.Atoi(context.Param("id"))
	todo,err := getTodoById(id)
	
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK,todo)
}

func toggleCompletion(context *gin.Context){
	id,_ := strconv.Atoi(context.Param("id"))
	todo,err := getTodoById(id)
	
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK,todo)
}

func main(){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id",getTodo)
	router.PATCH("/todos/:id",toggleCompletion)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}