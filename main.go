// Package main runs the shopping list
package main

import (
	"flag"
	"go-snake/pkg/canvas"
	"go-snake/pkg/game"
	"go-snake/pkg/snake"
	"log"

	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
)

const (
	size          = 8
	defaultConfig = "eight-by-eight.json"
)

func main() {
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	var configName = flag.String("config", defaultConfig, "run in debug mode")
	flag.Parse()

	if *configName == defaultConfig {
		log.Printf("Using default config file %s\n", defaultConfig)
	}

	cfg, err := config.LoadConfig(*configName)
	if err != nil {
		log.Fatal(err)
	}

	snk := snake.NewSnake([]int{4, 2}, [][]int{{5, 2}, {6, 2}, {7, 2}}, size, size)

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

	cvs := canvas.NewCanvas(size, size, &snk)

	if err := game.Loop(scrn, cvs, &snk); err != nil && err != game.ErrGameFin {
		log.Panicln(err)
	}
	log.Println("GOOD GAME!")
	log.Printf("Your score was %d", snk.Score())
}
