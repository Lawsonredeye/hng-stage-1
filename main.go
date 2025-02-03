package main

import (
	"stage-two/controllers"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/api/classify-number", controllers.ClassifyNumber)
	r.Run(":5050")
}