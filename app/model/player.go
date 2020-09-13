package model

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type PlayerInterface interface {
	AnimatedInterface
	ObjectInterface
	UpdatableInterface
	MovableObjectInterface
	SetTarget(target pixel.Vec)
}

type Player struct {
	*AnimatedObject
	*MovableObject
	Target    pixel.Vec
	JumpSpeed float64
}

func NewPlayer(picture pixel.Picture) *Player {
	return &Player{
		AnimatedObject: NewAnimatedObject(picture, 256, 256),
		MovableObject:  NewMovableDisplayObject(250),
		JumpSpeed:      100,
	}
}

func (p *Player) SetTarget(target pixel.Vec) {
	p.Target = target
}

func (p *Player) Update(win *pixelgl.Window, dt float64) {
	newPosition := p.Move(p.Position, p.Target, dt)
	if newPosition.X < p.Position.X {
		p.Scale = pixel.V(-1, 1)
	} else if newPosition.X > p.Position.X {
		p.Scale = pixel.V(1, 1)
	}

	p.SetPosition(newPosition)
}
