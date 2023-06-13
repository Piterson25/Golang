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
	direction    int
	nextColor    sdl.Color
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
		direction:    0,
		nextColor:    sdl.Color{0, 0, 0, 255},
	}
}

func (a *Ant) Update(ants []*Ant, visitedGrid map[point]sdl.Color) {
	dx := int32(0)
	dy := int32(0)

	checkGoing := rand.Intn(5)

	direction := 0

	if checkGoing == 0 {
		direction = rand.Intn(4)
	} else {
		direction = a.direction
	}

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

	nextPosX := a.pointX + dx
	nextPosY := a.pointY + dy

	// Spotkanie z inna mrowka
	for _, a := range ants {
		if a.pointX == nextPosX && a.pointY == nextPosY {
			a.nextColor = sdl.Color{255, 255, 255, 255}
			nextPosX = a.pointX - dx
			nextPosY = a.pointY - dy
			break
		}
	}

	// Kolorki sladow
	for p, color := range visitedGrid {
		if a.pointX == p.x && a.pointY == p.y {
			a.nextColor = color
			break
		}
	}

	a.pointX = nextPosX
	a.pointY = nextPosY

	a.checkWalls()
}

func (a *Ant) Render() {

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

func (a *Ant) markVisited(x, y int32, direction int) (point, sdl.Color) {
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

	return point{x, y}, color

}

func (a *Ant) checkWalls() {

	if a.pointX < 0 {
		a.pointX = a.windowWidth - a.gridSize
	} else if a.pointX > a.windowWidth-a.gridSize {
		a.pointX = 0
	}

	if a.pointY < 0 {
		a.pointY = a.windowHeight - a.gridSize
	} else if a.pointY > a.windowHeight-a.gridSize {
		a.pointY = 0
	}
}
