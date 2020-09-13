package model

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type UpdatableInterface interface {
	Update(win *pixelgl.Window, dt float64)
}

type DrawableInterface interface {
	Draw(win pixel.Target)
}

type ObjectInterface interface {
	MatObjectInterface
	GetSprite() *pixel.Sprite
}

type StaticObjectInterface interface {
	ObjectInterface
	DrawableInterface
}

type StaticObject struct {
	MatObject
	sprite *pixel.Sprite
}

func NewStaticObject(picture pixel.Picture) *StaticObject {
	return &StaticObject{
		sprite: pixel.NewSprite(picture, picture.Bounds()),
		//MatObject: NewMatObject(pixel.V(0, 0), 0, pixel.ZV, pixel.V(1, 1), pixel.ZV),
	}
}

func (o StaticObject) GetSprite() *pixel.Sprite {
	return o.sprite
}
