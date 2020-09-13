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
	p.Position.X = 300
	p.Position.Y = 300
	p.AnimationSpeed = .1
	p.Target = pixel.Vec{X: 1000, Y: 300}
	g.player = p

	//tiles:=NewSpriteList(g.pictureRepository.GetPicture("tiles"), 64, 64)

	//g.GameLayer = append(g.player, NewTile(g.pictureRepository.GetPicture("tiles"), 0))
}

func (g *Game) Draw(win *pixelgl.Window, dt float64) {
	win.Clear(colornames.Skyblue)

	if win.JustPressed(pixelgl.KeyLeft) {
		g.player.SetTarget(pixel.Vec{X: 0, Y: 300})
	}

	if win.JustPressed(pixelgl.KeyRight) {
		g.player.SetTarget(pixel.Vec{X: 1024, Y: 300})
	}

	/*
		for _, s := range g.staticObjects {
			s.GetSprite().Draw(win, s.GetMatrix())
		}*/
	//g.player.Update(win,dt)
	g.player.Animate(dt)
	g.player.Update(win, dt)
	g.player.GetSprite().Draw(win, g.player.GetMatrix())
}
