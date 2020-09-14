package models

/*
Старый код: нужный в app/model/

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
)

type TileMap struct {
	batch      *pixel.Batch
	items      [][]*pixel.Sprite
	picture    pixel.Picture
	tileFrames []pixel.Rect
}

func NewTileMap(picture pixel.Picture) *TileMap {
	var tileFrames []pixel.Rect
	for x := picture.Bounds().Min.X; x < picture.Bounds().Max.X; x += 64 {
		for y := picture.Bounds().Min.Y; y < picture.Bounds().Max.Y; y += 64 {
			tileFrames = append(tileFrames, pixel.R(x, y, x+64, y+64))
		}
	}

	n := 16
	m := 12
	matrix := make([][]*pixel.Sprite, n)
	rows := make([]*pixel.Sprite, n*m)
	for i := 0; i < n; i++ {
		matrix[i] = rows[i*m : (i+1)*m]
	}

	tm := &TileMap{
		items:      matrix,
		picture:    picture,
		tileFrames: tileFrames,
		batch:      pixel.NewBatch(&pixel.TrianglesData{}, picture),
	}
	tm.generateMap()
	return tm
}

func (t *TileMap) generateMap() {
	for x, row := range t.items {
		for y, _ := range row {
			t.items[x][y] = pixel.NewSprite(t.picture, t.tileFrames[2+(rand.Intn(4)*10)])
		}
	}
}

func (t *TileMap) Draw(win *pixelgl.Window, dt float64) {
	t.batch.Clear()

	for x, row := range t.items {
		for y, col := range row {
			mat := pixel.IM
			mat = mat.Moved(pixel.V(float64(x)*64+32, float64(y)*64+32))
			if col != nil {
				col.Draw(t.batch, mat)
			}
		}
	}

	t.batch.Draw(win)
}

*/
