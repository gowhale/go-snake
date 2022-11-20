// Package main runs the shopping list
package main

import (
	"go-snake/pkg/canvas"
	"go-snake/pkg/snake"
	"log"
	"time"

	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
	term "github.com/nsf/termbox-go"
)

const (
	size = 8
)

func main() {
	snk := snake.NewSnake([]int{4, 2}, [][]int{{5, 2}, {6, 2}, {7, 2}}, size, size)

	scrn := gui.NewTerminalGui(config.PinConfig{
		RowPins: make([]int, size),
		ColPins: make([]int, size),
	})

	cvs := canvas.NewCanvas(size, size, &snk)
	err := gameLoop(scrn, cvs, &snk)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("GOOD GAME!")
	log.Printf("Your score was %d", snk.Score())

	log.Println("fin!")
}

func gameLoop(scrn gui.Screen, cvs canvas.Canvas, snk *snake.Snake) error {
	term.Init()
	defer term.Close()
	for {
		snk.NextMove()
		if snk.Dead() {
			return nil
		}
		scrn.DisplayMatrix(cvs.GetMatrix(), time.Millisecond)
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				log.Panicln("esc pressed")
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

			}
		case term.EventError:
			return ev.Err
		}

	}
}
