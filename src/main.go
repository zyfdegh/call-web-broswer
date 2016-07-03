package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	url := "https://github.com"
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = open_linux(url)
	case "darwin":
		cmd = open_darwin(url)
	case "windows":
		cmd = open_windows(url)
	default:
		log.Fatalf("Cannot open url: %s", "operating system not supported")
	}

	if err := cmd.Run(); err != nil {
		log.Fatal("Open url error: %v", err)
	}
}

func open_linux(url string) *exec.Cmd {
	return exec.Command("xdg-open", url)
}

func open_darwin(url string) *exec.Cmd {
	return exec.Command("open", url)
}

func open_windows(url string) *exec.Cmd {
	cmd := "url.dll,FileProtocolHandler"
	runDll32 := filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
	return exec.Command(runDll32, cmd, url)
}
