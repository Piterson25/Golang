package ant

import "github.com/veandco/go-sdl2/sdl"

type AntsManager struct {
	ants        []*Ant
	visitedGrid map[point]sdl.Color
}

func NewManager() *AntsManager {
	return &AntsManager{
		ants:        []*Ant{},
		visitedGrid: make(map[point]sdl.Color),
	}
}

func (am *AntsManager) AddAnts(renderer *sdl.Renderer, width, height, X int) {
	for i := 0; i < X; i++ {
		am.ants = append(am.ants, NewAnt(renderer, int32(width), int32(height), 30))
	}
}

func (am *AntsManager) Update() {
	for _, a := range am.ants {
		a.Update()
		p, color := a.markVisited(a.pointX, a.pointY, a.direction)
		am.visitedGrid[p] = color
	}
}

func (am *AntsManager) Render() {
	for _, a := range am.ants {
		for p, color := range am.visitedGrid {
			a.renderer.SetDrawColor(color.R, color.G, color.B, color.A)
			a.renderer.FillRect(&sdl.Rect{X: p.x, Y: p.y, W: a.gridSize, H: a.gridSize})
		}
	}

	for _, a := range am.ants {
		a.Render()
	}
}
