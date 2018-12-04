package brawl

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type GameBrawl struct {
	currentLevel *level
	nextLevel    *level
}

func NewGameBrawl() *GameBrawl {
	// Choose and load level
	// Load level
	l := newDevelopmentLevel()
	// add players and settings
	NewRambo(l)

	gb := &GameBrawl{currentLevel: l}
	return gb
}

// GetType get the type of game mode/status
func (gm GameBrawl) GetType() string {
	return "game"
}

// Display called as main game loop (normally update but Display is for a whole game state)
func (gm GameBrawl) Display(screen *ebiten.Image) error {

	if finished, err := gm.currentLevel.update(); err != nil {
		if finished {
			fmt.Println("Finished")
		}
		return err
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err := gm.currentLevel.draw(screen)

	return err
}
