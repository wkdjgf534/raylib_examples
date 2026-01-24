package config

const (
	windowWidth  = 1280
	widnowHeight = 720
	title        = "Stars"
	fps          = 60
)

type Config struct {
	WindowWidth  int32
	WidnowHeight int32
	Title        string
	FPS          int32
}

func New() *Config {
	return &Config{
		WindowWidth:  windowWidth,
		WidnowHeight: widnowHeight,
		Title:        title,
		FPS:          fps,
	}
}
