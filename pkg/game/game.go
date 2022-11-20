// Package game is responsible for Snake's game logic
package game

import (
	"fmt"
	"go-snake/pkg/canvas"
	"go-snake/pkg/snake"
	"log"
	"time"

	"github.com/gowhale/led-matrix-golang/pkg/gui"
	term "github.com/nsf/termbox-go"
)

var (
	ErrGameFin = fmt.Errorf("fin") // ErrGameFin is the error which returns when the game is finished
)

// passTime makes things happen at certain timeframes
func passTime(snk *snake.Snake, errs chan error) {
	startTime := time.Now()
	lastGrow := startTime
	lastScreenRefresh := startTime
	for {
		timeSinceGrow := time.Since(lastGrow)
		if timeSinceGrow > time.Second*5 {
			snk.SetGrow(true)
			lastGrow = time.Now()
		}
		timeSinceRefresh := time.Since(lastScreenRefresh)
		if timeSinceRefresh > time.Millisecond*100 {
			lastScreenRefresh = time.Now()
			snk.NextMove()
			if snk.Dead() {
				errs <- ErrGameFin
			}
		}
	}
}

// checkKeyInput gets the users key inputs
func checkKeyInput(snk *snake.Snake, errs chan error) {
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				errs <- fmt.Errorf("esc pressed")
			case term.KeyArrowUp:
				log.Println("Going North")
				snk.ChangeDirection(snake.North)
			case term.KeyArrowLeft:
				log.Println("Going west")
				snk.ChangeDirection(snake.West)
			case term.KeyArrowRight:
				log.Println("Going East")
				snk.ChangeDirection(snake.East)
			case term.KeyArrowDown:
				log.Println("Going South")
				snk.ChangeDirection(snake.South)
			default:
				log.Println("Invalid key")
			}
		case term.EventError:
			errs <- ev.Err
		}
	}
}

// displayGame displays the gamescreen
func displayGame(scrn gui.Screen, cvs canvas.Canvas, errs chan error) {
	for {
		log.Println("Displaying!")
		if err := scrn.DisplayMatrix(cvs.GetMatrix(), time.Millisecond*50); err != nil {
			errs <- err
		}
		if err := scrn.AllLEDSOff(); err != nil {
			errs <- err
		}
	}
}

// Loop runs the games main loop
func Loop(scrn gui.Screen, cvs canvas.Canvas, snk *snake.Snake) error {
	if err := term.Init(); err != nil {
		return err
	}
	defer term.Close()

	// Error handler
	errs := make(chan error, 1)

	// Run three async funcs at same time
	go passTime(snk, errs)
	go checkKeyInput(snk, errs)
	go displayGame(scrn, cvs, errs)

	// Handle any errs that come from the async loops
	for {
		err := <-errs
		if err != nil {
			return err
		}
	}
}
