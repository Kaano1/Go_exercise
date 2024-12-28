package main

import "github.com/gin-gonic/gin"

var g_list [100]List
var router *gin.Engine = gin.Default()

func main() {
	router.POST("/completedTask", completedTask)
	router.POST("/removeTask", removeTask)
	router.POST("/addTask", addTask)
	router.GET("/showList", showList)
	router.Run("localhost:9090");
}
