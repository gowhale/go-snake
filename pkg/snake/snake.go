// Package snake contains code to do with the Snake 'character'
package snake

import "log"

const (
	North = "N" // North means the snake is going up the matrix
	South = "S" // South means the snake is down up the matrix
	East  = "E" // East means the snake is right accross the matrix
	West  = "W" // West means the snake is left accross the matrix
)

// Snake represents the game character
// Has attributes such as direction and position etc.
type Snake struct {
	direction      string
	xLimit, yLimit int
	headCord       []int
	tailCord       [][]int
	moves          int
	grow           bool
}

// NewSnake creates a new instance of snake to be place on Canvas
func NewSnake(headCord []int, tailCord [][]int, xLimit, yLimit int) Snake {
	return Snake{
		direction: North,
		headCord:  headCord,
		tailCord:  tailCord,
		xLimit:    xLimit,
		yLimit:    yLimit,
		moves:     0,
		grow:      false,
	}
}

// Body returns all coords the snake is at
// tail and head!
func (s *Snake) Body() [][]int {
	return append([][]int{s.headCord}, s.tailCord...)
}

// SetGrow marks the snake as ready to grow
func (s *Snake) SetGrow(g bool) {
	s.grow = g
}

// GetGrow returns if the snake needs to grow
func (s *Snake) GetGrow() bool {
	return s.grow
}

func (s *Snake) println() {
	log.Printf("direction=%s head=%+v tail=%+v", s.direction, s.headCord, s.tailCord)
}

func (s *Snake) updatePositions(x, y int) {
	if !s.GetGrow() {
		s.tailCord = s.tailCord[:len(s.tailCord)-1]
	} else {
		s.SetGrow(false)
	}
	s.moves++
	s.tailCord = append([][]int{{s.headCord[0], s.headCord[1]}}, s.tailCord...)
	headX, headY := s.headCord[0]+x, s.headCord[1]+y
	if headX < 0 {
		headX = s.xLimit - 1
	}
	if headY < 0 {
		headY = s.yLimit - 1
		s.headCord[1] = headY
	}
	if headX > s.xLimit-1 {
		headX = 0
	}
	if headY > s.yLimit-1 {
		headY = 0
	}
	s.headCord = []int{headX, headY}
}

// Score returns the snakes score
func (s *Snake) Score() int {
	return s.moves
}

// Dead returns if the snake is dead as it has collided with itself
func (s *Snake) Dead() bool {
	allCords := s.Body()
	for i := range allCords {
		for j := i + 1; j < len(allCords); j++ {
			if allCords[i][0] == allCords[j][0] && allCords[i][1] == allCords[j][1] {
				return true
			}
		}
	}
	return false
}

// NextMove updates the snakes cords based on direction
func (s *Snake) NextMove() {
	switch s.direction {
	case North:
		s.updatePositions(0, -1)
	case South:
		s.updatePositions(0, 1)
	case East:
		s.updatePositions(1, 0)
	case West:
		s.updatePositions(-1, 0)
	}
}

var possibleDirections = map[string]map[string]bool{
	North: {
		North: true,
		East:  true,
		West:  true,
	},
	South: {
		South: true,
		East:  true,
		West:  true,
	},
	East: {
		North: true,
		East:  true,
		South: true,
	},
	West: {
		North: true,
		South: true,
		West:  true,
	},
}

// ChangeDirection changes the direction the snake is heading
// Note: Snake cannot turn in opposite way
func (s *Snake) ChangeDirection(direction string) {
	if _, ok := possibleDirections[s.direction][direction]; ok {
		s.direction = direction
	}
}
