package model

import (
	"github.com/faiface/pixel"
)

type MatObjectInterface interface {
	GetMatrix() pixel.Matrix
	SetPosition(p pixel.Vec) MatObjectInterface
	SetScale(s pixel.Vec) MatObjectInterface
	SetRotate(rotate float64) MatObjectInterface
}

type MatObject struct {
	Position     pixel.Vec
	Rotate       float64
	RotateAround pixel.Vec
	Scale        pixel.Vec
	ScaleAround  pixel.Vec
}

func NewMatObject(position pixel.Vec, rotate float64, rotateAround pixel.Vec, scale pixel.Vec, scaleAround pixel.Vec) *MatObject {
	return &MatObject{
		Position:     position,
		Rotate:       rotate,
		RotateAround: rotateAround,
		Scale:        scale,
		ScaleAround:  scaleAround,
	}
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
