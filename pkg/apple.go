package pkg

type Apple struct {
	x, y int
}

func NewApple(x, y int) *Apple {
	return &Apple{
		x: x,
		y: y,
	}
}
