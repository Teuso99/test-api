package handlers

import (
	"net/http"
	"test-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var items = make(map[string]models.Item)

func GetItems(c *gin.Context) {
	list := make([]models.Item, 0, len(items))
	for _, v := range items {
		list = append(list, v)
	}
	c.JSON(http.StatusOK, list)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	if item, exists := items[id]; exists {
		c.JSON(http.StatusOK, item)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
}

func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = uuid.New().String()
	items[item.ID] = item
	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	if _, exists := items[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = id
	items[id] = item
	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if _, exists := items[id]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	delete(items, id)
	c.JSON(http.StatusNoContent, nil)
}
