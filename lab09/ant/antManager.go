package ant

import "github.com/veandco/go-sdl2/sdl"

type antManager struct {
	ants        []*Ant
	visitedGrid map[point]sdl.Color
}

func NewManager() *antManager {
	return &antManager{
		ants:        []*Ant{},
		visitedGrid: make(map[point]sdl.Color),
	}
}

func (am *antManager) AddAnts(renderer *sdl.Renderer, width, height, X int) {
	for i := 0; i < X; i++ {
		am.ants = append(am.ants, NewAnt(renderer, int32(width), int32(height), 10))
	}
}

func (am *antManager) Update() {
	for _, a := range am.ants {
		a.Update()
		p, color := a.markVisited(a.pointX, a.pointY, a.direction)
		am.visitedGrid[p] = color

		for p, color := range am.visitedGrid {
			a.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			rect := sdl.Rect{
				X: p.x,
				Y: p.y,
				W: a.gridSize,
				H: a.gridSize,
			}
			a.renderer.FillRect(&rect)
		}
		a.Render()
	}
}
