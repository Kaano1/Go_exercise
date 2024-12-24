package main

import "fmt"

type List struct {
	task string
	completed bool
}

var g_list []List

func showQuestions() int {
	result := -1
	questions := []string{"If you want to add a task, write 1", "If you want to remove a task, write 2", "If you want to see the list, write 3", "If you want to exit, write 4"}

	for index, question := range questions {
		fmt.Println(index, question)
	}
	fmt.Scan(&result)
	return result
}

func addTask() {
	fmt.Println("Write the task you want to add:")
	fmt.Scan(&g_list[len(g_list) - 1].task)
	g_list[len(g_list) - 1].completed = false
}

func completedTask() {
	var index int
	fmt.Println("Write the index of the task you want to complete:")
	fmt.Scan(&index)
	if index < len(g_list) {
		g_list[index].completed = true
	} else {
		fmt.Println("Index out of range")
	}
}

func removeTask() {
	var index int
	fmt.Println("Write the index of the task you want to remove:")
	fmt.Scan(&index)
	if index < len(g_list) {
		g_list = append(g_list[:index], g_list[index+1:]...)
	} else {
		fmt.Println("Index out of range")
	}
}

func showList() {
	for index, task := range g_list {
		fmt.Println(index, task)
	}
}

func main() {
	var runFunction = []func() {
		addTask,
		completedTask,
		removeTask,
		showList,}
	var key int

	for key != 5 {
		key = showQuestions()
		if key <= 4 && key >= 1{
			runFunction[key-1]()
		} else {
			fmt.Println("Invalid value!")
		}
	}

}
