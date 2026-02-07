package main

import (
	"go-scene-manager/internal/config"
	"go-scene-manager/internal/game"
)

func main() {
	cfg := config.New()
	game := game.New(cfg)
	game.Run()
}
