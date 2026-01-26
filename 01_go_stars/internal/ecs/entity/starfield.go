package entity

type StarField struct {
	Stars []*Star
}

func NewStarField(starCount uint16, width, height float32) *StarField {
	stars := make([]*Star, starCount)
	for i := range stars {
		star := NewStar(width, height)
		star.Color = star.starColor()
		stars[i] = star
	}

	return &StarField{
		Stars: stars,
	}
}
