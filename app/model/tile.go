package model

import "game/app/domain"

type Tile struct {
	*StaticObject
	Type int
}

func NewTile(picture domain.PictureInterface, t int) *Tile {
	return &Tile{
		Type:         t,
		StaticObject: NewStaticObject(picture),
	}
}
