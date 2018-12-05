package loader

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/T0PC4T/BossFight/assets"
	"github.com/hajimehoshi/ebiten"
)

var (
	// player
	RamboSpriteSheet *ebiten.Image
	// tiles
	Grasstile *ebiten.Image
	BrickType *ebiten.Image
	// Background image
	BackgroundImg *ebiten.Image
)

func InitialLoad() error {

	if img, _, err := image.Decode(bytes.NewReader(assets.Background)); err != nil {
		return err
	} else {
		BackgroundImg, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	if img, _, err := image.Decode(bytes.NewReader(assets.BrickTile)); err != nil {
		return err
	} else {
		BrickType, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	if img, _, err := image.Decode(bytes.NewReader(assets.Runner_png)); err != nil {
		return err
	} else {
		RamboSpriteSheet, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}

	if img, _, err := image.Decode(bytes.NewReader(assets.TilePng)); err != nil {
		return err
	} else {
		Grasstile, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	}
	return nil

}
