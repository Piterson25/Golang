package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

func DrawCircle(renderer *sdl.Renderer, x, y, r int) {
	cx, cy := x+r, y+r
	for i := 0; i < 360; i++ {
		rad := float64(i) * math.Pi / 180.0
		x := int32(int(float64(r)*math.Cos(rad)) + cx)
		y := int32(int(float64(r)*math.Sin(rad)) + cy)
		renderer.DrawPoint(x, y)
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Error initializing SDL:", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Shapes", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Error creating window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error creating renderer:", err)
		return
	}
	defer renderer.Destroy()

	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawLine(0, 0, 800, 600)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.FillRect(&sdl.Rect{X: 200, Y: 200, W: 400, H: 200})

	renderer.SetDrawColor(0, 255, 0, 255)
	DrawCircle(renderer, 400, 300, 100)

	renderer.Present()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}
}
