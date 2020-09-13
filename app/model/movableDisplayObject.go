package model

import (
	"game"
	"github.com/faiface/pixel"
)

type MovableObjectInterface interface {
	Move(from pixel.Vec, to pixel.Vec, dt float64) pixel.Vec
	SetVelocity(v float64) MovableObjectInterface
}

type MovableObject struct {
	Velocity float64
}

func (o *MovableObject) SetVelocity(v float64) MovableObjectInterface {
	o.Velocity = v
	return o
}

func (o *MovableObject) Move(from pixel.Vec, to pixel.Vec, dt float64) pixel.Vec {
	d := game.Distance(from, to)

	t := (o.Velocity * dt) / d
	if t < 1 {
		return pixel.Lerp(from, to, t)
	} else {
		return to
	}
}
