// cursor movement in byte slice
package cursor

// Index returns the index within the buffer for a given column and line
// If column overflows in either positive or negative buffer length determins endpoints.
func Index(buf []rune, sep rune, line, column int) (i int) {
	if line < 0 { // don't handle negative lines
		return 0
	}
FINDLINE:
	for i = 0; i < len(buf); i++ {
		if line == 0 {
			break FINDLINE
		}
		if sep == buf[i] { // line ending found
			line--
		}
	}
	if i == len(buf) { // We reached the end
		i--
	}
	i = i + column
	if i < 0 {
		i = 0
	}
	return
}

// IndexLeft returns the right most index of the sep character
// to the left of start
func IndexLeft(buf []rune, sep rune, start int) (i int) {
	if start >= len(buf) {
		start = len(buf) - 1
	}
	if start > 0 && start < len(buf) {
		return IndexLast(buf[:start], sep)
	}
	return 0
}

// IndexLast returns the last index of sep inside buf or 0 if none is found
// Similar to IndexLeft but works on buf alone.
func IndexLast(buf []rune, sep rune) (i int) {
	for i := len(buf) - 1; i > 0; i-- {
		if sep == buf[i] {
			return i
		}
	}
	return 0
}

// IndexRune returns the index of the first position of sep in buf
func IndexRune(buf []rune, sep rune) (i int) {
	for i := 0; i < len(buf); i++ {
		if sep == buf[i] {
			return i
		}
	}
	return len(buf) - 1
}

// IndexUp moves the cursor one line up from the given index.
// It tries to keep the column position if possible.
func IndexUp(buf []rune, sep rune, start int) (i int) {
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
func IndexDown(buf []rune, sep rune, start int) (i int) {
	// make sure we're within the buffer limits
	if start >= len(buf) {
		return len(buf) - 1
	}
	if start < 0 {
		start = 0
	}
	newlineleft := IndexLast(buf[:start], sep)
	currentcol := start - newlineleft - 1
	// index of the next new line
	begin := IndexRune(buf[start:], sep) + start
	if begin > 0 {
		begin = begin + 1
	}
	endcol := IndexRune(buf[begin:], sep)
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

// Position returns the line and column the given index is on using sep as newline separator
func Position(buf []rune, sep rune, start int) (line, column int) {
	if start >= len(buf) {
		start = len(buf) - 1
	}
	if start < 0 {
		return
	}
	line = Count(buf[:start], sep)
	if line == 0 {
		column = start
		return
	}
	column = start - IndexLast(buf[:start], sep) - 1
	return
}

// Count returns the number of occurences of sep in buf
func Count(buf []rune, sep rune) (c int) {
	for i := 0; i < len(buf); i++ {
		if sep == buf[i] {
			c++
		}
	}
	return
}
