package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type todoList struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tasks []task `json:"tasks"`
}

var todoLists = []todoList{}
var todoListsId int = 0
var taskId int = 0

func getTodoLists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"items": todoLists})
}

func postTodoLists(c *gin.Context) {
	var newTodoList todoList
	if err := c.BindJSON(&newTodoList); err != nil {
		// TODO: send some response?
		return
	}
	newTodoList.ID = strconv.Itoa(todoListsId) // TODO: use uuid instead
	todoListsId += 1

	todoLists = append(todoLists, newTodoList)
	c.IndentedJSON(http.StatusCreated, newTodoList)
}

func getTodoListsById(c *gin.Context) {
	id := c.Param("id")
	for _, t := range todoLists {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TodoList not found"})
}

func patchTodoListsById(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not implemented"})
}

func deleteTodoListsById(c *gin.Context) {
	id := c.Param("id")
	for i := range todoLists {
		if todoLists[i].ID == id {
			todoLists = append(todoLists[:i], todoLists[i+1:]...)
			c.Status(http.StatusOK)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TodoList not found"})
}

func postTasks(c *gin.Context) {
	var newTask task
	if err := c.BindJSON(&newTask); err != nil {
		// TODO: send some response?
		return
	}
	newTask.ID = strconv.Itoa(taskId) // TODO: use uuid instead
	taskId += 1

	id := c.Param("id")
	for i := range todoLists {
		if todoLists[i].ID == id {
			todoLists[i].Tasks = append(todoLists[i].Tasks, newTask)
			c.IndentedJSON(http.StatusCreated, todoLists[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TodoList not found"})
}

func getTasksById(c *gin.Context) {
	id := c.Param("id")
	taskId := c.Param("taskId")
	for _, t := range todoLists {
		if t.ID == id {
			for _, tk := range t.Tasks {
				fmt.Println("compare: ", tk.ID, taskId)
				if tk.ID == taskId {
					c.IndentedJSON(http.StatusOK, tk)
					return
				}
			}
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TodoList not found"})
}

func patchTasksById(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not implemented"})
}

func deleteTasksById(c *gin.Context) {
	id := c.Param("id")
	taskId := c.Param("taskId")
	for i := range todoLists {
		if todoLists[i].ID == id {
			for j := range todoLists[i].Tasks {
				if todoLists[i].Tasks[j].ID == taskId {
					todoLists[i].Tasks = append(todoLists[i].Tasks[:j], todoLists[i].Tasks[j+1:]...)
					c.Status(http.StatusOK)
					return
				}
			}
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TodoList not found"})
}

func main() {
	router := gin.Default()

	router.GET("/todolists", getTodoLists)
	router.POST("/todolists", postTodoLists)
	router.GET("/todolists/:id", getTodoListsById)
	router.PATCH("/todolists/:id", patchTodoListsById)
	router.DELETE("/todolists/:id", deleteTodoListsById)

	router.POST("/todolists/:id/tasks", postTasks)
	router.GET("/todolists/:id/tasks/:taskId", getTasksById)
	router.PATCH("/todolists/:id/tasks/:taskId", patchTasksById)
	router.DELETE("/todolists/:id/tasks/:taskId", deleteTasksById)

	router.Run("localhost:8089")
}
