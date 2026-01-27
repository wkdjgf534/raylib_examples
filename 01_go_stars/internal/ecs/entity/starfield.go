package entity

type StarField struct {
	Stars []Star
}

func NewStarField(starCount uint16, width, height float32) *StarField {
	stars := make([]Star, starCount)

	for i := range stars {
		stars[i] = NewStar(width, height)
	}

	return &StarField{
		Stars: stars,
	}
}
