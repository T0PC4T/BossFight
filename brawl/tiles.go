package brawl

import (
	"github.com/T0PC4T/BossFight/loader"
	"github.com/hajimehoshi/ebiten"
)

type tile struct {
	x, y        float64
	tileType    string
	tileUpdate  func()
	tileDraw    func(*ebiten.Image)
	tileCollide func(*element)
	active      bool
}

func (t *tile) isActive() bool {
	if t == nil {
		return false
	}
	return t.active
}
func (t *tile) getPos() (float64, float64) {
	return t.x, t.y
}
func (t *tile) getDim() (float64, float64) {
	return tileSizeF, tileSizeF
}

func (t *tile) collisionTileUpdate() {}

func (t *tile) collisionTileDraw(screen *ebiten.Image) {
	drawImg := loader.Grasstile
	w, _ := drawImg.Size()
	scalePercentage := tileSizeF / float64(w)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scalePercentage, scalePercentage)
	op.GeoM.Translate(-tileSizeF/2, -tileSizeF/2)
	op.GeoM.Translate(t.x, t.y)
	screen.DrawImage(drawImg, op)
}
func (t *tile) collisionTileCollide(e *element) {
	if isPointInBox(e.x, e.y-e.h/2, t) {
		e.vy = 0
		e.y = t.y + tileSizeF/2 + e.h/2
	} else if isPointInBox(e.x, e.y+e.h/2, t) {
		e.vy = 0
		e.y = t.y - tileSizeF/2 - e.h/2
	}
	if isPointInBox(e.x-e.w/2, e.y, t) {
		e.vx = 0
		e.x = t.x + tileSizeF/2 + e.w/2
	} else if isPointInBox(e.x+e.w/2, e.y, t) {
		e.vx = 0
		e.x = t.x - tileSizeF/2 - e.w/2
	}
}

func (l *level) newTile(x, y int) {
	ctx := (float64(x) * tileSizeF) + tileSizeF/2
	cty := (float64(y) * tileSizeF) + tileSizeF/2
	t := &tile{x: ctx, y: cty, active: true}

	t.tileUpdate = t.collisionTileUpdate
	t.tileDraw = t.collisionTileDraw
	t.tileCollide = t.collisionTileCollide

	l.mapLayout[x][y] = t
}

type collider interface {
	getPos() (float64, float64)
	getDim() (float64, float64)
}

// Rectangle 1’s bottom edge is less than Rectangle 2’s top edge.
// OR
// Rectangle 1’s top edge is more than Rectangle 2’s bottom edge.
// AND
// Rectangle 1’s left edge is more than of Rectangle 2’s right edge.
// OR
// Rectangle 1’s right edge is less than of Rectangle 2’s left edge.
func isPointInBox(x, y float64, box collider) bool {
	bx, by := box.getPos()
	bw, bh := box.getDim()
	return x <= bx+bw/2 && x >= bx-bw/2 && y <= by+bh/2 && y >= by-bh/2
}

func isColliding(c1, c2 collider) (tr, tl, br, bl bool) {
	c1x, c1y := c1.getPos()
	c1w, c1h := c1.getDim()

	rx := c1x + c1w/2
	lx := c1x - c1w/2

	ty := c1y - c1h/2
	by := c1y + c1h/2

	tr = isPointInBox(rx, ty, c2)
	tl = isPointInBox(lx, ty, c2)
	br = isPointInBox(rx, by, c2)
	bl = isPointInBox(lx, by, c2)

	return
}
