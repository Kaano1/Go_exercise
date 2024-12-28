package main

func findEmptyIndex() int {
	for index, list := range g_list {
		if list.Task == "" {
			return index
		}
	}
	return -1
}