package main

import (
	"log"

	"astralwm/internal/wm"
)

func main() {
	manager, err := wm.New()
	if err != nil {
		log.Fatalf("AstralWM: %v", err)
	}

	if err := manager.Run(); err != nil {
		log.Fatalf("AstralWM: %v", err)
	}
}
