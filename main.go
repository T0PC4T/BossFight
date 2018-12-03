package main

import (
	"log"

	. "github.com/T0PC4T/BossFight/global"
	"github.com/T0PC4T/BossFight/loader"
	"github.com/hajimehoshi/ebiten"
)

type displayView interface {
	getType() string
	loop(*ebiten.Image) error
}

var (
	currentScreen displayView
	allScreens    map[string]displayView
)

func init() {
	allScreens = make(map[string]displayView)
	brawl := newGameBrawl()
	allScreens["development"] = brawl
	currentScreen = allScreens["development"]
}

func update(s *ebiten.Image) error {
	return currentScreen.loop(s)
}

func main() {
	if err := loader.InitialLoad(); err != nil {
		log.Fatal("Failed loading initial assets with: ", err)
	}

	ebiten.SetFullscreen(true)
	ebiten.SetRunnableInBackground(true)

	if err := ebiten.Run(update, ScreenWidth, ScreenHeight, 1, "Boss Fight"); err != nil {
		panic(err)
	}
}
