// Package canvas is regarding the thing the snake is placed on
package canvas

import "github.com/gowhale/go-snake/pkg/snake"

// Canvas represents the canvas which the snake is placed
type Canvas struct {
	width, height int
	snk           *snake.Snake
}

// NewCanvas returns a Canvas based on the height, width and snake passed
func NewCanvas(width, height int, snk *snake.Snake) Canvas {
	return Canvas{
		width:  width,
		height: height,
		snk:    snk,
	}
}

// GetMatrix returns a matrix made up of 1 and 0
// 1's mean the snake's body lays there. BEWARE!
func (c *Canvas) GetMatrix() [][]int {
	matrix := [][]int{}
	for y := 0; y < c.height; y++ {
		matrix = append(matrix, []int{})
		for x := 0; x < c.width; x++ {
			matrix[y] = append(matrix[y], 0)
		}
	}

	snakeCords := c.snk.Body()
	for _, cord := range snakeCords {
		matrix[cord[1]][cord[0]] = 1
	}
	return matrix
}
