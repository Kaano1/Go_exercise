package main

import "fmt"


var g_list [100]List


func showQuestions() int {
	result := -1
	questions := []string{"If you want to add a task, write", "If you want to remove a task, write", "If you want to see the list, write", "If you want to complete one of the task",  "If you want to exit, write"}

	cleanScreen()
	for index, question := range questions {
		fmt.Println(question, index)
	}
	fmt.Scan(&result)
	return result
}

func main() {
	var runFunction = []func() {
		addTask,
		removeTask,
		showList,
		completedTask,
		}
	var key int

	for key != 4 {
		key = showQuestions()
		if key <= 3 && key >= 0{
			runFunction[key]()
		} else if key != 4 {
			fmt.Println("Invalid value!")
		}
	}
	fmt.Println("See you again bro!")
}
