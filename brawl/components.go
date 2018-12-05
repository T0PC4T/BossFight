package brawl

import (
	"github.com/hajimehoshi/ebiten"
)

// Component definition

type component interface {
	update(*element) error
	getName() string
	isActive() bool
	setActive(a bool)
}

// Components

type componentActive struct {
	active bool
}

func (c *componentActive) isActive() bool {
	return c.active
}

func (c *componentActive) setActive(a bool) {
	c.active = a
}

// Gravity Componenet

type compGravity struct {
	componentActive
	gravityAmount float64
}

func (g *compGravity) update(e *element) error {
	vx, vy := e.getVel()
	e.setVel(vx, vy+g.gravityAmount)
	return nil
}

func (g *compGravity) getName() string {
	return "gravity"
}

func (e *element) newGravityComponent(gravityAmount float64) {
	cg := &compGravity{gravityAmount: gravityAmount}
	e.addComponent(cg)
}

// Velocity Applier

type compVelocity struct {
	componentActive
	max, scale, friction float64
}

func (v *compVelocity) update(e *element) error {
	x, y := e.getPos()
	vx, vy := e.getVel()
	if vx > v.max {
		vx = v.max
	} else if vx < -v.max {
		vx = -v.max
	}

	if vy > v.max {
		vy = v.max
	} else if vy < -v.max {
		vy = -v.max
	}
	e.setVel(vx, vy)
	e.setPos(x+vx, y+vy)

	e.vx *= v.friction
	e.vy *= v.friction

	return nil
}

func (v *compVelocity) getName() string {
	return "velocity"
}

func (e *element) newVelocityApplier(max float64, scale float64, friction float64) {
	v := &compVelocity{max: max, scale: scale, friction: friction}
	e.addComponent(v)
}

// BlockCollider

type compBlockCollider struct {
	componentActive
}

func (v *compBlockCollider) update(e *element) error {
	// Get tiles which we might have to interact with
	x, y := e.getTopLeft()

	xi := (int(x) - int(x)%tileSize) / tileSize
	yi := (int(y) - int(y)%tileSize) / tileSize

	// get element size used to know which blocks to interact with
	tileWidth := int(e.w/tileSizeF) + 1
	tileHeight := int(e.h/tileSizeF) + 1

	tileSet := make([]*tile, 0, 6)

	for tx, ty := 0, 0; tx <= tileWidth || ty < tileHeight; {

		// Ensure valid indexes
		if xi+tx < 0 || xi+tx >= len(e.l.mapLayout) || yi+ty < 0 || yi+ty >= len(e.l.mapLayout[0]) {
			goto postLoop
		}

		{
			// Get tile may be a nil tile
			t := e.l.mapLayout[xi+tx][yi+ty]
			// Ensure the tile is active aka inited
			if t.isActive() {
				tileSet = append(tileSet, t)
			}
		}

	postLoop:
		if tx <= tileWidth {
			tx++
		} else {
			ty++
			tx = 0
		}
	}
	for _, t := range tileSet {
		t.tileCollide(e)
	}

	// interact with a block depending on the direction you are moving

	return nil
}

func (v *compBlockCollider) getName() string {
	return "block collider"
}

func (e *element) newBlockCollider() {
	v := &compBlockCollider{}
	e.addComponent(v)
}

// KeyboardController

type compKeyboardController struct {
	componentActive
	left          func()
	right         func()
	up            func()
	down          func()
	upRight       func()
	upLeft        func()
	downRight     func()
	downLeft      func()
	actionBtnA    func()
	actionBtnB    func()
	actionBtnX    func()
	actionBtnY    func()
	actionBtnMenu func()
}

func (v *compKeyboardController) update(e *element) error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) && ebiten.IsKeyPressed(ebiten.KeyUp) {
		v.upRight()
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyUp) {
		v.upLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyDown) {
		v.downLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyDown) {
		v.downLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		v.left()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		v.right()
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		v.down()
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		v.up()
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		v.actionBtnA()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		v.actionBtnB()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		v.actionBtnX()
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		v.actionBtnY()
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		v.actionBtnMenu()
	}
	return nil
}

func (v *compKeyboardController) getName() string {
	return "keyboard controller"
}

func (e *element) newKeyboardController(
	right func(),
	left func(),
	up func(),
	down func(),
	upRight func(),
	upLeft func(),
	downRight func(),
	downLeft func(),
	actionBtnA func(),
	actionBtnB func(),
	actionBtnX func(),
	actionBtnY func(),
	actionBtnMenu func()) {
	c := &compKeyboardController{
		left:          left,
		right:         right,
		up:            up,
		down:          down,
		upRight:       upRight,
		upLeft:        upLeft,
		downRight:     downRight,
		downLeft:      downLeft,
		actionBtnA:    actionBtnA,
		actionBtnB:    actionBtnB,
		actionBtnX:    actionBtnX,
		actionBtnY:    actionBtnY,
		actionBtnMenu: actionBtnMenu}
	e.addComponent(c)
}
