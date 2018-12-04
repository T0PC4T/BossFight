package brawl

import "github.com/hajimehoshi/ebiten"

// Elements

type element struct {
	ID         int
	s          *sprite
	status     string
	x, y       float64
	vx, vy     float64
	w, h       int
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

func (e *element) getVel() (float64, float64) {
	return e.vx, e.vy
}
func (e *element) setVel(vx, vy float64) error {
	e.vx, e.vy = vx, vy
	return nil
}
func (e *element) getDim() (int, int) {
	return e.w, e.h
}
func (e *element) setDim(w, h int) error {
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
	canvas, _ := ebiten.NewImage(e.w, e.h, ebiten.FilterDefault)
	e.s.draw(canvas)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(canvas, op)
	return nil
}

func (e *element) addComponent(c component) {
	c.setID(len(e.components) - 1)
	e.components = append(e.components, c)
}

func (e *element) removeComponent(ID int) bool {
	if ID < len(e.components) {
		e.components[ID] = e.components[len(e.components)-1]
		e.components = e.components[:len(e.components)-1]
		e.components[ID].setID(ID)
	} else {
		return false
	}
	return true
}
