package model

import (
	"game/app/domain"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type GameInterface interface {
	Init()
	Draw(win *pixelgl.Window, dt float64)
}

type Game struct {
	pictureRepository domain.PictureRepositoryInterface
	player            PlayerInterface
	GameLayer         LayerInterface
	UILayer           LayerInterface
}

func NewGame(pr domain.PictureRepositoryInterface) *Game {
	return &Game{
		pictureRepository: pr,
		GameLayer:         NewLayer(),
		UILayer:           NewLayer(),
	}
}

func (g *Game) Init() {
	p := NewPlayer(g.pictureRepository.GetPicture("player"))
	p.SetPosition(pixel.Vec{X: 300, Y: 300}).SetScale(pixel.Vec{X: .5, Y: .5})
	p.SetVelocity(250)
	p.AnimationSpeed = .1
	p.Target = pixel.Vec{X: 1000, Y: 300}
	g.player = p
	g.GameLayer.AddItem(g.player)
}

func (g *Game) Draw(win *pixelgl.Window, dt float64) {
	win.Clear(colornames.Black)

	if win.JustPressed(pixelgl.KeyLeft) {
		g.player.SetTarget(pixel.Vec{X: 0, Y: 300})
	}

	if win.JustPressed(pixelgl.KeyRight) {
		g.player.SetTarget(pixel.Vec{X: 1024, Y: 300})
	}

	g.player.Update(win, dt)
	g.GameLayer.Draw(win)
}
