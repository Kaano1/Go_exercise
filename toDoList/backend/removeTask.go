package main

import "net/http"
import "github.com/gin-gonic/gin"


func removeFunc(index int) {
	g_list[index].Task = ""
	g_list[index].Completed = false
	for ; index < len(g_list); index++ {
		if index + 1 < len(g_list) {
			g_list[index].ID = g_list[index+1].ID
			g_list[index].Task = g_list[index+1].Task
			g_list[index].Completed = g_list[index+1].Completed
			g_list[index + 1].Task = ""
			g_list[index + 1].Completed = false
			g_list[index + 1].ID = -1
		}
	}
}

func removeTask(context *gin.Context) {
	var list List

	index := list.ID
	context.BindJSON(&index)

	if index > -1 && index < len(g_list) {
		removeFunc(index)
		context.IndentedJSON(http.StatusCreated, index)
	} else {
		context.IndentedJSON(http.StatusBadRequest, "error: there is not task in the list")
		return
	}
}
