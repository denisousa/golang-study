package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetProducts ... Get all products
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":      "Notebook",
		"price":     2000.50,
		"categorie": "technology",
	})
}
