package brawl

// Component definition

type component interface {
	update(element) error
	setID(int) error
}

func addComponent(e element, c component) {
	cSlice := e.getComponentSlicePtr()
	c.setID(len(*cSlice) - 1)
	*cSlice = append(*cSlice, c)
}

func removeComponent(e element, id int) bool {
	cSlice := e.getComponentSlicePtr()
	if tSlice := *cSlice; id < len(tSlice) {
		tSlice[id] = tSlice[len(tSlice)-1]
		tSlice = tSlice[:len(tSlice)-1]
		tSlice[id].setID(id)
		*cSlice = tSlice
	} else {
		return false
	}
	return true
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
