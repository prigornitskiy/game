package model

import (
	"github.com/faiface/pixel"
)

type Tile struct {
	*StaticObject
	Type int
}

func NewTile(picture pixel.Picture, t int) *Tile {
	return &Tile{
		Type:         t,
		StaticObject: NewStaticObject(picture),
	}
}
