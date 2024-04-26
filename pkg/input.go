package pkg

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func NewInput() Input {
	return &Input{}
}

func (iInput) Dir() (ebiten.Key, bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return ebiten.KeyArrowUp, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return ebiten.KeyArrowLeft, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return ebiten.KeyArrowRight, true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		return ebiten.KeyArrowDown, true
	}

	return 0, false
}
