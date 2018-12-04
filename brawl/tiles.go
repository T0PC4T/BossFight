package brawl

import (
	"github.com/T0PC4T/BossFight/loader"
	"github.com/hajimehoshi/ebiten"
)

type tile struct {
	x           int
	y           int
	tileType    string
	tileUpdate  func(*tile)
	tileDraw    func(*tile, *ebiten.Image)
	tileCollide func(*tile, element)
	active      bool
}

func (t *tile) isActive() bool {
	if t == nil {
		return false
	} else {
		return t.active
	}
}

func collisionTileUpdate(t *tile) {}

func collisionTileDraw(t *tile, screen *ebiten.Image) {
	drawImg := loader.Grasstile
	w, _ := drawImg.Size()
	scalePercentage := tileSizeF / float64(w)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scalePercentage, scalePercentage)
	op.GeoM.Translate(float64(t.x*tileSize), float64(t.y*tileSize))
	screen.DrawImage(drawImg, op)
}
func collisionTileCollide(t *tile, e element) {}

func newTile(x, y int) *tile {
	return &tile{x: x,
		y:           y,
		tileUpdate:  collisionTileUpdate,
		tileDraw:    collisionTileDraw,
		tileCollide: collisionTileCollide,
		active:      true}
}

func newBlankTile(x, y int) *tile {
	return &tile{x: x, y: y, tileType: "blank"}
}

func newTileEx(x, y int, tileType string,
	tUpdate func(*tile),
	tDraw func(*tile, *ebiten.Image),
	tCollide func(*tile, element)) *tile {

	return &tile{x, y, "col", tUpdate, tDraw, tCollide, true}
}
