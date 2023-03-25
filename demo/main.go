package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nomainc/frame"
	"github.com/nomainc/frame/version"
	"github.com/sirupsen/logrus"
)

func main() {

	fmt.Printf("commit: %20s\n", version.GitCommit)
	fmt.Printf("built on %20s\n", version.BuildGoVersion)
	fmt.Printf("built on %20s\n", version.BuildSystem)
	router := frame.Default()

	router.GET("/hello", func(c *gin.Context) {
		logrus.Info("哈哈")
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.Run()
}