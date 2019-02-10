package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

const tileSize = 16

func main() {
	// Layout of the room
	layout := []string{
		"xxxxxxxx",
		"x   x-|x",
		"x   xx|x",
		"x     |x",
		"x|-|  |x",
		"x|x|  |x",
		"x|x|--|x",
		"xxxxxxxx",
	}
	w := newWorld(layout)

	start, dest := w.room.Get(1, 6), w.room.Get(5, 1)

	c := Creature{
		pos:  gfx.V(float64(start.X)*tileSize, float64(start.Y)*tileSize),
		path: w.findPath(start, dest),
	}

	animation := gfx.Animation{
		Delay: 10, // Delay between frames
	}

	// Draw frame
	for _, stepPosition := range c.path {
		img := gfx.NewPaletted(w.width*tileSize, w.width*tileSize, gfx.PaletteEDG32)

		// Background
		draw.Draw(img, img.Bounds(), w.background, image.ZP, draw.Over)

		// Start and finish
		drawTile(img, start.X, start.Y, colornames.Blue)
		drawTile(img, dest.X, dest.Y, colornames.Red)

		// Creature
		creatureRect := gfx.R(-2, -2, 2, 2).Moved(stepPosition)
		creatureRect = creatureRect.Moved(gfx.V(tileSize/2, tileSize/2)) // Move to center of tile
		gfx.DrawImageRectangle(img, creatureRect.Bounds(), colornames.Pink)

		animation.AddPalettedImage(img)
	}
	animation.SaveGIF("images/platformer_2.gif")
}

func drawTile(img draw.Image, x, y int, c color.Color) {
	tileRectangle := image.Rect(0, 0, tileSize, tileSize).Add(image.Pt(x*tileSize, y*tileSize))
	gfx.DrawImageRectangle(img, tileRectangle, c)
}
