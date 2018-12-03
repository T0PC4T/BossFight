package brawl

import "github.com/hajimehoshi/ebiten"

type level interface {
	Update() (bool, error)
	Draw(*ebiten.Image) error
	GetElementSlicePtr() *[]element
}
