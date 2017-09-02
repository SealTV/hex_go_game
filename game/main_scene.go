package game

import (
	"github.com/SealTV/hex_go_game/resources"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type Game struct {
	Settings     resources.Settings
	Localization resources.Localizator
	Sprites      resources.SpriteMap
	Spritesheet  resources.SpritesheetData
}

func (g *Game) Run() {
	cfg := pixelgl.WindowConfig{
		Title:  g.Localization.GetString(g.Settings.AppName),
		Bounds: pixel.R(0, 0, g.Settings.WindowW, g.Settings.WindowH),
		VSync:  g.Settings.IsUseVSync,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)
	pic := g.Spritesheet.GetSprite("waveWater.png")
	sprites = append(sprites, pic)
	sprites = append(sprites, pic)

	batch = pixel.NewBatch(&pixel.TrianglesData{}, g.Spritesheet.Picture)
	for !win.Closed() {
		g.update(win)
		win.Update()
	}
}

var angle float64
var sprites []*pixel.Sprite
var batch *pixel.Batch

func (g *Game) update(win *pixelgl.Window) {
	angle += 0.05
	win.Clear(colornames.Black)
	batch.Clear()
	for i, s := range sprites {
		mat := pixel.IM
		var a float64
		if i%2 == 0 {
			a = angle
		} else {
			a = -angle
		}
		mat = mat.Rotated(pixel.ZV, a)
		pos := win.Bounds().Center()
		pos.X = pos.X + float64(i)*30
		mat = mat.Moved(pos)

		s.Draw(batch, mat)
	}
	batch.Draw(win)
}
