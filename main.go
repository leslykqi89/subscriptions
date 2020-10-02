package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leslykqi89/subscriptions/database"
)

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Subscription API.",
	})
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusNotFound,
		"message": "This endpoint does not exist.",
	})
}

func handleRequest() {
	router := gin.Default()

	router.GET("/", homePage)
	router.POST("/subscriptor", database.CreateSubscriptor)
	router.GET("/subscriptor/:id", database.GetSubscriptor)
	router.GET("/subscriptors", database.GetAllSubscriptors)
	router.PUT("/subscriptor/:id", database.ModifySubscritor)
	router.DELETE("/subscript/:id", database.DeleteSubscritor)
	router.NoRoute(notFound)
	router.Run(":8080")
}

func main() {
	database.Connection()
	handleRequest()
}
