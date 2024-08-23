package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func closeLinux() {
	cmd := exec.Command("wmctrl", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error listing windows:", err)
		return
	}

	for _, line := range strings.Split(string(output), "\n") {
		if len(line) > 0 {
			windowID := strings.Fields(line)[0]
			cmd := exec.Command("wmctrl", "-i", "-c", windowID)
			err := cmd.Run()
			if err != nil {
				log.Println("Error closing window:", err)
			}
		}
	}
}

func shutdownLinux() {
	err := exec.Command("shutdown", "-h", "now").Run()
	if err != nil {
		fmt.Println("Failed to shut down the system:", err)
		return
	}
	fmt.Println("System shutting down...")
}
