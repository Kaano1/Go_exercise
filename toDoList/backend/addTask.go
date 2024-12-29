package main

import "github.com/gin-gonic/gin"

func addTask(context *gin.Context) {

	var task List;

	context.BindJSON(&task);
	g_list[findEmptyIndex()].ID = task.ID;
	g_list[findEmptyIndex()].Task = task.Task;
	g_list[findEmptyIndex()].Completed = task.Completed;
}
