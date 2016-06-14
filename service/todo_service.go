package service

import (
	"starter/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// TodoService is an unexported type
type TodoService struct {
	DB *gorm.DB
}

// Add is an unexported type
func (obj *TodoService) Add(c *gin.Context) {
	title := c.DefaultQuery("title", "")
	completeDateStr := c.DefaultQuery("completedDate", "")
	priorityStr := c.DefaultQuery("priority", "")
	color := c.DefaultQuery("color", "")

	if title != "" && priorityStr != "" && color != "" {

		var modelObj model.Todo

		modelObj.TodoTitle = title
		modelObj.Color = color

		priority, _ := strconv.Atoi(priorityStr)
		modelObj.Priority = priority

		completeDateStr += "+03:00"
		completeDate, _ := time.Parse(time.RFC3339, completeDateStr)
		modelObj.CompletedDate = completeDate

		obj.DB.Create(&modelObj)
		c.JSON(200, modelObj)
	} else {
		c.JSON(404, gin.H{"error": "title, priority, color parameters required"})
	}
}

// Get is an unexported type
func (obj *TodoService) Get(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		title := c.DefaultQuery("title", "")
		priorityStr := c.DefaultQuery("priority", "")
		color := c.DefaultQuery("color", "")

		priority, _ := strconv.Atoi(priorityStr)

		var modelObjs []model.Todo

		obj.DB.Where(&model.Todo{
			TodoTitle: title,
			Priority:  priority,
			Color:     color,
		}).Find(&modelObjs)

		c.JSON(200, modelObjs)
		return
	}

	var modelObj model.Todo
	if obj.DB.Where("todo_id = ?", id).First(&modelObj).RecordNotFound() {
		c.JSON(404, gin.H{"error": "record not found"})
	} else {
		c.JSON(200, modelObj)
	}
}

// Update is an unexported type
func (obj *TodoService) Update(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(404, gin.H{"error": "id parameter required"})
		return
	}

	var existing model.Todo

	if obj.DB.Where("todo_id = ?", id).First(&existing).RecordNotFound() {
		c.JSON(404, gin.H{"error": "record not found"})
	} else {

		title := c.DefaultQuery("title", "")
		if title != "" {
			existing.TodoTitle = title
		}

		completeDateStr := c.DefaultQuery("completedDate", "")
		if completeDateStr != "" {
			completeDateStr += "+03:00"
			completeDate, _ := time.Parse(time.RFC3339, completeDateStr)
			existing.CompletedDate = completeDate
		}

		priorityStr := c.DefaultQuery("priority", "")
		if priorityStr != "" {
			priority, _ := strconv.Atoi(priorityStr)
			existing.Priority = priority
		}

		color := c.DefaultQuery("color", "")
		if color != "" {
			existing.Color = color
		}

		obj.DB.Save(&existing)
		c.JSON(200, existing)
	}
}

// Delete is an unexported type
func (obj *TodoService) Delete(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		c.JSON(404, gin.H{"error": "id parameter required"})
		return
	}

	var modelObj model.Todo

	if obj.DB.Where("todo_id = ?", id).First(&modelObj).RecordNotFound() {
		c.JSON(404, gin.H{"error": "record not found"})
	} else {
		obj.DB.Delete(&modelObj)
		c.JSON(200, modelObj)
	}
}
