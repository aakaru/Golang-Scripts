package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("Running on %s...", runtime.GOOS)

	switch runtime.GOOS {
	case "windows":
		closeWindows()
		shutdownWindows()
	case "darwin":
		closeMac()
		shutdownMac()
	case "linux":
		closeLinux()
		//shutdownLinux()
	default:
		fmt.Println("Unsupported OS.")
	}

}
