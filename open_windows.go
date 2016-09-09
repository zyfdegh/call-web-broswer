package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func open(url string) *exec.Cmd {
	cmd := "url.dll,FileProtocolHandler"
	runDll32 := filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
	return exec.Command(runDll32, cmd, url)
}
