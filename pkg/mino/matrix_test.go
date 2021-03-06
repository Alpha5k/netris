package mino

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	t.Parallel()

	m, err := NewTestMatrix()
	if err != nil {
		t.Error(err)
	}

	m.AddTestBlocks()

	ok := m.SetBlock(9, 0, BlockSolidT, false)
	if !ok {
		t.Error("failed to set final block after test blocks")
	}
	ok = m.SetBlock(9, 1, BlockSolidT, false)
	if !ok {
		t.Error("failed to set final block after test blocks")
	}
	ok = m.SetBlock(9, 3, BlockSolidT, false)
	if !ok {
		t.Error("failed to set final block after test blocks")
	}

	cleared := m.clearFilled()
	if cleared != 3 {
		t.Errorf("failed to clear lines, wanted 3 got %d", cleared)
	}

	m.Clear()

	err = m.Add(m.P, BlockSolidJ, Point{3, 3}, false)
	if err != nil {
		t.Errorf("failed to add initial mino to matrix: %s", err)
	}

	err = m.Add(m.P, BlockSolidJ, Point{3, 3}, false)
	if err == nil {
		t.Error("failed to detect collision when adding second mino to matrix")
	}

	err = m.Add(m.P, BlockSolidJ, Point{9, 9}, false)
	if err == nil {
		t.Error("failed to detect out of bounds when adding third mino to matrix")
	}

	m.Clear()

	for i := 0; i < 11; i++ {
		ok := m.RotatePiece(1, 0)
		if !ok {
			t.Errorf("failed to rotate piece on iteration %d", i)
		}
	}

	for i := 0; i < 4; i++ {
		ok := m.movePiece(1, 0)
		if !ok {
			t.Errorf("failed to Move piece on iteration %d", i)
		}
	}

	ok = m.RotatePiece(1, 0)
	if !ok {
		t.Errorf("failed to rotate piece for right 1wall kick")
	}

	for i := 0; i < 7; i++ {
		ok := m.movePiece(-1, 0)
		if !ok {
			t.Errorf("failed to Move piece on iteration %d", i)
		}
	}

	ok = m.addGarbage(1)
	if !ok {
		t.Error("failed to add 1 line of garbage")
	}

	ok = m.addGarbage(3)
	if !ok {
		t.Error("failed to add 3 line of garbage")
	}
}
