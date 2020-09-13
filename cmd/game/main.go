package main

import (
	"fmt"
	"game/app/domain"
	"game/app/model"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"time"
)

import _ "image/png"

func run() {
	rand.Seed(time.Now().UnixNano())
	cfg := pixelgl.WindowConfig{
		Title:  "Clean!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)
	win.SetVSync(true)
	pr := domain.NewPictureRepository()

	loader := domain.NewLoader(pr)
	err = loader.LoadPicture("player", "img/spritestrip.png")

	if err != nil {
		panic(err)
	}

	err = loader.LoadPicture("tiles", "img/tileset.png")

	if err != nil {
		panic(err)
	}

	err = loader.LoadPicture("boom", "img/boom.png")

	if err != nil {
		panic(err)
	}

	g := model.NewGame(pr)
	g.Init()
	last := time.Now()

	frames := 0
	second := time.Tick(time.Second)
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		g.Draw(win, dt)
		frames++
		win.Update()
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
