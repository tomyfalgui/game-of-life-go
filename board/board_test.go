package board_test

import (
	"testing"

	board "github.com/tomyfalgui/game-of-life-go/board"
)

func TestGenerateXYGridState(t *testing.T) {
	t.Parallel()

	wantWidth := 5
	wantHeight := 2
	newBoard := board.GenerateRandom(wantWidth, wantHeight)

	gotHeight := len(newBoard)

	if wantHeight != gotHeight {
		t.Errorf("grid height is unmatched. want %d vs got %d", wantHeight, gotHeight)
	}

	for _, list := range newBoard {
		gotWidth := len(list)
		if gotWidth != wantWidth {
			t.Errorf("grid width is unmatched. want %d vs got %d", wantWidth, gotWidth)
		}
	}
}

func TestRenderGrid(t *testing.T) {
	t.Parallel()

	testBoard := [][]int{
		{0, 1, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}

	want := " #  #\n    #\n ### \n"
	got := board.Render(testBoard)

	if want != got {
		t.Errorf("want %v vs got %v", want, got)
	}
}

// behaviors
/**
Any live cell with 0 or 1 live neighbors becomes dead, because of underpopulation
Any live cell with 2 or 3 live neighbors stays alive, because its neighborhood is just right
Any live cell with more than 3 live neighbors becomes dead, because of overpopulation
Any dead cell with exactly 3 live neighbors becomes alive, by reproduction
*/
