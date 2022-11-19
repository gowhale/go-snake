package snake

type Snake struct {
	direction string
	headCord  []int
	tailCord  [][]int
}

func NewSnake(headCord []int, tailCord [][]int) Snake {
	return Snake{
		headCord: headCord,
		tailCord: tailCord,
	}
}

func (s *Snake) Body() [][]int {
	return append([][]int{s.headCord}, s.tailCord...)
}
