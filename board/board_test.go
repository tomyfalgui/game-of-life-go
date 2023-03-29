package board_test

import (
	"testing"

	board "github.com/tomyfalgui/gol-go/board"
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

	board := [][]int{
		{0, 1, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}

	want := " #   #\n    #\n ### \n"
	got := board.Render(board)

	if want != got {
		t.Errorf("want %v vs got %v", want, got)
	}
}
