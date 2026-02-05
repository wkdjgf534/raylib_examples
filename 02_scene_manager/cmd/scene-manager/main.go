package main

import (
	"scene-manager/internal/config"
	"scene-manager/internal/game"
)

func main() {
	cfg := config.New()
	game := game.New(cfg)
	game.Run()
}
