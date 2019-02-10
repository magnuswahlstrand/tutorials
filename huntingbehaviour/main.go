package main

import (
	"image"
	"image/color"

	"github.com/SolarLune/paths"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

const tileSize = 16

func main() {
	// Layout of the room
	layout := []string{
		"xxxxxxxx",
		"x   x  x",
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

			// Draw the tile
			drawTile(img, x, y, c)
		}
	}

	room := paths.NewGridFromStringArrays(layout)
	start, dest := room.Get(1, 6), room.Get(5, 1)

	// Turn off movement in walls
	for _, cell := range room.GetCellsByRune('x') {
		cell.Walkable = false
	}
	path := room.GetPath(start, dest, false)

	// Draw path and start and finish
	for _, tile := range path.Cells {
		drawTile(img, tile.X, tile.Y, colornames.Pink)
	}

	drawTile(img, start.X, start.Y, colornames.Blue)
	drawTile(img, dest.X, dest.Y, colornames.Red)
	gfx.SavePNG("images/basic_3.png", img)
}

func drawTile(img *image.NRGBA, x, y int, c color.Color) {
	tileRectangle := image.Rect(0, 0, tileSize, tileSize).Add(image.Pt(x*tileSize, y*tileSize))
	gfx.DrawImageRectangle(img, tileRectangle, c)
}
