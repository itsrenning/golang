package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "New Test AWS Code-Pipeline",
			"deploy":  "success",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"server": "Golang server",
			"status": 200,
		})
	})
	err := r.Run(":5000")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func method1() {

}

func method2() {}

func method3() {}
