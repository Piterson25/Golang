package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	width, height := parseArguments(os.Args)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Nie udalo sie zaladowac SDL2: %s\n", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("MRUFKI", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Blad w tworzeniu okna: %s\n", err)
		return
	}
	defer window.Destroy()

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
	}

}

func parseArguments(args []string) (int, int) {
	width := 800
	height := 600

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
		}
	}

	return width, height
}

func parseInt(str string) (int, error) {
	var value int
	_, err := fmt.Sscanf(str, "%d", &value)
	return value, err
}
