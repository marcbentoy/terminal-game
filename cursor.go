package main

import "github.com/gdamore/tcell/v2"

type Cursor struct {
	X, Y   int
	Char   rune
	Mode   int
	Screen *tcell.Screen
}

func NewCursor(screen *tcell.Screen) *Cursor {
	sWidth, sHeight := (*screen).Size()

	return &Cursor{
		Screen: screen,
		Mode:   0,
		X:      sWidth / 2,
		Y:      sHeight / 2,
		Char:   '.',
	}
}

func (c *Cursor) Draw() {

}
