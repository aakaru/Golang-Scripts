package main

import (
	"fmt"
	"os/exec"
)

func closeMac() {
	cmd := exec.Command("osascript", "-e", `tell application "System Events" to quit every application`)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to close windows:", err)
	}
}

func shutdownMac() {
	err := exec.Command("shutdown", "-h", "now").Run()
	if err != nil {
		fmt.Println("Failed to shut down the system:", err)
		return
	}
	fmt.Println("System shutting down...")
}
