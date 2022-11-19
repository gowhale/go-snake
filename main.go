// Package main runs the shopping list
package main

import (
	"go-snake/pkg/canvas"
	"go-snake/pkg/snake"
	"time"

	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
)

func main() {
	scrn := gui.NewTerminalGui(config.PinConfig{
		RowPins: make([]int, 10),
		ColPins: make([]int, 10),
	})

	snk := snake.NewSnake([]int{1, 2}, [][]int{{2, 2}, {3, 2}, {4, 2}})

	cvs := canvas.NewCanvas(10, 10, snk)
	cvs.GetMatrix()
	scrn.DisplayMatrix(cvs.GetMatrix(), time.Second)
}
