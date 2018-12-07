package brawl

import (
	"github.com/T0PC4T/BossFight/loader"
)

const (
	gravity     float64 = 0.3
	ramboSpeed  float64 = 0.4
	ramboMaxVel float64 = 10
)

type rambo struct {
	e         *element
	speed     float64
	isRunning bool
	canJump   bool
	cooldown  int
}

func (r *rambo) right() {
	r.e.vx += r.speed
}

func (r *rambo) left() {
	r.e.vx -= r.speed
}

func (r *rambo) up() {
	if r.canJump {
		r.canJump = false
		r.e.vy -= 10
	}
}

func (r *rambo) down() {
	r.e.vy += 1
}

func (r *rambo) upRight() {
	r.e.vx += r.speed
	if r.canJump {
		r.canJump = false
		r.e.vy -= 10
	}
}

func (r *rambo) upLeft() {
	r.e.vx -= r.speed
	if r.canJump {
		r.canJump = false
		r.e.vy -= 10
	}
}

func (r *rambo) downRight() {
	r.e.vx += r.speed
	r.e.vy += 1
}

func (r *rambo) downLeft() {
	r.e.vx -= r.speed
	r.e.vy += 1
}

func (r *rambo) ground() {
	r.canJump = true
}

func NewRambo(l *level) {
	// Definitions //
	s := new(sprite)
	s.animations = make(map[string]*animation)
	s.status = "idle"

	var a *animation
	// idle
	a = new(animation)
	a.spriteSheet = loader.RamboSpriteSheet
	a.frameOX = 0
	a.frameOY = 0
	a.frameWidth = 32
	a.frameHeight = 32
	a.frameNum = 4
	a.ticksPerFrame = 5
	s.animations["idle"] = a

	// running
	a = new(animation)
	a.spriteSheet = loader.RamboSpriteSheet
	a.frameOX = 0
	a.frameOY = 32
	a.frameWidth = 32
	a.frameHeight = 32
	a.frameNum = 7
	a.ticksPerFrame = 5
	s.animations["running"] = a

	// Create rambo
	// Create element
	e := &element{s: s, status: "alive", w: 32, h: 32, x: 100, y: 100}
	// Create rambo
	r := &rambo{e: e, speed: ramboSpeed}
	e.newPusherComponent(0, gravity)
	e.newKeyboardController(r.right, r.left, r.up, r.down, r.upRight, r.upLeft, r.downRight, r.downLeft, r.right, r.right, r.right, r.right, r.right)
	e.newVelocityApplier(ramboMaxVel, 1, 0.9)
	e.newBlockCollider()
	e.newAnimator(r.ground)
	l.addElement(e)
}
