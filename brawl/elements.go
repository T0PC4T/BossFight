package brawl

import "github.com/hajimehoshi/ebiten"

// Elements

type element interface {
	setID(int) error
	update() error
	draw(*ebiten.Image) error

	// getters and setters
	getPos() (float64, float64)
	setPos(float64, float64) error
	getVel() (float64, float64)
	setVel(float64, float64) error
	getDim() (int, int)
	setDim(int, int) error

	// components
	getComponentSlicePtr() *[]component
}

func addElement(l level, e element) {
	eSlice := l.GetElementSlicePtr()
	e.setID(len(*eSlice) - 1)
	*eSlice = append(*eSlice, e)
}

func removeElement(l level, id int) bool {
	eSlice := l.GetElementSlicePtr()
	if tSlice := *eSlice; id < len(tSlice) {
		tSlice[id] = tSlice[len(tSlice)-1]
		tSlice = tSlice[:len(tSlice)-1]
		tSlice[id].setID(id)
		*eSlice = tSlice
	} else {
		return false
	}
	return true
}
