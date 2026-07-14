package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCase(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func GetAllCases(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func CreateCase(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func UpdateCase(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func DeleteCaseByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
