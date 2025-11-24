package game

import (
	"fmt"
	"image/color"

	"github.com/PauloFH/A-Star/astar"
	"github.com/PauloFH/A-Star/data"
	"github.com/PauloFH/A-Star/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type State int

const (
	StateMenu State = iota
	StateMap
)

type Game struct {
	State               State
	FinalPath           []string
	PathCost            int
	CurrentSegmentIndex int
	SegmentProgress     float32
	AnimationSpeed      float32
	DdStart             *ui.Dropdown
	DdTarget            *ui.Dropdown
}

func NewGame() *Game {
	return &Game{
		State:          StateMenu,
		AnimationSpeed: 0.02,
		DdStart: &ui.Dropdown{
			X: 50, Y: 120, W: 250, H: 30,
			Label: "Origem (Arad)", Options: data.SortedCities, Selected: "Arad",
		},
		DdTarget: &ui.Dropdown{
			X: 600, Y: 120, W: 250, H: 30,
			Label: "Destino (Bucharest)", Options: data.SortedCities, Selected: "Bucharest",
		},
	}
}

func (g *Game) Update() error {
	if g.State == StateMenu {
		g.updateMenu()
	} else {
		g.updateMap()
	}
	return nil
}

func (g *Game) updateMenu() {
	if g.DdStart.IsOpen {
		g.DdStart.Update()
		return
	}
	if g.DdTarget.IsOpen {
		g.DdTarget.Update()
		return
	}

	if g.DdStart.Update() {
		return
	}
	if g.DdTarget.Update() {
		return
	}

	if isClicked(300, 500, 300, 50) {
		if g.DdStart.Selected != "" && g.DdTarget.Selected != "" && g.DdStart.Selected != g.DdTarget.Selected {
			g.FinalPath, g.PathCost = astar.FindPath(g.DdStart.Selected, g.DdTarget.Selected)
			g.CurrentSegmentIndex = 0
			g.SegmentProgress = 0.0
			g.State = StateMap
		}
	}
}

func (g *Game) updateMap() {
	if isClicked(10, 10, 100, 30) {
		g.State = StateMenu
		return
	}
	if len(g.FinalPath) <= 1 {
		return
	}
	if g.CurrentSegmentIndex < len(g.FinalPath)-1 {
		g.SegmentProgress += g.AnimationSpeed
		if g.SegmentProgress >= 1.0 {
			g.SegmentProgress = 0
			g.CurrentSegmentIndex++
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State == StateMenu {
		g.drawMenu(screen)
	} else {
		g.drawMap(screen)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 40, G: 40, B: 50, A: 255})

	text.Draw(screen, "Distâncias com A*", ui.TitleFace, 250, 60, color.White)
	text.Draw(screen, "Origem:", ui.MediumFace, 110, 100, color.RGBA{R: 150, G: 255, B: 150, A: 255})
	text.Draw(screen, "Destino:", ui.MediumFace, 680, 100, color.RGBA{R: 255, G: 150, B: 150, A: 255})

	btnColor := color.RGBA{G: 100, B: 255, A: 255}
	if g.DdStart.Selected == "" || g.DdTarget.Selected == "" || g.DdStart.Selected == g.DdTarget.Selected {
		btnColor = color.RGBA{R: 100, G: 100, B: 100, A: 255}
	}
	vector.FillRect(screen, 300, 500, 300, 50, btnColor, true)
	text.Draw(screen, "VISUALIZAR ROTA", ui.MediumFace, 350, 530, color.White)

	if g.DdStart.IsOpen {
		g.DdTarget.Draw(screen)
		g.DdStart.Draw(screen)
	} else {
		g.DdStart.Draw(screen)
		g.DdTarget.Draw(screen)
	}
}

func (g *Game) drawMap(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 30, G: 30, B: 30, A: 255})

	vector.FillRect(screen, 10, 10, 80, 25, color.RGBA{R: 200, G: 50, B: 50, A: 255}, true)
	text.Draw(screen, "< VOLTAR", ui.SmallFace, 15, 27, color.White)
	for city, edges := range data.Graph {
		startPos := data.Positions[city]
		for _, edge := range edges {
			endPos := data.Positions[edge.To]
			vector.StrokeLine(screen, startPos.X, startPos.Y, endPos.X, endPos.Y, 2, color.RGBA{R: 80, G: 80, B: 80, A: 255}, true)
		}
	}
	if len(g.FinalPath) > 1 {
		pathColor := color.RGBA{G: 255, B: 100, A: 255}
		for i := 0; i < g.CurrentSegmentIndex; i++ {
			u := g.FinalPath[i]
			v := g.FinalPath[i+1]
			p1 := data.Positions[u]
			p2 := data.Positions[v]
			vector.StrokeLine(screen, p1.X, p1.Y, p2.X, p2.Y, 5, pathColor, true)
		}
		if g.CurrentSegmentIndex < len(g.FinalPath)-1 {
			u := g.FinalPath[g.CurrentSegmentIndex]
			v := g.FinalPath[g.CurrentSegmentIndex+1]
			startPos := data.Positions[u]
			endPos := data.Positions[v]
			currX := startPos.X + (endPos.X-startPos.X)*g.SegmentProgress
			currY := startPos.Y + (endPos.Y-startPos.Y)*g.SegmentProgress
			vector.StrokeLine(screen, startPos.X, startPos.Y, currX, currY, 5, pathColor, true)
			vector.FillCircle(screen, currX, currY, 6, color.RGBA{R: 150, G: 255, B: 150, A: 255}, true)
		}
	}

	for city, edges := range data.Graph {
		startPos := data.Positions[city]
		for _, edge := range edges {
			endPos := data.Positions[edge.To]
			if city < edge.To {
				midX := (startPos.X + endPos.X) / 2
				midY := (startPos.Y + endPos.Y) / 2

				vector.FillRect(screen, midX-10, midY-6, 20, 12, color.RGBA{R: 30, G: 30, B: 30, A: 255}, true)
				text.Draw(screen, fmt.Sprintf("%d", edge.Cost), ui.SmallFace, int(midX)-8, int(midY)+4, color.RGBA{R: 255, G: 255, A: 255})
			}
		}
	}

	for city, pos := range data.Positions {
		col := color.RGBA{R: 200, G: 200, B: 200, A: 255}
		if city == g.DdStart.Selected {
			col = color.RGBA{R: 100, G: 100, B: 255, A: 255}
		} else if city == g.DdTarget.Selected {
			col = color.RGBA{R: 255, G: 100, B: 100, A: 255}
		}

		for i := 0; i <= g.CurrentSegmentIndex && i < len(g.FinalPath); i++ {
			if g.FinalPath[i] == city {
				col = color.RGBA{G: 255, B: 100, A: 255}
			}
		}

		vector.FillCircle(screen, pos.X, pos.Y, 8, col, true)
		text.Draw(screen, city, ui.SmallFace, int(pos.X)-15, int(pos.Y)-12, color.White)
	}

	msg := fmt.Sprintf("Rota: %s -> %s | Custo: %d km", g.DdStart.Selected, g.DdTarget.Selected, g.PathCost)
	text.Draw(screen, msg, ui.MediumFace, 200, 680, color.White)
}

func (g *Game) Layout(w, h int) (int, int) { return 900, 700 }

func isClicked(x, y, w, h int) bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		return mx >= x && mx <= x+w && my >= y && my <= y+h
	}
	return false
}
