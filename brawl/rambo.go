package brawl

import (
	"github.com/T0PC4T/BossFight/loader"
	"github.com/hajimehoshi/ebiten"
)

type Rambo struct {
	s          *sprite
	id         int
	status     string
	x, y       float64
	vx, vy     float64
	w, h       int
	components []component
}

// getters and setters
func (r *Rambo) getPos() (float64, float64) {
	return r.x, r.y
}
func (r *Rambo) setPos(x, y float64) error {
	r.x, r.y = x, y
	return nil
}

func (r *Rambo) getVel() (float64, float64) {
	return r.vx, r.vy
}
func (r *Rambo) setVel(vx, vy float64) error {
	r.vx, r.vy = vx, vy
	return nil
}
func (r *Rambo) getDim() (int, int) {
	return r.w, r.h
}
func (r *Rambo) setDim(w, h int) error {
	r.w, r.h = w, h
	return nil
}

func (r *Rambo) getComponentSlicePtr() *[]component {
	return &r.components
}

func (r *Rambo) setID(id int) error {
	r.id = id
	return nil
}

func (r *Rambo) update() error {
	return nil
}

func (r *Rambo) draw(screen *ebiten.Image) error {
	canvas, _ := ebiten.NewImage(r.w, r.h, ebiten.FilterDefault)
	r.s.draw(canvas)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.x, r.y)
	screen.DrawImage(canvas, op)
	return nil
}

func NewRambo(l level) {
	// Definitions //
	s := new(sprite)
	s.animations = make(map[string]*animation)

	a := new(animation)
	a.spriteSheet = loader.RamboSpriteSheet
	a.frameOX = 0
	a.frameOY = 32
	a.frameWidth = 32
	a.frameHeight = 32
	a.frameNum = 6
	a.ticksPerFrame = 10

	// Assignments //

	// Sprite
	s.status = "running"
	s.animations["running"] = a

	// Create rambo
	r := &Rambo{s: s, status: "alive", w: 100, h: 100, x: 100, y: 100}

	addElement(l, r)
}
