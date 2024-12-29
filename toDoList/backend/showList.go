package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showList(context *gin.Context) {
	var end int = findEmptyIndex()

	if end > 0  {
		context.IndentedJSON(http.StatusOK, g_list)
	} else if end == 0 {
		context.IndentedJSON(http.StatusNotFound, "error: there is not task in the list")
	}
}
