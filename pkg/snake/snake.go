package snake

import "log"

const (
	North = "N"
	South = "S"
	East  = "E"
	West  = "W"
)

type Snake struct {
	direction string
	headCord  []int
	tailCord  [][]int
}

func NewSnake(headCord []int, tailCord [][]int) Snake {
	return Snake{
		direction: North,
		headCord:  headCord,
		tailCord:  tailCord,
	}
}

func (s *Snake) Body() [][]int {
	return append([][]int{s.headCord}, s.tailCord...)
}

func (s *Snake) NextMove() {
	switch s.direction {
	case North:
		log.Println("Going north!")
		s.tailCord = s.tailCord[:len(s.tailCord)-1]
		s.tailCord = append([][]int{s.headCord}, s.tailCord...)
		s.headCord = []int{s.headCord[0], s.headCord[1] - 1}
	case South:
		log.Println("Going south!")
	case East:
		log.Println("Going east!")
	case West:
		log.Println("Going west!")
	}
}
