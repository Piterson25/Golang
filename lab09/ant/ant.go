package ant

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

type Ant struct {
	renderer     *sdl.Renderer
	pointX       int32
	pointY       int32
	windowWidth  int32
	windowHeight int32
	gridSize     int32
	visitedGrid  map[point]sdl.Color
	direction    int
}

type point struct {
	x int32
	y int32
}

func NewAnt(renderer *sdl.Renderer, windowWidth, windowHeight, gridSize int32) *Ant {
	return &Ant{
		renderer:     renderer,
		pointX:       windowWidth / 2,
		pointY:       windowHeight / 2,
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		gridSize:     gridSize,
		visitedGrid:  make(map[point]sdl.Color),
		direction:    0,
	}
}

func (a *Ant) Update() {
	dx := int32(0)
	dy := int32(0)

	direction := rand.Intn(4)
	switch direction {
	case 0:
		dx = a.gridSize
	case 1:
		dx = -a.gridSize
	case 2:
		dy = a.gridSize
	case 3:
		dy = -a.gridSize
	}

	a.direction = direction

	a.pointX += dx
	a.pointY += dy

	a.markVisited(a.pointX, a.pointY, a.direction)
}

func (a *Ant) Render() {
	for p, color := range a.visitedGrid {
		a.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
		rect := sdl.Rect{
			X: p.x,
			Y: p.y,
			W: a.gridSize,
			H: a.gridSize,
		}
		a.renderer.FillRect(&rect)
	}

	a.renderer.SetDrawColor(255, 255, 255, 255)
	rect := sdl.Rect{
		X: a.pointX,
		Y: a.pointY,
		W: a.gridSize,
		H: a.gridSize,
	}
	a.renderer.FillRect(&rect)

	a.renderer.Present()
}

func (a *Ant) markVisited(x, y int32, direction int) {
	color := sdl.Color{}
	switch direction {
	case 0:
		color = sdl.Color{255, 0, 0, 255} // Czerwony
	case 1:
		color = sdl.Color{0, 255, 0, 255} // Zielony
	case 2:
		color = sdl.Color{255, 255, 0, 255} // Żółty
	case 3:
		color = sdl.Color{0, 0, 255, 255} // Niebieski
	}

	p := point{x, y}
	a.visitedGrid[p] = color
}

func (a *Ant) HitEdge() bool {
	if a.pointX <= 0 || a.pointX >= a.windowWidth-a.gridSize || a.pointY <= 0 || a.pointY >= a.windowHeight-a.gridSize {
		return true
	}
	return false
}
