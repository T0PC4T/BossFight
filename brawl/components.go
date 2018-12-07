package brawl

import (
	"github.com/hajimehoshi/ebiten"
)

// Component definition

type component interface {
	update(*element) error
	isActive() bool
	setActive(a bool)
}

// Components

type componentActive bool

func (c *componentActive) isActive() bool {
	return bool(*c)
}

func (c *componentActive) setActive(a bool) {
	v := componentActive(a)
	c = &v
}

// Pusher Componenet

type compPusher struct {
	componentActive
	vx, vy float64
}

func (g *compPusher) update(e *element) error {
	vx, vy := e.getVel()
	e.setVel(vx+g.vx, vy+g.vy)
	return nil
}

func (g *compPusher) getName() string {
	return "gravity"
}

func (e *element) newPusherComponent(vx, vy float64) {
	cg := &compPusher{vx: vx, vy: vy}
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
	} else if vx > -0.1 && vx < 0.1 {
		vx = 0
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

type compAnimator struct {
	componentActive
	groundFunc func()
	vx         float64
	vy         float64
}

func (v *compAnimator) update(e *element) error {
	if e.vy == 0 && v.vy == 0 {
		v.groundFunc()
	}
	if e.vx == 0 && e.s.status != "idle" {
		e.s.UpdateStatus("idle")
	} else if e.s.status != "running" {
		e.s.UpdateStatus("running")
	}

	if e.vx > 0 && e.s.movingRight == false {
		e.s.movingRight = true
	} else if e.vx < 0 && e.s.movingRight == true {
		e.s.movingRight = false
	}

	v.vy = e.vy
	return nil
}

func (e *element) newAnimator(groundFunc func()) {
	c := &compAnimator{groundFunc: groundFunc}
	e.addComponent(c)
}
