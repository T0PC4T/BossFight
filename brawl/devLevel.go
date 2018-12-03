package brawl

import (
	gl "github.com/T0PC4T/BossFight/global"
	"github.com/hajimehoshi/ebiten"
)

type devLevel struct {
	levelName   string
	mapLayout   [gl.TilesWidth][gl.TilesHeight]*tile
	allElements []element
}

func (dl *devLevel) GetElementSlicePtr() *[]element {
	return &dl.allElements
}

// Update is the loop function for devLevel
func (dl *devLevel) Update() (bool, error) {
	return true, nil
}

// Draw is the draw function for devLevel
func (dl *devLevel) Draw(canvas *ebiten.Image) error {
	// Updating

	// Update Elements
	for _, e := range dl.allElements {
		e.update()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw Tiles
	for _, col := range dl.mapLayout {
		for _, curTile := range col {
			if curTile != nil {
				curTile.tileDraw(curTile, canvas)
			}
		}
	}

	// Draw Elements

	for _, e := range dl.allElements {
		e.draw(canvas)
	}

	return nil
}

// NewDevelopmentLevel Creates a new level in this case its a development level
func NewDevelopmentLevel() level {
	dl := &devLevel{levelName: "DevLevel", allElements: make([]element, 0, 50)}
	for tx, col := range dl.mapLayout {
		for ty := range col {
			if ty == 20 {
				dl.mapLayout[tx][ty] = newTile(tx, ty)
			}
		}
	}
	return dl
}
