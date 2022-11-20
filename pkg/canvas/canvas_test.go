// Package gui is responsible for visual output
// File gui_test.go tests the gui.go file
package canvas

import (
	"go-snake/pkg/snake"
	"testing"

	"github.com/stretchr/testify/suite"
)

type canvasSuite struct {
	suite.Suite
}

func (*canvasSuite) SetupTest() {}

func TestQuizTestSuite(s *testing.T) {
	suite.Run(s, new(canvasSuite))
}

func (c *canvasSuite) Test_NewCanvas() {
	cnvs := NewCanvas(10, 10, &snake.Snake{})
	expected := Canvas{
		width:  10,
		height: 10,
		snk:    &snake.Snake{},
	}
	c.Equal(expected, cnvs)
}

func (c *canvasSuite) Test_GetMatrix_A() {
	snk := snake.NewSnake([]int{0, 0}, [][]int{{1, 0}, {2, 0}}, 5, 5)
	cnvs := NewCanvas(5, 5, &snk)
	matrix := cnvs.GetMatrix()
	expexted := [][]int{
		{1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	c.Equal(expexted, matrix)
}

func (c *canvasSuite) Test_GetMatrix_B() {
	snk := snake.NewSnake([]int{3, 0}, [][]int{{3, 1}, {3, 2}}, 5, 5)
	cnvs := NewCanvas(5, 5, &snk)
	matrix := cnvs.GetMatrix()
	expexted := [][]int{
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	c.Equal(expexted, matrix)
}
