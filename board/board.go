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
