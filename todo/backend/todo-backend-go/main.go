package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type todo struct {
    ID string `json:"id"`
    Title string `json:"title"`
}

var todos = []todo{
    {ID: "1", Title: "first task"},
    {ID: "2", Title: "second task"},
}

func getTodos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(c *gin.Context) {
    id := c.Param("id")
    for _, t := range todos {
        if t.ID == id {
            c.IndentedJSON(http.StatusOK, t)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func postTodos(c *gin.Context) {
    var newTodo todo
    if err := c.BindJSON(&newTodo); err != nil {
        // send some response?
        return
    }

    todos = append(todos, newTodo)
    c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
    router := gin.Default()
    router.GET("/todos", getTodos)
    router.GET("/todos/:id", getTodoById)
    router.POST("/todos", postTodos)

    router.Run("localhost:8089")
}
