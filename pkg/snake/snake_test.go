// Package gui is responsible for visual output
// File gui_test.go tests the gui.go file
package snake

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type snakeSuite struct {
	suite.Suite
}

func (*snakeSuite) SetupTest() {}

func TestQuizTestSuite(s *testing.T) {
	suite.Run(s, new(snakeSuite))
}

func (s *snakeSuite) Test_Grow() {
	snk := Snake{
		grow: false,
	}
	snk.SetGrow()
	s.Equal(true, snk.Grow())
}

func (s *snakeSuite) Test_Dead_True() {
	snk := Snake{
		headCord: []int{1, 2},
		tailCord: [][]int{{2, 2}, {3, 2}, {4, 2}, {2, 2}},
	}
	s.True(snk.Dead())
}

func (s *snakeSuite) Test_Dead_False() {
	snk := Snake{
		headCord: []int{1, 2},
		tailCord: [][]int{{2, 2}, {3, 2}, {4, 2}, {5, 2}},
	}
	s.False(snk.Dead())
}

func (s *snakeSuite) Test_NextMove_North_Middle() {
	snk := Snake{
		direction: North,
		headCord:  []int{1, 2},
		tailCord:  [][]int{{2, 2}, {3, 2}, {4, 2}, {5, 2}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.NextMove()
	expectedSnake := Snake{
		direction: North,
		headCord:  []int{1, 1},
		tailCord:  [][]int{{1, 2}, {2, 2}, {3, 2}, {4, 2}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NextMove_North_Top() {
	snk := Snake{
		direction: North,
		headCord:  []int{1, 0},
		tailCord:  [][]int{{1, 1}, {1, 2}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.NextMove()
	expectedSnake := Snake{
		direction: North,
		headCord:  []int{1, 9},
		tailCord:  [][]int{{1, 0}, {1, 1}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NextMove_South_Middle() {
	snk := Snake{
		direction: South,
		headCord:  []int{1, 3},
		tailCord:  [][]int{{1, 2}, {1, 1}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.NextMove()
	expectedSnake := Snake{
		direction: South,
		headCord:  []int{1, 4},
		tailCord:  [][]int{{1, 3}, {1, 2}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NextMove_South_Bottom() {
	snk := Snake{
		direction: South,
		headCord:  []int{1, 9},
		tailCord:  [][]int{{1, 8}, {1, 7}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.NextMove()
	expectedSnake := Snake{
		direction: South,
		headCord:  []int{1, 0},
		tailCord:  [][]int{{1, 9}, {1, 8}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NextMove_East_Middle() {
	snk := Snake{
		direction: East,
		headCord:  []int{1, 3},
		tailCord:  [][]int{{1, 2}, {1, 1}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.NextMove()
	expectedSnake := Snake{
		direction: East,
		headCord:  []int{2, 3},
		tailCord:  [][]int{{1, 3}, {1, 2}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NextMove_East_Edge() {
	snk := Snake{
		direction: East,
		headCord:  []int{9, 0},
		tailCord:  [][]int{{8, 0}, {7, 0}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.NextMove()
	expectedSnake := Snake{
		direction: East,
		headCord:  []int{0, 0},
		tailCord:  [][]int{{9, 0}, {8, 0}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NextMove_West_Middle() {
	snk := Snake{
		direction: West,
		headCord:  []int{1, 3},
		tailCord:  [][]int{{1, 2}, {1, 1}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.NextMove()
	expectedSnake := Snake{
		direction: West,
		headCord:  []int{0, 3},
		tailCord:  [][]int{{1, 3}, {1, 2}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NextMove_West_Edge() {
	snk := Snake{
		direction: West,
		headCord:  []int{0, 0},
		tailCord:  [][]int{{1, 0}, {2, 0}},
		moves:     1,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.Println()
	snk.NextMove()
	expectedSnake := Snake{
		direction: West,
		headCord:  []int{9, 0},
		tailCord:  [][]int{{0, 0}, {1, 0}},
		moves:     2,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}
	snk.Println()

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_NewSnake() {
	snk := NewSnake([]int{0, 0}, [][]int{{1, 0}, {2, 0}}, 10, 10)

	expectedSnake := Snake{
		direction: North,
		headCord:  []int{0, 0},
		tailCord:  [][]int{{1, 0}, {2, 0}},
		moves:     0,
		xLimit:    10,
		yLimit:    10,
		grow:      false,
	}

	s.Equal(expectedSnake, snk)
}

func (s *snakeSuite) Test_Score() {
	snk := Snake{
		moves: 99,
	}

	s.Equal(99, snk.Score())
}

func (s *snakeSuite) Test_ChangeDirection() {
	snk := Snake{
		direction: North,
		moves:     99,
	}
	
	snk.ChangeDirection(East)
	snk.ChangeDirection(South)
	snk.ChangeDirection(West)

	s.Equal(West, snk.direction)

	// Test that snake cannot go back on itself
	snk.ChangeDirection(East)
	s.Equal(West, snk.direction)
}
