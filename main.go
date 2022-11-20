// Package main runs the shopping list
package main

import (
	"flag"
	"fmt"
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

var (
	GameFinErr = fmt.Errorf("fin")
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

	if err := gameLoop(scrn, cvs, &snk); err != nil && err != GameFinErr {
		log.Panicln(err)
	}
	log.Println("GOOD GAME!")
	log.Printf("Your score was %d", snk.Score())
}

func gameLoop(scrn gui.Screen, cvs canvas.Canvas, snk *snake.Snake) error {
	if err := term.Init(); err != nil {
		return err
	}
	defer term.Close()

	startTime := time.Now()
	errs := make(chan error, 1)

	go func() {
		lastGrow := startTime
		lastScreenRefresh := startTime
		for {
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
					errs <- GameFinErr
				}
			}
		}
	}()

	go func() {
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
	}()

	go func() {
		for {
			log.Println("Displaying!")
			if err := scrn.DisplayMatrix(cvs.GetMatrix(), time.Millisecond*50); err != nil {
				errs <- err
			}
			if err := scrn.AllLEDSOff(); err != nil {
				errs <- err
			}
		}
	}()

	for {
		err := <-errs
		if err != nil {
			return err
		}
	}
}
