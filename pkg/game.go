package pkg

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 600
	boardRows    = 20
	boardCols    = 20
)

var (
	backgroundColor = color.RGBA{0, 0, 0, 255}
	snakeColor      = color.RGBA{0, 255, 0, 255}
	AppleColor      = color.RGBA{255, 0, 0, 255}
)

type Game struct {
	input *Input
	board *Board
}

func NewGame() *Game {
	return &Game{
		input: NewInput(),
		board: NewBoard(boardRows, boardCols),
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	return g.board.Update(g.input)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	if g.board.gameOver {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Vous avez perdu. Votre score: %d", g.board.points))
	} else {
		width := ScreenHeight / boardRows

		for _, p := range g.board.snake.body {
			ebitenutil.DrawRect(screen, float64(p.y*width), float64(p.x*width), float64(width), float64(width), snakeColor)
		}
		if g.board.Apple != nil {
			ebitenutil.DrawRect(screen, float64(g.board.Apple.y*width), float64(g.board.Apple.x*width), float64(width), float64(width), AppleColor)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("votre score: %d", g.board.points))
	}
}
