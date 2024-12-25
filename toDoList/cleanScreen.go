package main

import "os"
import "os/exec"


func cleanScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}