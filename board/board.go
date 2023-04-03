package board

import (
	"errors"
	"math/rand"
	"strings"
)

func GenerateRandom(width, height int) ([][]int, error) {
	if width < 3 || height < 3 {
		return [][]int{}, errors.New("Grid should at least be 3x3")
	}
	grid := [][]int{}

	for y := 0; y < height; y++ {
		wSlice := []int{}
		for x := 0; x < width; x++ {
			wSlice = append(wSlice, rand.Intn(2))
		}
		grid = append(grid, wSlice)
	}
	return grid, nil
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

// behaviors
/**
1. Any live cell with 0 or 1 live neighbors becomes dead, because of underpopulation
Any live cell with 2 or 3 live neighbors stays alive, because its neighborhood is just right
Any live cell with more than 3 live neighbors becomes dead, because of overpopulation
Any dead cell with exactly 3 live neighbors becomes alive, by reproduction
*/
func NextState(grid [][]int) [][]int {
	nextState := make([][]int, len(grid))
	maxY := len(grid)
	for iy, y := range grid {
		maxX := len(y)
		nextState[iy] = make([]int, len(grid[iy]))
		for ix, cell := range y {
			nIndices := NeighborIndices(ix, iy, maxX, maxY)
			deadCount := 0
			liveCount := 0
			for _, nIdx := range nIndices {
				yNeighbor := nIdx[0]
				xNeighbor := nIdx[1]

				if grid[yNeighbor][xNeighbor] == 0 {
					deadCount++
				} else {
					liveCount++
				}
			}

			if cell == 1 && (liveCount == 2 || liveCount == 3) {
				nextState[iy][ix] = 1
			} else if cell == 0 && (liveCount == 3) {
				nextState[iy][ix] = 1
			} else {
				nextState[iy][ix] = 0
			}
		}

	}

	return nextState
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
