package main

import (
	"mission_service/infrastructure/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	defer db.DB.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
