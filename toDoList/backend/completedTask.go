package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func completedTask(context *gin.Context) {
	var index int

	context.BindJSON(&index)

	if index < len(g_list) {
		g_list[index].Completed = !g_list[index].Completed
	} else {
		fmt.Println("Index out of range")
	}
}