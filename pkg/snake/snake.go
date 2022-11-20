package snake

import "log"

const (
	North = "N"
	South = "S"
	East  = "E"
	West  = "W"
)

type Snake struct {
	direction      string
	xLimit, yLimit int
	headCord       []int
	tailCord       [][]int
	moves          int
	grow           bool
}

func NewSnake(headCord []int, tailCord [][]int, xLimit, yLimit int) Snake {
	return Snake{
		direction: West,
		headCord:  headCord,
		tailCord:  tailCord,
		xLimit:    xLimit,
		yLimit:    yLimit,
		moves:     0,
		grow:      false,
	}
}

func (s *Snake) Body() [][]int {
	return append([][]int{s.headCord}, s.tailCord...)
}

func (s *Snake) SetGrow() {
	s.grow = true
}

func (s *Snake) Grow() bool {
	return s.grow
}

func (s *Snake) Println() {
	log.Printf("direction=%s head=%+v tail =%+v", s.direction, s.headCord, s.tailCord)
}

func (s *Snake) UpdatePositions(x, y int) {
	if s.moves%10 != 0 {
		s.tailCord = s.tailCord[:len(s.tailCord)-1]
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
	s.Println()
}

func (s *Snake) Score() int {
	return s.moves
}

func (s *Snake) Dead() bool {
	allCords := append([][]int{s.headCord}, s.tailCord...)
	for i, _ := range allCords {
		for j := i + 1; j < len(allCords); j++ {
			if allCords[i][0] == allCords[j][0] && allCords[i][1] == allCords[j][1] {
				return true
			}
		}
	}
	return false
}

func (s *Snake) NextMove() {
	switch s.direction {
	case North:
		log.Println("Going north!")
		s.UpdatePositions(0, -1)
	case South:
		log.Println("Going south!")
		s.UpdatePositions(0, 1)
	case East:
		log.Println("Going east!")
		s.UpdatePositions(1, 0)
	case West:
		log.Println("Going west!")
		s.UpdatePositions(-1, 0)
	}

}

var possibleDirections = map[string]map[string]bool{
	North: map[string]bool{
		North: true,
		East:  true,
		West:  true,
	},
	South: map[string]bool{
		South: true,
		East:  true,
		West:  true,
	},
	East: map[string]bool{
		North: true,
		East:  true,
		South: true,
	},
	West: map[string]bool{
		North: true,
		South: true,
		West:  true,
	},
}

func (s *Snake) ChangeDirection(direction string) {
	if _, ok := possibleDirections[s.direction][direction]; ok {
		s.direction = direction
	}
}
