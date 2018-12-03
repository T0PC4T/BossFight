package brawl

import "github.com/hajimehoshi/ebiten"

type hero interface {
	Update() error
	Draw(*ebiten.Image) error
}

func heroesUpdate() {

}
