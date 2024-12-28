package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showList(context *gin.Context) {
	var end int = findEmptyIndex()

	fmt.Println(g_list[0])
	if end > 0  {
		context.IndentedJSON(http.StatusOK, g_list)
	} else if end == 0 {
		context.IndentedJSON(http.StatusOK, g_list)
	}
}
