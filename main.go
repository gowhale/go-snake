// Package main runs the shopping list
package main

import (
	"flag"
	"log"

	"github.com/gowhale/go-snake/pkg/canvas"
	"github.com/gowhale/go-snake/pkg/game"
	"github.com/gowhale/go-snake/pkg/snake"

	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
)

const (
	defaultConfig = "eight-by-eight.json"
)

func main() {
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	var configName = flag.String("config", defaultConfig, "run in debug mode")
	flag.Parse()

	if *configName == defaultConfig {
		log.Printf("Using default config file %s\n", defaultConfig)
	}

	// // To make a custom size game uncomment the following code:
	// cfg := config.PinConfig{
	// 	RowPins: make([]int, 20),
	// 	ColPins: make([]int, 20),
	// }
	cfg, err := config.LoadConfig(*configName)
	if err != nil {
		log.Fatal(err)
	}

	snk := snake.NewSnake([]int{4, 2}, [][]int{{5, 2}, {6, 2}, {7, 2}}, cfg.ColCount(), cfg.RowCount())

	scrn := gui.NewTerminalGui(cfg)
	if !*debugMode {
		var err error
		scrn, err = gui.NewledGUI(cfg)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		err := scrn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	cvs := canvas.NewCanvas(cfg.ColCount(), cfg.RowCount(), &snk)

	if err := game.Loop(scrn, cvs, &snk); err != nil && err != game.ErrGameFin {
		log.Panicln(err)
	}
	log.Println("GOOD GAME!")
	log.Printf("Your score was %d", snk.Score())
}
