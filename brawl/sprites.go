package brawl

import (
	"image"

	"github.com/T0PC4T/BossFight/loader"
	"github.com/hajimehoshi/ebiten"
)

type animation struct {
	spriteSheet   *ebiten.Image
	frameOX       int
	frameOY       int
	frameWidth    int
	frameHeight   int
	frameNum      int
	ticksPerFrame int
	currentFrame  int
	currentTick   int
}

type sprite struct {
	movingRight bool
	animations  map[string]*animation
	status      string
}

func (s *sprite) UpdateStatus(status string) error {
	s.animations[s.status].reset()
	s.status = status
	return nil
}

func (s *sprite) draw(canvas *ebiten.Image) error {
	w, h := canvas.Size()
	subCanvas, _ := ebiten.NewImage(w, h, ebiten.FilterDefault)
	s.animations[s.status].draw(subCanvas)

	// if moving Left flip this (and in future for enacting other effects)
	op := &ebiten.DrawImageOptions{}
	canvas.DrawImage(subCanvas, op)

	return nil
}

func (a *animation) draw(canvas *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	sx, sy := a.frameOX+a.currentFrame*a.frameWidth, a.frameOY
	canvas.DrawImage(loader.RamboSpriteSheet.SubImage(image.Rect(sx, sy, sx+a.frameWidth, sy+a.frameHeight)).(*ebiten.Image), op)

	// Get next frame

	if a.currentTick < a.ticksPerFrame {
		a.currentTick++
	} else {
		if a.currentFrame < a.frameNum {
			a.currentFrame++
		} else {
			a.currentFrame = 0
		}
		a.currentTick = 0
	}

	return nil
}

func (a *animation) reset() error {
	a.currentFrame = 0
	a.currentTick = 0
	return nil
}
