package ui

import (
	"game/app/model"
	"github.com/faiface/pixel"
)

type ButtonInterface interface {
	model.StaticObjectInterface
}

type Button struct {
	*model.StaticObject
	Text string
}

func NewButton(picture domain.PictureInterface, text string) *Button {
	return &Button{
		StaticObject: model.NewStaticObject(picture),
		Text:         text,
	}
}
