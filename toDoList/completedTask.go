package main

import "fmt"

func completedTask() {
	var index int

	cleanScreen()
	showList()

	fmt.Println("Write the index of the task you want to complete:")
	fmt.Scan(&index)

	if index < len(g_list) {
		g_list[index].completed = true
	} else {
		fmt.Println("Index out of range")
	}
}