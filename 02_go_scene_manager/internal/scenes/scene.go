package scenes

import "go-scene-manager/internal/config"

type Scene interface {
	Init()
	Update() string
	Draw(cfg *config.Config)
	Unload()
}
