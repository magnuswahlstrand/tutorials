# Hunting behaviour

### Background

The other day I found [this presentation](https://www.youtube.com/watch?v=sVntwsrjNe4&t=346s8) about some of the animation techniques used in the survival platformer game [Rain World](https://en.wikipedia.org/wiki/Rain_World). I really liked a few of the concepts:

- [Procedurally generated animation](https://ksr-ugc.imgix.net/assets/001/529/743/7984c7dfe0ae05dcbc820a6c20d14925_original.gif?ixlib=rb-1.1.0&w=639&fit=max&v=1389817681&auto=format&gif-q=50&q=92&s=b885beac30be769418377538ba9683d1) (as opposed to sprites)
- [Credible creature AI](https://ksr-ugc.imgix.net/assets/017/313/774/0519921b0e69724a97f59647e917bf00_original.gif?ixlib=rb-1.1.0&w=639&fit=max&v=1498766217&auto=format&gif-q=50&q=92&s=fabe3347c62481e86b387e0d0569577c)
- ...

I haven't actually played the game, but after binge watching a few hours of game play, I decided to try to implement some of the concepts myself. The first thing I will try is the hunting behaviour of the lizards.

![hunting](https://ksr-ugc.imgix.net/assets/002/919/650/83c9175342c8aacb552aaf90bcac7f4d_original.gif?ixlib=rb-1.1.0&w=639&fit=max&v=1416160417&auto=format&gif-q=50&q=92&s=6e8d173b25bfe54f0b7ebfc1f90989a1)

To keep this article as simple as possible, I will try to use existing libraries whenever possible.

## Our goal

1. Drawing the world
2. Basic path finding
3. Smooth path finding
4. (optional) Improve the graphics

## 1. Drawing the world

First, let's draw our world. We define it using string slices (will be used by the path finding algorithm later).

- `'x'` is a wall
- `' '` is empty space

```go
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
```

Here I use Peter Hellberg's `gfx` package for creating a `NewImage`, `DrawImageRectangle` and finally `SavePNG`.

```go
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
            tileRectangle := image.Rect(0, 0, tileSize, tileSize).Add(image.Pt(x*tileSize, y*tileSize))
            gfx.DrawImageRectangle(img, tileRectangle, c)
        }
    }
    gfx.SavePNG("images/out.png", img)
}
```

Here is the result. Gray is wall (`x`), black is empty space (``).

![room](images/out.png)

Done!

## 2. Basic path finding

Now that we can visualize our world, let's try finding a way from point A to point B in it!

For this, we will use an A\* (a star) algorithm which is popular in game development. Read [this article](http://theory.stanford.edu/~amitp/GameProgramming/AStarComparison.html) for an overview of the algorithm, and [this](https://www.redblobgames.com/pathfinding/a-star/introduction.html) for a more in-depth version with some code. There are [quite a few](https://godoc.org/?q=astar) go implementations of the algorithm. I chose [SolarLune/paths](https://github.com/SolarLune/paths), because I have used some of his packages before.

### SolarLune/paths

Using the package is very straight forward. First we define our grid from the string slice from before, then we can get the path from `start` to `dest` (_we use `false` to turn of diagonal movement_).

```go
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
    // Abbrieviated
    ...
    room := paths.NewGridFromStringArrays(layout)
    start, dest := room.Get(1, 6), room.Get(5, 1)
    path := room.GetPath(start, dest, false)

    // Draw path, start and finish
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

```

**Result:**

|                                | + path                         |
| ------------------------------ | ------------------------------ |
| ![basic_1](images/basic_1.png) | ![basic_2](images/basic_2.png) |

- start (blue)
- dest (blue)
- path (pink)

As you can see, it _almost_ works, beside from the fact that we walk through a wall left of the goal. Let's fix that by saying that the cells marked by `x` should not be walkable.

```go
    ...
    // Turn off movement in walls
    for _, cell := range room.GetCellsByRune('x') {
        cell.Walkable = false
    }

    path := room.GetPath(start, dest, false)
    ...
```

Much better! **Done**

|                                | + path                         |
| ------------------------------ | ------------------------------ |
| ![basic_1](images/basic_1.png) | ![basic_3](images/basic_3.png) |
