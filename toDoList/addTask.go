package main

import "fmt"

func addTask() {

	var task string

	cleanScreen()
	fmt.Println("Write the task you want to add:")
	fmt.Scan(&task)
	g_list[findEmptyIndex()].task = task
	g_list[findEmptyIndex()].completed = false
}