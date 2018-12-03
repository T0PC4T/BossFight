package components

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type controllerGamepad struct {
	gamepadId int
}

type controllerKeyboardMouse struct {
	gamepadId int
}

func (c *controllerKeyboardMouse) getType() string {
	return "KeyboardMouse"
}

func (c *controllerKeyboardMouse) setId(id int) bool {
	if id > 0 {
		panic("cant change controller id from keyboardMouse to controller")
	}
	return true
}

func (c *controllerKeyboardMouse) getMovementRotation() (float64, bool) {
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyW) {
		return 1.75 * math.Pi, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyW) {
		return 1.25 * math.Pi, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && ebiten.IsKeyPressed(ebiten.KeyS) {
		return 0.25 * math.Pi, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && ebiten.IsKeyPressed(ebiten.KeyS) {
		return 0.75 * math.Pi, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		return math.Pi, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		return 1.5 * math.Pi, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		return 0, true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		return 0.5 * math.Pi, true
	}
	return 0, false

}

type controller interface {
	getType() string
	setId(id int) bool
	getMovementRotation() (float64, bool)
}
