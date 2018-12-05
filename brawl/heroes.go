package brawl

import (
	"github.com/T0PC4T/BossFight/loader"
)

const (
	gravity     float64 = 0.3
	ramboSpeed  float64 = 0.2
	ramboMaxVel float64 = 3
)

type rambo struct {
	e        *element
	speed    float64
	cooldown int
}

func (r *rambo) right() {
	r.e.vx += r.speed
}

func (r *rambo) left() {
	r.e.vx -= r.speed
}

func (r *rambo) up() {
	r.e.vy -= 4
}

func (r *rambo) down() {
	r.e.vy += 1
}

func (r *rambo) upRight() {
	r.e.vx += r.speed
	r.e.vy -= 4
}

func (r *rambo) upLeft() {
	r.e.vx -= r.speed
	r.e.vy -= 4
}

func (r *rambo) downRight() {
	r.e.vx += r.speed
	r.e.vy += 1
}

func (r *rambo) downLeft() {
	r.e.vx -= r.speed
	r.e.vy += 1
}

func NewRambo(l *level) {
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
	// Create element
	e := &element{s: s, status: "alive", w: 32, h: 32, x: 100, y: 100}
	// Create rambo
	r := &rambo{e: e, speed: ramboSpeed}
	e.newGravityComponent(gravity)
	e.newKeyboardController(r.right, r.left, r.up, r.down, r.upRight, r.upLeft, r.downRight, r.downLeft, r.right, r.right, r.right, r.right, r.right)
	e.newVelocityApplier(ramboMaxVel, 1, 0.9)
	e.newBlockCollider()
	l.addElement(e)
}
