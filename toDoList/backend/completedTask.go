package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func completedTask(context *gin.Context) {
	var index int

	fmt.Println("Write the index of the task you want to complete:")	

	if index < len(g_list) {
		g_list[index].Completed = true
	} else {
		fmt.Println("Index out of range")
	}
}