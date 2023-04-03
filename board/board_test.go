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
	wantHeight := 3
	newBoard, err := board.GenerateRandom(wantWidth, wantHeight)
	if err != nil {
		t.Errorf("board should not have thrown an error")
	}

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

func TestGenerateInvalidXYGrid(t *testing.T) {
	t.Parallel()

	_, err := board.GenerateRandom(2, 1)
	if err == nil {
		t.Errorf("grid should not have been created")
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
		{9, 4},
		{9, 0},
		{9, 1},
		{0, 4},
		{0, 1},
		{1, 4},
		{1, 0},
		{1, 1},
	}
	gotNeigborIndices1 := board.NeighborIndices(currXIndex1, currYIndex1, maxX, maxY)
	if !cmp.Equal(
		wantNeighborIndices1,
		gotNeigborIndices1,
		cmpopts.SortSlices(func(x, y [][]int) bool {
			return x[0][1] < y[0][0]
		}),
	) {
		t.Errorf("diff NeighborIndices1: %v", cmp.Diff(wantNeighborIndices1, gotNeigborIndices1))
	}
}

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
		{0, 1, 0},
		{0, 0, 0},
		{1, 1, 0},
	}
	gotState2 := board.NextState(initState2)

	if !cmp.Equal(wantState2, gotState2) {
		t.Errorf("state2 diff: %v", cmp.Diff(wantState2, gotState2))
	}
}
