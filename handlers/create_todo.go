package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTodoHandler(client *firestore.Client) func(c *gin.Context) {
	
	return func(c *gin.Context) {
		// what do write here?
		// 1-) bind the received/request json into a struct
		// 2-) fill in a few files
		// 3-) save into firestore
		
		var todo types.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJson(http.StatusBadRequest, gin.H{"error:",err.Error()})
			return
		}
		now := time.Now()
		todo.CreatedAt = now
		todo.UpdatedAt = now

		ref := client.Collection(types.TODO_COLLECTION).NewDoc()

		_, err := ref.Set(c, map[string]interface{}) {
			"title": todo.Title,
			"description": todo.description,
			"createdAt": todo.createdAt,
			"updatedAt": todo.updatedAt,
			"completed": todo.completed,
		}

		if err != nil {
			log.Printf("An error has occurred: %s", err)
			c.JSON(http.StatusInternalServerError, "")
		}

		c.JSON(http.StatusCreated, todo)
	}
}