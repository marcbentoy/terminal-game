package main

import (
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

	player := NewSprite('v', 10, 10)

	ticker := time.NewTicker(time.Second)
	doneTicking := make(chan bool)
	prevColor := 0

	go func() {
		for {
			select {
			case <-doneTicking:
				return
			case _ = <-ticker.C:
				if prevColor == 0 {
					player.Color = tealColor
					prevColor = 1
				} else {
					player.Color = blueColor
					prevColor = 0
				}

				screen.Clear()
				player.Draw(screen)
				screen.Show()
			}
		}
	}()

	running := true
	for running {
		screen.Clear()

		player.Draw(screen)

		screen.Show()

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Rune() {
			case 'j':
				player.Y += 1
			case 'k':
				player.Y -= 1
			case 'h':
				player.X -= 1
			case 'l':
				player.X += 1
			case 'q':
				ticker.Stop()
				doneTicking <- true
				running = false
			}
		}
	}
}
