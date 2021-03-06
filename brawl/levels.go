package brawl

import (
	"github.com/T0PC4T/BossFight/loader"
	"github.com/hajimehoshi/ebiten"
)

type level struct {
	name        string
	mapLayout   [][]*tile
	allElements []*element
}

// Update is the loop function for devLevel
func (l *level) update() (bool, error) {
	return false, nil
}

// Draw is the draw function for devLevel
func (l *level) draw(canvas *ebiten.Image) error {
	// Updating

	// updateTiles

	// TODO

	// Update Elements
	for _, e := range l.allElements {
		e.update()
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw Background
	op := &ebiten.DrawImageOptions{}
	w, h := loader.BackgroundImg.Size()
	scaleW := float64(screenWidth) / float64(w)
	scaleH := float64(screenHeight) / float64(h)
	op.GeoM.Scale(scaleW, scaleH)
	canvas.DrawImage(loader.BackgroundImg, op)

	// Draw Tiles
	for _, col := range l.mapLayout {
		for _, curTile := range col {
			if curTile != nil {
				curTile.tileDraw(canvas)
			}
		}
	}

	// Draw Elements

	for _, e := range l.allElements {
		e.draw(canvas)
	}

	return nil
}

func (l *level) addElement(e *element) {
	e.l = l
	l.allElements = append(l.allElements, e)
}

func (l *level) removeElement(ID int) bool {
	if ID < len(l.allElements) {
		l.allElements[ID] = l.allElements[len(l.allElements)-1]
		l.allElements = l.allElements[:len(l.allElements)-1]
		l.allElements[ID].setID(ID)
	} else {
		return false
	}
	return true
}

// Levels

// newDevelopmentLevel Creates a new level in this case its a development level
func newDevelopmentLevel() *level {
	l := &level{name: "DevLevel", allElements: make([]*element, 0, 50)}
	l.mapLayout = make([][]*tile, tilesWidth)
	for tx := range l.mapLayout {
		l.mapLayout[tx] = make([]*tile, tilesHeight)
		for ty := range l.mapLayout[tx] {
			if ty == 26 {
				l.newTile(tx, ty)
			}
		}
	}
	l.newTile(3, 25)
	l.newTile(4, 25)
	l.newTile(5, 25)
	l.newTile(6, 25)
	l.newTile(7, 25)

	l.newTile(13, 20)
	l.newTile(14, 20)
	l.newTile(15, 20)

	return l
}
