package brawl

import "github.com/T0PC4T/BossFight/loader"

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
	e := &element{s: s, status: "alive", w: 100, h: 100, x: 100, y: 100}
	e.newGravityComponent(0.25)
	e.newBlockCollider()
	e.newVelocityApplier(2, 1)
	l.addElement(e)
}
