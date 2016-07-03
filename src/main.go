package main

import (
	"log"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	if err := open.Run("https://github.com"); err != nil {
		log.Fatal("Open url error: %v", err)
	}
}
