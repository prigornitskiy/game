package model

import "github.com/faiface/pixel"

type LayerInterface interface {
	DrawableInterface
	AddItem(item ObjectInterface)
}
type Layer struct {
	Items []ObjectInterface
}

func NewLayer() *Layer {
	return &Layer{}
}

func (o *Layer) AddItem(item ObjectInterface) {
	o.Items = append(o.Items, item)
}

func (o *Layer) Draw(win pixel.Target) {
	for _, s := range o.Items {
		s.GetSprite().Draw(win, s.GetMatrix())
	}
}
