package ui

import (
	"game/app/model"
	"github.com/faiface/pixel"
)

type MenuItemsInterface interface {
	model.ObjectInterface
	model.StaticObjectInterface
}

type MenuInterface interface {
	model.MatObjectInterface
}

type Menu struct {
	items []MenuItemsInterface
}

func NewMenu(picture pixel.Picture) *Menu {
	m := &Menu{}
	return m
}
