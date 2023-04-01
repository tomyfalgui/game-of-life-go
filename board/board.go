package board

import (
	"math/rand"
	"strings"
)

func GenerateRandom(width, height int) [][]int {
	grid := [][]int{}

	for y := 0; y < height; y++ {
		wSlice := []int{}
		for x := 0; x < width; x++ {
			wSlice = append(wSlice, rand.Intn(2))
		}
		grid = append(grid, wSlice)
	}
	return grid
}

func Render(grid [][]int) string {
	var sb strings.Builder

	for _, y := range grid {
		for _, x := range y {
			if x == 0 {
				sb.WriteString(" ")
			}

			if x == 1 {
				sb.WriteString("#")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func NextState(grid [][]int) [][]int {
	// yDimMax := len(grid)
	// xDimMax := len(grid[0])
	// nextState := [][]int{}

	return grid
}

/*
*
(x-1, y-1)
(x, y-1)
(x+1, y-1)
(x-1, y)
(x+1, y)
(x-1, y+1)
(x, y+1)
(x+1, y+1)
*/
func NeighborIndices(x, y, maxX, maxY int) [][]int {
	neighbors := make([][]int, 8)

	// y + 1
	// x + 1
	// y - 1
	// x - 1
	yInc := y + 1
	xInc := x + 1
	yDec := y - 1
	xDec := x - 1

	if yInc >= maxY {
		yInc = 0
	}
	if xInc >= maxX {
		xInc = 0
	}
	if yDec < 0 {
		yDec = maxY - 1
	}
	if xDec < 0 {
		xDec = maxX - 1
	}

	neighbors[0] = []int{yDec, xDec}
	neighbors[1] = []int{yDec, x}
	neighbors[2] = []int{yDec, xInc}
	neighbors[3] = []int{y, xDec}
	neighbors[4] = []int{y, xInc}
	neighbors[5] = []int{yInc, xDec}
	neighbors[6] = []int{yInc, x}
	neighbors[7] = []int{yInc, xInc}

	return neighbors
}
