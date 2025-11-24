package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Dropdown struct {
	X, Y, W, H int
	Label      string
	Options    []string
	Selected   string
	IsOpen     bool
}

func (d *Dropdown) Update() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if mx >= d.X && mx <= d.X+d.W && my >= d.Y && my <= d.Y+d.H {
			d.IsOpen = !d.IsOpen
			return true
		}
		if d.IsOpen {
			optionHeight := 25
			listHeight := len(d.Options) * optionHeight
			if mx >= d.X && mx <= d.X+d.W && my > d.Y+d.H && my <= d.Y+d.H+listHeight {
				clickedIndex := (my - (d.Y + d.H)) / optionHeight
				if clickedIndex >= 0 && clickedIndex < len(d.Options) {
					d.Selected = d.Options[clickedIndex]
					d.IsOpen = false
					return true
				}
			} else {
				d.IsOpen = false
			}
		}
	}
	return false
}

func (d *Dropdown) Draw(screen *ebiten.Image) {
	bgColor := color.RGBA{R: 60, G: 60, B: 70, A: 255}
	borderColor := color.RGBA{R: 200, G: 200, B: 200, A: 255}
	textColor := color.White

	vector.FillRect(screen, float32(d.X), float32(d.Y), float32(d.W), float32(d.H), bgColor, true)
	vector.StrokeRect(screen, float32(d.X), float32(d.Y), float32(d.W), float32(d.H), 1, borderColor, true)

	displayTxt := d.Selected
	if displayTxt == "" {
		displayTxt = d.Label
	}
	text.Draw(screen, displayTxt, basicfont.Face7x13, d.X+10, d.Y+20, textColor)

	arrow := "v"
	if d.IsOpen {
		arrow = "^"
	}
	text.Draw(screen, arrow, basicfont.Face7x13, d.X+d.W-20, d.Y+20, color.RGBA{R: 150, G: 150, B: 150, A: 255})
	if d.IsOpen {
		optionHeight := 25
		totalH := len(d.Options) * optionHeight
		vector.FillRect(screen, float32(d.X), float32(d.Y+d.H), float32(d.W), float32(totalH), color.RGBA{R: 50, G: 50, B: 60, A: 255}, true)
		vector.StrokeRect(screen, float32(d.X), float32(d.Y+d.H), float32(d.W), float32(totalH), 1, borderColor, true)

		for i, opt := range d.Options {
			yPos := d.Y + d.H + (i * optionHeight)
			mx, my := ebiten.CursorPosition()
			if mx >= d.X && mx <= d.X+d.W && my >= yPos && my < yPos+optionHeight {
				vector.FillRect(screen, float32(d.X+1), float32(yPos), float32(d.W-2), float32(optionHeight), color.RGBA{R: 80, G: 80, B: 100, A: 255}, true)
			}

			var optColor color.Color = textColor
			if opt == d.Selected {
				optColor = color.RGBA{R: 100, G: 255, B: 100, A: 255}
			}
			text.Draw(screen, opt, basicfont.Face7x13, d.X+10, yPos+18, optColor)
		}
	}
}
