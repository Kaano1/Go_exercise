package main

import "net/http"
import "github.com/gin-gonic/gin"


func removeFunc(index int) {
	g_list[index].Task = ""
	g_list[index].Completed = false
	for ; index < 100; index++ {
		if index + 1 < 100 {
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
	var index int 

	context.BindJSON(&index)

	for i := 0; i < 100; i++ {
		if g_list[i].ID == index {
			removeFunc(i)
			context.IndentedJSON(http.StatusCreated, index)
			return ;
		}
	}			
	context.IndentedJSON(http.StatusBadRequest, "error: there is not task in the list")
}
