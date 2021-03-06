package model

import (
	"game/app/domain"
	"github.com/faiface/pixel"
)

type AnimatedInterface interface {
	ObjectInterface
	Animate(dt float64)
}

type AnimatedObject struct {
	MatObject
	Sprites            []*pixel.Sprite
	CurrentFrame       int
	AnimationSpeed     float64
	SinceLastAnimation float64
}

func NewAnimatedObject(picture domain.PictureInterface, sizeX float64, sizeY float64) *AnimatedObject {
	var frames []pixel.Rect
	for x := picture.Bounds().Min.X; x < picture.Bounds().Max.X; x += sizeX {
		for y := picture.Bounds().Min.Y; y < picture.Bounds().Max.Y; y += sizeY {
			frames = append(frames, pixel.R(x, y, x+sizeX, y+sizeY))
		}
	}

	sprites := make([]*pixel.Sprite, len(frames), len(frames))

	for i := 0; i < len(frames); i++ {
		sprites[i] = pixel.NewSprite(picture, frames[i])
	}

	return &AnimatedObject{
		Sprites: sprites,
	}
}

func (o *AnimatedObject) Animate(dt float64) {
	o.SinceLastAnimation = o.SinceLastAnimation + dt

	if o.SinceLastAnimation > o.AnimationSpeed {
		o.SinceLastAnimation = 0
		o.CurrentFrame++
		if o.CurrentFrame >= len(o.Sprites) {
			o.CurrentFrame = 0
		}
	}
}

func (o *AnimatedObject) GetSprite() *pixel.Sprite {
	return o.Sprites[o.CurrentFrame]
}
