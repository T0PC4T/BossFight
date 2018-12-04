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

func (g *compGravity) setID(ID int) {
	g.ID = ID
}

func (e *element) newGravityComponent(gravityAmount float64) {
	cg := &compGravity{gravityAmount: gravityAmount}
	e.addComponent(cg)
}

// Velocity Applier

type newVelocityApplier struct {
	ID            int
	gravityAmount float64
}

func (e *element) newVelocityApplier(scale float64) {

}
