package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {
	fmt.Println("Task Manager API")

	router := gin.Default()
	// Get all tasks
	router.GET("/tasks", getTasks)

	// get task by id
	router.GET("/tasks/:id", getTask)

	// add task
	router.POST("/tasks", addTask)

	// update task by id
	router.PATCH("/tasks/:id", updateTask)

	//delete task by id
	router.DELETE("/tasks/:id", removeTask)

	router.Run()
}

func getTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func getTask(ctx *gin.Context) {
	task_id := ctx.Param("id")
	for _, task := range tasks {
		if task.ID == task_id {
			ctx.JSON(http.StatusOK, gin.H{
				"task": task,
			})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "task not found",
	})
}

func addTask(ctx *gin.Context) {
	var newTask Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Task Created",
	})
}

func updateTask(ctx *gin.Context) {
	task_id := ctx.Param("id")
	var updatedTask Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	for i, task := range tasks {
		if task.ID == task_id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "task updated"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func removeTask(ctx *gin.Context) {
	var task_id = ctx.Param("id")
	for i, task := range tasks {
		if task.ID == task_id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
