package main

import (
	"log"
	pkg "snakeGame/pkg"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := pkg.NewGame()

	ebiten.SetWindowSize(pkg.ScreenWidth, pkg.ScreenHeight)
	ebiten.SetWindowTitle("Jeu du serpent")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
