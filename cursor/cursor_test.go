package cursor

import (
	"testing"
)

var newline = []byte("\n")
var b = []byte("01\n\n4\n67\n9")
var i int

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
