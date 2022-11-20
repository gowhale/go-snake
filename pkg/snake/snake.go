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
}

func NewSnake(headCord []int, tailCord [][]int, xLimit, yLimit int) Snake {
	return Snake{
		direction: North,
		headCord:  headCord,
		tailCord:  tailCord,
		xLimit:    xLimit,
		yLimit:    yLimit,
	}
}

func (s *Snake) Body() [][]int {
	return append([][]int{s.headCord}, s.tailCord...)
}

func (s *Snake) UpdatePositions(x, y int) {
	s.tailCord = s.tailCord[:len(s.tailCord)-1]
	s.tailCord = append([][]int{s.headCord}, s.tailCord...)
	headX, headY := s.headCord[0]+x, s.headCord[1]+y
	if headX < 0 {
		log.Println("SETTING X LIM")
		headX = s.xLimit - 1
	}
	if headY < 0 {
		log.Println("SETTING Y LIM")
		headY = s.yLimit - 1
		log.Println(s.yLimit - 1)
		log.Println(headY)
		s.headCord[1] = headY
	}
	if headX > s.xLimit-1 {
		headX = 0
	}
	if headY > s.yLimit-1 {
		headY = 0
	}
	log.Println(headX, headY)
	s.headCord = []int{headX, headY}
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
