package model

import "github.com/faiface/pixel"

type LayerInterface interface {
	DrawableInterface
}
type Layer struct {
	Items []ObjectInterface
}

func NewLayer() *Layer {
	return &Layer{}
}

func (o Layer) Draw(win pixel.Target) {
	for _, s := range o.Items {
		s.GetSprite().Draw(win, s.GetMatrix())
	}
}
