package main

import (
	"image"
	"image/color"

	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

const tileSize = 16

func main() {
	// Layout of the room
	layout := []string{
		"xxxxxxxx",
		"x      x",
		"x   xx x",
		"x      x",
		"x      x",
		"x x    x",
		"x x    x",
		"xxxxxxxx",
	}
	height := len(layout)
	width := len(layout[0])
	img := gfx.NewImage(width*tileSize, height*tileSize, color.Transparent)

	// Draw the room
	for y, row := range layout {
		for x, tile := range row {
			c := colornames.Black
			switch tile {
			case 'x':
				c = colornames.Grey
			}
			tileRectangle := image.Rect(0, 0, tileSize, tileSize).Add(image.Pt(x*tileSize, y*tileSize))
			gfx.DrawImageRectangle(img, tileRectangle, c)
		}
	}
	gfx.SavePNG("images/out.png", img)
}
