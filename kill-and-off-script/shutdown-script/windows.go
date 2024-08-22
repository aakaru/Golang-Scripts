package main

import (
	"fmt"
	"os/exec"
)

func closeWindows() {
	cmd := exec.Command("taskkill", "/F", "/FI", "STATUS eq RUNNING")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to close windows:", err)
	}
}

func shutdownWindows() {
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to shut down the system:", err)
		return
	}
	fmt.Println("System shutting down...")
}

func main() {
	closeWindows()
	shutdownWindows()
}
