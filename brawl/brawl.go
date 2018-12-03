package brawl

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type GameBrawl struct {
	currentLevel level
	nextLevel    level
}

func NewGameBrawl() *GameBrawl {
	// Choose and load level
	// Load level
	dl := newDevelopmentLevel()
	// add players and settings
	NewRambo(dl)

	g := &GameBrawl{currentLevel: dl}
	return g
}

// GetType get the type of game mode/status
func (gm GameBrawl) GetType() string {
	return "game"
}

// Display called as main game loop (normally update but Display is for a whole game state)
func (gm GameBrawl) Display(screen *ebiten.Image) error {

	if finished, err := gm.currentLevel.Update(); err != nil {
		if finished {
			fmt.Println("Finished")
		}
		return err
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err := gm.currentLevel.Draw(screen)

	return err
}
