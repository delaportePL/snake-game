package pkg

import (
	"math/rand"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	rows                int
	cols                int
	Apple               *Apple
	snake               *Snake
	points              int
	gameOver            bool
	gameOverSoundPlayed bool
	timer               time.Time
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
		if !b.gameOverSoundPlayed {
			err := playAudio("medias/fail.mp3")
			if err != nil {
				return err
			}
			b.gameOverSoundPlayed = true // Mark that the game over sound has been played
		}
		return nil
	}

	// The more points, the shorter the interval, and the faster the snake moves.
	baseInterval := time.Millisecond * 200
	speedIncreaseFactor := time.Millisecond * 10 // This value controls how much faster the game gets per point.
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

func playAudio(filePath string) error {
	// Ouvrir le fichier audio
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Décoder le fichier audio
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		return err
	}
	defer streamer.Close()

	// Configurer le lecteur audio
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}

	// Lecture de l'audio une seule fois
	done := make(chan struct{})
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		os.Exit(1)
		close(done)
	})))

	// Attendre que la lecture soit terminée

	<-done
	speaker.Close()
	// Arrêter le lecteur audio

	return nil
}
