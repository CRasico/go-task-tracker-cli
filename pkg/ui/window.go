package ui

import (
	"errors"
	"fmt"
	"strings"
)

type Window struct {
	Pixels    [][]rune
	HasBorder bool
}

type WindowArgs struct {
	Height    int
	Width     int
	HasBorder bool
}

func NewWindow(args WindowArgs) Window {
	pixels := [][]rune{}
	for range args.Height {
		row := make([]rune, args.Width)
		for i := range args.Width {
			row[i] = ' '
		}
		pixels = append(pixels, row)
	}

	return Window{
		Pixels:    pixels,
		HasBorder: args.HasBorder,
	}
}

func (window *Window) Draw(x int, y int, character rune) error {
	if x >= len(window.Pixels) || y >= len(window.Pixels[0]) {
		return errors.New("Cannot draw pixel outside of window")
	}

	window.Pixels[x][y] = character

	return nil
}

func (window *Window) Render() {
	screen := ""

	for i := range window.Pixels {
		row := string(window.Pixels[i])

		// add border to row item if we have one
		if window.HasBorder {
			row = "|" + row + "|"
		}
		screen += row + "\n"
	}

	if window.HasBorder {
		border := strings.Repeat("-", len(window.Pixels)+2)
		screen = border + "\n" + screen + border + "\n"
	}

	fmt.Print(screen)
}
