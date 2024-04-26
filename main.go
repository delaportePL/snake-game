package main

import (
	"log"
	snake "snakeGame/pkg"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := snake.NewGame()

	ebiten.SetWindowSize(snake.ScreenWidth, snake.ScreenHeight)
	ebiten.SetWindowTitle("Jeu du serpent")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
