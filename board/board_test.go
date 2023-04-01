package board_test

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestGetNeighborIndices_XYPair(t *testing.T) {
	t.Parallel()
	maxY := 10
	maxX := 5

	currYIndex1 := 0
	currXIndex1 := 0

	// in (Y, X) form
	wantNeighborIndices1 := [][]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{9, 4},
		{9, 0},
		{9, 1},
		{0, 4},
		{1, 4},
	}
	gotNeigborIndices1 := board.NeighborIndices(currXIndex1, currYIndex1, maxX, maxY)
	if !cmp.Equal(
		wantNeighborIndices1,
		gotNeigborIndices1,
		cmpopts.SortSlices(func(x, y [][]int) bool {
			return x[0][0] == y[0][0]
		}),
	) {
		t.Errorf("diff NeighborIndices1: %v", cmp.Diff(wantNeighborIndices1, gotNeigborIndices1))
	}
}

// behaviors
/**
1. Any live cell with 0 or 1 live neighbors becomes dead, because of underpopulation
Any live cell with 2 or 3 live neighbors stays alive, because its neighborhood is just right
Any live cell with more than 3 live neighbors becomes dead, because of overpopulation
Any dead cell with exactly 3 live neighbors becomes alive, by reproduction
*/

func TestNextStateCells_WithZeroOrOneLiveNeighbors_ShouldBeDead(t *testing.T) {
	t.Parallel()

	// base case
	initState1 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	wantState1 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	gotState1 := board.NextState(initState1)
	if !cmp.Equal(wantState1, gotState1) {
		t.Errorf("state1 diff: %v", cmp.Diff(wantState1, gotState1))
	}

	initState2 := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{1, 1, 0},
	}
	wantState2 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	gotState2 := board.NextState(initState2)
	if !cmp.Equal(wantState2, gotState2) {
		t.Errorf("state2 diff: %v", cmp.Diff(wantState2, gotState2))
	}
}
