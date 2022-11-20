// Package main runs the shopping list
package main

import (
	"go-snake/pkg/canvas"
	"go-snake/pkg/snake"
	"time"

	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
)

const (
	size = 10
)

func main() {
	scrn := gui.NewTerminalGui(config.PinConfig{
		RowPins: make([]int, size),
		ColPins: make([]int, size),
	})

	snk := snake.NewSnake([]int{4, 2}, [][]int{{5, 2}, {6, 2}, {7, 2}}, size, size)

	cvs := canvas.NewCanvas(size, size, &snk)
	cvs.GetMatrix()
	scrn.DisplayMatrix(cvs.GetMatrix(), time.Second)

	for j := 0; j < 100; j++ {
		snk.NextMove()
		scrn.DisplayMatrix(cvs.GetMatrix(), time.Second)
	}
	
	// snk.ChangeDirection(snake.East)
	snk.NextMove()
	scrn.DisplayMatrix(cvs.GetMatrix(), time.Second)
	// snk.ChangeDirection(snake.South)
	snk.NextMove()
	scrn.DisplayMatrix(cvs.GetMatrix(), time.Second)
	snk.NextMove()
	scrn.DisplayMatrix(cvs.GetMatrix(), time.Second)
	// snk.ChangeDirection(snake.West)
	snk.NextMove()
	scrn.DisplayMatrix(cvs.GetMatrix(), time.Second)
}
