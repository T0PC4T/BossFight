package brawl

import "github.com/hajimehoshi/ebiten"

// Elements

type element struct {
	ID     int
	s      *sprite
	l      *level
	status string
	x, y   float64
	vx, vy float64
	w, h   float64

	components []component
}

// getters and setters
func (e *element) getPos() (float64, float64) {
	return e.x, e.y
}
func (e *element) setPos(x, y float64) error {
	e.x, e.y = x, y
	return nil
}
func (e *element) getTopLeft() (float64, float64) {
	return e.x - e.w/2, e.y - e.h/2
}
func (e *element) getVel() (float64, float64) {
	return e.vx, e.vy
}
func (e *element) setVel(vx, vy float64) error {
	e.vx, e.vy = vx, vy
	return nil
}
func (e *element) getDim() (float64, float64) {
	return e.w, e.h
}
func (e *element) setDim(w, h float64) error {
	e.w, e.h = w, h
	return nil
}

func (e *element) setID(ID int) error {
	e.ID = ID
	return nil
}

// loop functions (update and draw)
func (e *element) update() error {
	for _, c := range e.components {
		c.update(e)
	}
	return nil
}

func (e *element) draw(screen *ebiten.Image) error {
	canvas, _ := ebiten.NewImage(int(e.w), int(e.h), ebiten.FilterDefault)
	e.s.draw(canvas)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-e.w/2, -e.h/2)
	op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(canvas, op)
	return nil
}

func (e *element) addComponent(c component) {
	e.components = append(e.components, c)
}

// func (e *element) removeComponent(name string) bool {
// 	for i, c := range e.components {
// 		if c.getName() == name {

// 		}
// 	}
// 	if index < len(e.components) {
// 		e.components[ID] = e.components[len(e.components)-1]
// 		e.components = e.components[:len(e.components)-1]
// 	} else {
// 		return false
// 	}
// 	return true
// }
