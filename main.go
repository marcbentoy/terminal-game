package main

import (
	// "fmt"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	tealColor = tcell.ColorTeal
	blueColor = tcell.ColorBlue
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("error creating new screen:", err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatal("error initializing screen:", err)
	}

	player := NewSprite('.', 10, 10)
	player2 := NewSprite('.', 10, 11)

	sprites := []*Sprite{}

	ticker := time.NewTicker(time.Second)
	doneTicking := make(chan bool)
	prevColor := 0

	width, height := screen.Size()
	cursor := Cursor{
		X:    width / 2,
		Y:    height / 2,
		Char: '.',
	}

	go func() {
		for {
			select {
			case <-doneTicking:
				return
			case _ = <-ticker.C:
				if prevColor == 0 {
					player.Color = tealColor
					player2.Color = blueColor
					prevColor = 1
				} else {
					player.Color = blueColor
					player2.Color = tealColor
					prevColor = 0
				}

				screen.Clear()

				player.Draw(screen)
				player2.Draw(screen)
				for _, s := range sprites {
					screen.SetContent(s.X, s.Y, s.Char, nil, tcell.StyleDefault.Foreground(s.Color))
				}

				screen.SetContent(cursor.X, cursor.Y, cursor.Char, nil, tcell.StyleDefault.Background(tcell.ColorTeal))
				screen.Show()
			}
		}
	}()

	// selectedChar := '.'

	running := true
	for running {
		screen.Clear()

		screen.Show()

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Rune() {
			case ' ':
				sprites = append(sprites, &Sprite{
					Char:  cursor.Char,
					X:     cursor.X,
					Y:     cursor.Y,
					Color: tcell.ColorWhite,
				})
				continue
			case 'j', 's':
				cursor.Y += 1
			case 'k', 'w':
				cursor.Y -= 1
			case 'h', 'a':
				cursor.X -= 1
			case 'l', 'd':
				cursor.X += 1
			case 'q':
				ticker.Stop()
				doneTicking <- true
				running = false
			}
		}
	}
}
