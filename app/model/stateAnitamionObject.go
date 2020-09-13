package model

type StateAnimatedObjectInterface interface {
	Add(state string, a *AnimatedInterface)
}

type StateAnimatedObject struct {
	Items map[string]*AnimatedInterface
	State string
}

func NewStateAnimatedObject() *StateAnimatedObject {
	return &StateAnimatedObject{
		Items: make(map[string]*AnimatedInterface),
	}
}

func (o *StateAnimatedObject) Add(state string, a *AnimatedInterface) {
	o.Items[state] = a
}
