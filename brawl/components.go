package brawl

// Component definition

type component interface {
	update(*element) error
	setID(int)
}

// Components

type compGravity struct {
	ID            int
	gravityAmount float64
}

func (g *compGravity) update(e *element) error {
	x, y := e.getPos()
	e.setPos(x, y+g.gravityAmount)
	return nil
}

func (g *compGravity) setID(ID int) {
	g.ID = ID
}

func (e *element) newGravityComponent(gravityAmount float64) {
	cg := &compGravity{gravityAmount: gravityAmount}
	e.addComponent(cg)
}
