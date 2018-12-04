package brawl

// Component definition

type component interface {
	update(*element) error
	getName() string
}

// Components

// Gravity Componenet

type compGravity struct {
	ID            int
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
	ID         int
	max, scale float64
}

func (v *compVelocity) update(e *element) error {
	x, y := e.getPos()
	vx, vy := e.getVel()
	if vx > v.max {
		vx = v.max
	}
	if vy > v.max {
		vy = v.max
	}
	e.setVel(vx, vy)
	e.setPos(x+vx, y+vy)
	return nil
}

func (v *compVelocity) getName() string {
	return "velocity"
}

func (e *element) newVelocityApplier(max float64, scale float64) {
	v := &compVelocity{max: max, scale: scale}
	e.addComponent(v)
}

// Velocity Applier

type compBlockCollider struct {
	ID int
}

func (v *compBlockCollider) update(e *element) error {

	tileWidth := (e.w / tileSize)
	tileHeight := (e.h / tileSize)
	x, y := e.getPos()
	xi := (int(x) - int(x)%tileSize) / tileSize
	yi := (int(y) - int(y)%tileSize) / tileSize

	if e.vx > 0 {
		if t := e.l.mapLayout[xi+tileWidth][yi]; t.isActive() {
			e.vx = 0
		}
	} else if e.vx < 0 {
		if t := e.l.mapLayout[xi][yi]; t.isActive() {
			e.vx = 0
		}
	}

	if e.vy > 0 {
		if t := e.l.mapLayout[xi][yi+tileHeight]; t.isActive() {
			e.vy = 0
		}
	} else if e.vy < 0 {
		if t := e.l.mapLayout[xi][yi]; t.isActive() {
			e.vy = 0
		}
	}

	return nil
}

func (v *compBlockCollider) getName() string {
	return "block collider"
}

func (e *element) newBlockCollider() {
	v := &compBlockCollider{}
	e.addComponent(v)
}
