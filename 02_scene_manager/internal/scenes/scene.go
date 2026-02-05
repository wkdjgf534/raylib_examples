package scenes

type Scene interface {
	Init()
	Update() string
	Draw()
	Unload()
}
