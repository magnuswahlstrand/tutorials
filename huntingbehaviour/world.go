package main

import (
	"image"
	"image/color"

	"github.com/SolarLune/paths"
	"github.com/peterhellberg/gfx"
	"golang.org/x/image/colornames"
)

type World struct {
	background    *image.NRGBA
	room          *paths.Grid
	width, height int
}

func newWorld(layout []string) World {
	height := len(layout)
	width := len(layout[0])
	img := gfx.NewImage(width*tileSize, height*tileSize, color.Transparent)

	// Draw the room
	for y, row := range layout {
		for x, tile := range row {
			var c color.RGBA
			switch tile {
			case 'x':
				c = colornames.Grey
			case '|':
				c = colornames.Brown
			case '-':
				c = colornames.Green
			default:
				c = colornames.Black
			}

			drawTile(img, x, y, c)
		}
	}

	room := paths.NewGridFromStringArrays(layout)

	// Turn off movement in walls
	for _, cell := range room.GetCellsByRune('x') {
		cell.Walkable = false
	}

	// Turn off movement in air
	for _, cell := range room.GetCellsByRune(' ') {
		cell.Walkable = false
	}

	return World{
		background: img,
		room:       room,
		width:      height,
		height:     width,
	}
}

func (w *World) findPath(start, dest *paths.Cell) []gfx.Vec {
	path := []gfx.Vec{}
	for _, cell := range w.room.GetPath(start, dest, false).Cells {
		path = append(path, gfx.V(float64(cell.X*tileSize), float64(cell.Y*tileSize)))
	}
	return path
}
