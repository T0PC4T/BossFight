package main

import (
	"fmt"
	"log"

	"github.com/T0PC4T/BossFight/brawl"
	"github.com/T0PC4T/BossFight/loader"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type displayView interface {
	GetType() string
	Display(*ebiten.Image) error
}

const (
	screenWidth  int     = 1920
	screenHeight int     = 1080
	tilesWidth   int     = 48
	tilesHeight  int     = 27
	tileSize     int     = screenWidth / tilesWidth
	tileSizeF    float64 = float64(tileSize)
)

var (
	currentScreen displayView
	allScreens    map[string]displayView
)

func init() {
	brawl.InitializeBrawl(screenWidth, screenHeight, tilesWidth, tilesHeight)
	allScreens = make(map[string]displayView)
	brawlMode := brawl.NewGameBrawl()
	allScreens["development"] = brawlMode
	currentScreen = allScreens["development"]
}

func update(s *ebiten.Image) error {
	msg := fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS())
	ebitenutil.DebugPrint(s, msg)
	return currentScreen.Display(s)
}

func main() {
	if err := loader.InitialLoad(); err != nil {
		log.Fatal("Failed loading initial assets with: ", err)
	}

	ebiten.SetFullscreen(true)
	ebiten.SetRunnableInBackground(true)

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Boss Fight"); err != nil {
		panic(err)
	}
}
