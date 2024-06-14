package pkg

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	rows     int
	cols     int
	Apple    *Apple
	snake    *Snake
	points   int
	gameOver bool
	timer    time.Time
}

func NewBoard(rows int, cols int) *Board {
	rand.Seed(time.Now().UnixNano())

	board := &Board{
		rows:  rows,
		cols:  cols,
		timer: time.Now(),
	}
	// start in top-left corner
	board.snake = NewSnake([]Coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, ebiten.KeyArrowRight)
	board.placeApple()

	return board
}

func (b *Board) Update(input *Input) error {
	if b.gameOver {
		err := playAudio("medias/fail.mp3")
		if err != nil {
			return err
		}
		return nil
	}

	// The more points, the shorter the interval, and the faster the snake moves.
	baseInterval := time.Millisecond * 200
	speedIncreaseFactor := time.Millisecond * 10
	interval := baseInterval - time.Duration(b.points)*speedIncreaseFactor

	// Ensure there is a minimum interval to prevent the game from becoming impossibly fast.
	if interval < time.Millisecond*50 {
		interval = time.Millisecond * 50
	}

	// Process input to change direction.
	if newDir, ok := input.Dir(); ok {
		b.snake.ChangeDirection(newDir)
	}

	// Move snake if enough time has passed.
	if time.Since(b.timer) >= interval {
		if err := b.moveSnake(); err != nil {
			return err
		}
		b.timer = time.Now()
	}

	return nil
}

func (b *Board) placeApple() {
	var x, y int

	for {
		x = rand.Intn(b.cols)
		y = rand.Intn(b.rows)

		// make sure we don't put a Apple on a snake
		if !b.snake.HeadHits(x, y) {
			break
		}
	}

	b.Apple = NewApple(x, y)
}

func (b *Board) moveSnake() error {
	// remove tail first, add 1 in front
	b.snake.Move()

	if b.snakeLeftBoard() || b.snake.HeadHitsBody() {
		b.gameOver = true
		return nil
	}

	if b.snake.HeadHits(b.Apple.x, b.Apple.y) {
		// the snake grows on the next move
		b.snake.justAte = true

		b.placeApple()
		b.points++
	}

	return nil
}

func (b *Board) snakeLeftBoard() bool {
	head := b.snake.Head()
	return head.x > b.cols-1 || head.y > b.rows-1 || head.x < 0 || head.y < 0
}
