package main

import (
	"log"

	"github.com/PauloFH/A-Star/data"
	"github.com/PauloFH/A-Star/game"
	"github.com/PauloFH/A-Star/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	data.InitData()
	ui.LoadFonts()

	ebiten.SetWindowSize(900, 700)
	ebiten.SetWindowTitle("A* Pathfinding")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
