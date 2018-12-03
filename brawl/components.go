package brawl

// Component definition

type component interface {
	update(element) error
	setID(int) error
}

// Components

type compGravity struct {
	ID            int
	gravityAmount float64
}

func (g *compGravity) update(e element) {
	x, y := e.getPos()
	e.setPos(x, y+g.gravityAmount)
}

func (g *compGravity) setID(ID int) {
	g.ID = ID
}

func newGravityComponent(gravityAmount float64) *compGravity {
	return &compGravity{}
}
