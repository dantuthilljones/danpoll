package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type calendar struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

var cal1 = calendar{
	ID: "1", Title: "Blue Train", StartDate: "John Coltrane", EndDate: "dsa",
}

func getCalendar(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cal1)
}
