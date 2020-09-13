package game

import (
	"github.com/faiface/pixel"
	"math"
)

func Distance(startVec pixel.Vec, endVec pixel.Vec) float64 {
	return math.Sqrt(math.Pow(endVec.X-startVec.X, 2) + math.Pow(endVec.Y-startVec.Y, 2))
}
