package main

import (
	"log"
	"os/exec"
	"runtime"
)

func main() {
	url := "https://github.com"
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin", "windows":
		cmd = open(url)
	default:
		log.Fatalf("Cannot open url: %s", "operating system not supported")
	}

	if err := cmd.Run(); err != nil {
		log.Fatal("Open url error: %v", err)
	}
}
