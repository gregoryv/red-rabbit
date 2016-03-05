package cursor

/*
	Moving cursor around a []rune.

	Index values start at 0, while line and columns start with 1.
*/
import (
	"testing"
)

var nlr = '\n'
var b = []byte("ab\n\nö\ncd\nf")
var r = []rune(string(b))

/*
	The indexes are
	column: 123456
	---------------
	line 1: 012
	     2: 3
	     3: 45
	     4: 678
	     5: 9
	---------------
*/

var i int

func TestSetup(t *testing.T) {
	expected := 11
	if len(b) != expected {
		t.Errorf("len(b) should be %v, but is %v", expected, len(b))
	}
	expected = 10
	if len(r) != expected {
		t.Errorf("len(r) should be %v, but is %v", expected, len(r))
	}
}

func TestIndex(t *testing.T) {
	data := []struct {
		Row, Column, Expected int
	}{
		{-1, 0, 0},
		{-1, -1, 0},
		{1, 1, 0},
		{1, 2, 1},
		{1, 3, 2},
		{1, 4, 3}, // overflow column goes to next line
		{2, 1, 3},
		{3, 1, 4},
		{3, 2, 5},
		{4, 1, 6},
		{4, 2, 7},
		{4, 3, 8},
		{5, 1, 9},
		{16, 1, 9},   // line overflow
		{5, -1, 8},   // backwards
		{5, -100, 0}, // backwards overflow column
	}
	for _, d := range data {
		if i := Index(r[:], nlr, d.Row, d.Column); i != d.Expected {
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
		{44, 8},
	}
	for _, d := range data {
		if i = IndexLeft(r, nlr, d.Start); i != d.Expected {
			t.Errorf("IndexLeft(%v) => %v expected %v\n%s",
				d.Start, i, d.Expected, b)
		}
	}
}

func TestIndexLast(t *testing.T) {
	data := []struct {
		End, Expected int
	}{
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
	}
	for _, d := range data {
		if i = IndexLast(r[:d.End], nlr); i != d.Expected {
			t.Errorf("IndexLast(buf[:%v], newline) => %v expected %v\n%s",
				d.End, i, d.Expected, b)
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
		if i = IndexUp(r, nlr, d.Start); i != d.Expected {
			msg := "IndexUp(%v) => %v expected %v \n%s"
			t.Errorf(msg, d.Start, i, d.Expected, b)
		}
	}
}

func TestIndexDown(t *testing.T) {
	data := []struct {
		Start, Expected int
	}{
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
		{29, 9},
	}
	for _, d := range data {
		if i = IndexDown(r, nlr, d.Start); i != d.Expected {
			msg := "IndexDown(%v) => %v expected %v \n%s"
			t.Errorf(msg, d.Start, i, d.Expected, b)
		}
	}
}

func TestPosition(t *testing.T) {
	data := []struct {
		Start, ExpectedRow, ExpectedColumn int
	}{
		{-1, 1, 1},
		{0, 1, 1},
		{1, 1, 2},
		{2, 1, 3}, // newline
		{3, 2, 1}, // newline
		{4, 3, 1},
		{5, 3, 2},
		{6, 4, 1},
		{7, 4, 2},
		{8, 4, 3},
		{9, 5, 1},
		{10, 5, 1},
	}
	for _, d := range data {
		if row, col := Position(r[:], nlr, d.Start); row != d.ExpectedRow || col != d.ExpectedColumn {
			t.Errorf("Position(%v) => %v, %v expected %v, %v\n", d.Start, row, col, d.ExpectedRow, d.ExpectedColumn)
		}
	}
	// Test empty buffer
	if row, col := Position(make([]rune, 0), nlr, 1); row != 1 || col != 1 {
		t.Errorf("Position() should handle empty buffers")
	}
}

func TestCount(t *testing.T) {
	data := []struct {
		Char     rune
		Expected int
	}{
		{nlr, 4},
		{'a', 1},
		{'b', 1},
		{'ö', 1},
		{'c', 1},
		{'d', 1},
		{'f', 1},
		{'n', 0},
	}
	for _, d := range data {
		result := Count(r[:], d.Char)
		if result != d.Expected {
			t.Errorf("Count('%v') => %v expected %v", d.Char, result, d.Expected)
		}
	}
}
