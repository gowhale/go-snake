// Package main runs the shopping list
package main

import (
	"flag"
	"go-snake/pkg/canvas"
	"go-snake/pkg/snake"
	"log"
	"time"

	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
	term "github.com/nsf/termbox-go"
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
	// scrn := gui.NewTerminalGui(cfg)

	cvs := canvas.NewCanvas(size, size, &snk)
	term.Init()
	if err := gameLoop(scrn, cvs, &snk); err != nil {
		log.Panicln(err)
	}
	term.Close()
	log.Println("GOOD GAME!")
	log.Printf("Your score was %d", snk.Score())
}

func gameLoop(scrn gui.Screen, cvs canvas.Canvas, snk *snake.Snake) error {
	running := true
	startTime := time.Now()

	go func() error {
		lastGrow := startTime
		lastScreenRefresh := startTime
		for running {
			timeSinceGrow := time.Since(lastGrow)
			if timeSinceGrow > time.Second*5 {
				lastGrow = time.Now()
				snk.Grow()
			}
			timeSinceRefresh := time.Since(lastScreenRefresh)
			if timeSinceRefresh > time.Millisecond*100 {
				lastScreenRefresh = time.Now()
				snk.NextMove()
				if snk.Dead() {
					running = false
					return nil
				}
			}
		}
		return nil
	}()

	go func() error {
		for {
			switch ev := term.PollEvent(); ev.Type {
			case term.EventKey:
				switch ev.Key {
				case term.KeyEsc:
					running = false
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
	}()
	for running {
		log.Println("Displaying!")
		scrn.DisplayMatrix(cvs.GetMatrix(), time.Millisecond*50)
		scrn.AllLEDSOff()
	}
	return nil
}
