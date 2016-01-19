package cursor

import (
	"testing"
)

var newline = []byte("\n")
var b = []byte("01\n\n4\n67\n9")
var i int

func TestIndex(t *testing.T) {
	data := []struct {
		Row, Column, Expected int
	}{
		{-1, 0, 0},
		{-1, -1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 2, 2},
		{0, 3, 3}, // overflow column goes to next line
		{1, 0, 3},
		{2, 0, 4},
		{2, 1, 5},
		{3, 0, 6},
		{3, 1, 7},
		{3, 2, 8},
		{4, 0, 9},
		{15, 0, 9},   // line overflow
		{4, -1, 8},   // backwards
		{4, -100, 0}, // backwards overflow column
	}
	for _, d := range data {
		if i := Index(b, newline, d.Row, d.Column); i != d.Expected {
			t.Errorf("Index(%v, %v) => %v expected %v", d.Row, d.Column, i, d.Expected)
		}
	}
}

func TestIndexLeft(t *testing.T) {
	data := []struct {
		Start, Expected int
	}{
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 2},
		{4, 3},
		{5, 3},
		{6, 5},
		{7, 5},
		{8, 5},
		{9, 8},
		{98, 8},
	}
	for _, d := range data {
		if i = IndexLeft(b, newline, d.Start); i != d.Expected {
			t.Errorf("IndexLeft(%v) => %v expected %v\n%s",
				d.Start, i, d.Expected, b)
		}
	}
}

func TestIndexUp(t *testing.T) {
	data := []struct {
		Start, Expected int
	}{ // 01\n\n4\n67\n9
		{-1, 0},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 3},
		{5, 3},
		{6, 4},
		{7, 5},
		{8, 5},
		{9, 6},
	}
	for _, d := range data {
		if i = IndexUp(b, newline, d.Start); i != d.Expected {
			msg := "IndexUp(%v) => %v expected %v \n%s"
			t.Errorf(msg, d.Start, i, d.Expected, b)
		}
	}
}

func TestIndexDown(t *testing.T) {
	data := []struct {
		Start, Expected int
	}{ // 01\n\n4\n67\n9
		{-1, 3},
		{0, 3},
		{1, 3},
		{2, 3},
		{3, 4},
		{4, 6},
		{5, 7},
		{6, 9},
		{7, 9},
		{8, 9},
		{9, 9},
		{10, 9},
	}
	for _, d := range data {
		if i = IndexDown(b, newline, d.Start); i != d.Expected {
			msg := "IndexDown(%v) => %v expected %v \n%s"
			t.Errorf(msg, d.Start, i, d.Expected, b)
		}
	}
}

func TestPosition(t *testing.T) {
	data := []struct {
		Start, ExpectedRow, ExpectedColumn int
	}{ // 01\n\n4\n67\n9
		{-1, 0, 0},
		{0, 0, 0},
		{1, 0, 1},
		{2, 0, 2},
		{3, 1, 0},
		{4, 2, 0},
		{5, 2, 1},
		{6, 3, 0},
		{7, 3, 1},
		{8, 3, 2},
		{9, 4, 0},
		{10, 4, 0},
	}
	for _, d := range data {
		if row, col := Position(b, newline, d.Start); row != d.ExpectedRow || col != d.ExpectedColumn {
			t.Errorf("Position(%v) => %v, %v expected %v, %v\n", d.Start, row, col, d.ExpectedRow, d.ExpectedColumn)
		}
	}
	// Test empty buffer
	if row, col := Position(make([]byte, 0), newline, 1); row != 0 || col != 0 {
		t.Errorf("Position() should handle empty buffers")
	}
}
