package canvas

import "go-snake/pkg/snake"

type Canvas struct {
	width, height int
	snake         *snake.Snake
}

func NewCanvas(width, height int, snake *snake.Snake) Canvas {
	return Canvas{
		width:  width,
		height: height,
		snake:  snake,
	}
}

func (c *Canvas) GetMatrix() [][]int {
	snakeCords := c.snake.Body()

	matrix := [][]int{}
	for y := 0; y < c.height; y++ {
		matrix = append(matrix, []int{})
		for x := 0; x < c.width; x++ {
			matrix[y] = append(matrix[y], 0)
		}
	}

	for _, cord := range snakeCords {
		matrix[cord[1]][cord[0]] = 1
	}
	return matrix
}
