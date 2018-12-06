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
	if !r.isRunning {
		r.e.s.UpdateStatus("running")
	}
	r.e.vx += r.speed
}

func (r *rambo) left() {
	if !r.isRunning {
		r.e.s.UpdateStatus("running")
	}
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
	r.e.s.UpdateStatus("idle")
}

func NewRambo(l *level) {
	// Definitions //
	s := new(sprite)
	s.animations = make(map[string]*animation)
	s.status = "idle"

	{
		// idle
		a := new(animation)
		a.spriteSheet = loader.RamboSpriteSheet
		a.frameOX = 0
		a.frameOY = 0
		a.frameWidth = 32
		a.frameHeight = 32
		a.frameNum = 5
		a.ticksPerFrame = 5
		s.animations["idle"] = a
	}

	{
		// running
		a := new(animation)
		a.spriteSheet = loader.RamboSpriteSheet
		a.frameOX = 0
		a.frameOY = 32
		a.frameWidth = 32
		a.frameHeight = 32
		a.frameNum = 8
		a.ticksPerFrame = 5
		s.animations["running"] = a
	}

	// Create rambo
	// Create element
	e := &element{s: s, status: "alive", w: 32, h: 32, x: 100, y: 100}
	// Create rambo
	r := &rambo{e: e, speed: ramboSpeed}
	e.newGravityComponent(gravity)
	e.newKeyboardController(r.right, r.left, r.up, r.down, r.upRight, r.upLeft, r.downRight, r.downLeft, r.right, r.right, r.right, r.right, r.right)
	e.newVelocityApplier(ramboMaxVel, 1, 0.9)
	e.newBlockCollider()
	e.newGrounder(r.ground)
	l.addElement(e)
}
