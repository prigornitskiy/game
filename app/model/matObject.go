package model

import (
	"github.com/faiface/pixel"
)

type MatObjectInterface interface {
	SetPosition(p pixel.Vec) MatObjectInterface
	SetScale(s pixel.Vec) MatObjectInterface
	SetRotate(rotate float64) MatObjectInterface
	GetMatrix() pixel.Matrix
}

type MatObject struct {
	Position     pixel.Vec
	Rotate       float64
	RotateAround pixel.Vec
	Scale        pixel.Vec
	ScaleAround  pixel.Vec
}

func (o *MatObject) SetPosition(p pixel.Vec) MatObjectInterface {
	o.Position = p
	return o
}

func (o *MatObject) SetScale(s pixel.Vec) MatObjectInterface {
	o.Scale = s
	return o
}

func (o *MatObject) SetRotate(rotate float64) MatObjectInterface {
	o.Rotate = rotate
	return o
}

func (o *MatObject) GetMatrix() pixel.Matrix {
	mat := pixel.IM
	mat = mat.ScaledXY(o.ScaleAround, o.Scale).Rotated(o.RotateAround, o.Rotate).Moved(o.Position)
	return mat
}
