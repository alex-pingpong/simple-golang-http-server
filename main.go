package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "1.0.0",
			"message": "index content",
			"code":    "200",
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "success",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		slog.Error("error", slog.String("error", err.Error()))
	}
}
