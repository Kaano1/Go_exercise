package main

import "fmt"
import "time"

func showList() {
	var end int = findEmptyIndex()
	for index := 0; index < end; index++ {
		fmt.Println("•", "[", index,"]", g_list[index].task)
		fmt.Println("Completed status:", g_list[index].completed)
		fmt.Println()
	}
	if end == 0 {
		fmt.Println("Sorry, We didn't find a list")
	}
	time.Sleep(3 * time.Second)
}
