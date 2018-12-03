package main

import (
	"fmt"

	"github.com/T0PC4T/BossFight/brawl"
	"github.com/hajimehoshi/ebiten"
)

type level interface {
	Update() (bool, error)
	Draw(*ebiten.Image) error
}

type gameBrawl struct {
	currentLevel level
	nextLevel    level
}

func newGameBrawl() *gameBrawl {
	// Choose and load level
	// Load level
	dl := brawl.NewDevelopmentLevel()
	// add players and settings
	brawl.NewRambo(dl)

	g := &gameBrawl{currentLevel: dl}
	return g
}

func (gm gameBrawl) getType() string {
	return "game"
}

func (gm gameBrawl) loop(screen *ebiten.Image) error {

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
