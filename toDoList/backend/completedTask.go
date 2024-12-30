package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func completedTask(context *gin.Context) {
	var index int

	context.BindJSON(&index)

	for i := 0; i < 100; i++ {
		if g_list[i].ID == index {
			g_list[index].Completed = !g_list[index].Completed;
			context.IndentedJSON(http.StatusOK, index)
			return ;
		}
	}			
}