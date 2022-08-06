package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/calendar", getCalendar)

	router.Run("localhost:8080")
}
