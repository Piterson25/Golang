package main

import (
	"fmt"
	"lab09/ant"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowTitle      = "ANTS"
	FPS              = 1000
	WindowTitleDelay = time.Second / FPS
)

func main() {
	width, height, X := parseArguments(os.Args)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Error in loading SDL2: %s\n", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in creating window: %s\n", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Nie można utworzyć renderera: %s\n", err)
		return
	}
	defer renderer.Destroy()

	am := ant.NewManager()

	am.AddAnts(renderer, width, height, X)

	fpsTicker := time.NewTicker(WindowTitleDelay)
	defer fpsTicker.Stop()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				keyEvent := event.(*sdl.KeyboardEvent)
				if keyEvent.Keysym.Sym == sdl.K_q && keyEvent.State == sdl.PRESSED {
					running = false
				}
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		am.Update()

		time.Sleep(WindowTitleDelay)
	}

}

func parseArguments(args []string) (int, int, int) {
	width := 800
	height := 600
	X := 1

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if arg == "-w" && i+1 < len(args) {
			if w, err := parseInt(args[i+1]); err == nil {
				width = w
			}
		} else if arg == "-h" && i+1 < len(args) {
			if h, err := parseInt(args[i+1]); err == nil {
				height = h
			}
		} else if arg == "-count" && i+1 < len(args) {
			if x, err := parseInt(args[i+1]); err == nil {
				X = x
			}
		}
	}

	return width, height, X
}

func parseInt(str string) (int, error) {
	var value int
	_, err := fmt.Sscanf(str, "%d", &value)
	return value, err
}
