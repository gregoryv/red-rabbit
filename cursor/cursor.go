// cursor movement in byte slice
package cursor

import "bytes"

// IndexLeft returns the right most index of the sep character
// to the left of start
func IndexLeft(buf, sep []byte, start int) (i int) {
	if start >= len(buf) {
		start = len(buf) - 1
	}
	if start > 0 && start < len(buf) {
		if i = bytes.LastIndex(buf[:start], sep); i > -1 {
			return i
		}
	}
	return 0
}

// IndexUp moves the cursor one line up from the given index.
// It tries to keep the column position if possible.
func IndexUp(buf, sep []byte, start int) (i int) {
	// make sure we're within the buffer limits
	if start >= len(buf) {
		start = len(buf) - 1
	}
	if start < 0 {
		return 0
	}
	// Move to end of previous line
	end := IndexLeft(buf, sep, start)
	// Which column are we at now?
	currentcol := start - end - 1
	// We are at the first character
	if end == 0 {
		return 0
	}
	// Find beginning of the previous line
	begin := IndexLeft(buf, sep, end)
	// move to the right of the sep so we end up on column 0
	if begin > 0 {
		begin = begin + 1
	}
	endcol := end - begin
	// the above line is empty
	if endcol == 0 {
		return end
	}
	// if the line above is shorter than the current column position
	if endcol < currentcol {
		return begin + endcol
	}
	i = begin + currentcol
	return
}

// IndexDown moves the cursor one line down from the given index.
// It tries to keep the column position if possible.
func IndexDown(buf, sep []byte, start int) (i int) {
	// make sure we're within the buffer limits
	if start >= len(buf) {
		return len(buf) - 1
	}
	if start < 0 {
		start = 0
	}
	currentcol := start - bytes.LastIndex(buf[:start], sep) - 1
	// index of the next new line
	begin := bytes.Index(buf[start:], sep) + start
	if begin > 0 {
		begin = begin + 1
	}
	endcol := bytes.Index(buf[begin:], sep)
	if endcol == -1 {
		endcol = len(buf[begin:]) - 1
	}
	if endcol == 0 { // means we're on empty line
		return begin
	}
	end := begin + endcol
	// line below is shorter than current column position
	if currentcol > endcol {
		return end
	}
	if begin == len(buf)-1 {
		return begin
	}
	i = begin + currentcol
	return
}
