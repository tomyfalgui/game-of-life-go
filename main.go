package main

import (
	"fmt"

	"github.com/tomyfalgui/gol-go/board"
)

func main() {
	randomBoard := board.GenerateRandom(20, 30)
	fmt.Print(randomBoard)
}
