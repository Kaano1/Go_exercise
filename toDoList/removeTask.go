package main

import "fmt"

func removeFunc(index int) {
	g_list[index].task = ""
	g_list[index].completed = false
	for ; index < len(g_list); index++ {
		if index + 1 < len(g_list) {
			g_list[index].task = g_list[index + 1].task
			g_list[index].completed = g_list[index + 1].completed
			g_list[index + 1].task = ""
			g_list[index + 1].completed = false
		}
	}
}

func removeTask() {
	var index int

	cleanScreen()
	showList()

	fmt.Println("Write the index of the task you want to remove:")
	fmt.Scan(&index)
	if index > -1 && index < len(g_list) {
		removeFunc(index)
	} else {
		fmt.Println("Index out of range")
	}
}