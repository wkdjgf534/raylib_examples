package main

import (
	"go-stars/internal/config"
	"go-stars/internal/game"
)

func main() {
	cfg := config.New()
	game := game.New(cfg)
	game.Run()
}
