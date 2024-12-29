package main

import "time"
import "github.com/gin-contrib/cors"
import "github.com/gin-gonic/gin"


var g_list [100]List
var router *gin.Engine = gin.Default()

func main() {
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST","PUT", "PATCH"},
        AllowHeaders:     []string{"*"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "https://github.com"
        },
        MaxAge: 12 * time.Hour,
    }))
	router.POST("/completedTask", completedTask)
	router.POST("/removeTask", removeTask)
	router.POST("/addTask", addTask)
	router.GET("/showList", showList)
	router.Run("localhost:9090")
}

