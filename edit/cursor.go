package edit

// the representation of a location of a point
// in the subeditor
// TODO: implement Cursor coordinate system
type Cursor struct {
	loc location
}

type line int

type col int

type location struct {
	line line
	col  col
}

/*
	Sets the Cursors location to the mark provided
*/
func (c *Cursor) SetToMark(mark *Mark) error {
	c.loc = mark.where.loc
	return nil
}

/*
	Swaps the location of the Cursor and the mark
*/
func (c *Cursor) Swap(mark *Mark) {
	// TODO: set to internal coordinates of the Cursor
	// this probably won't work right due to pointer things
	tmp := Cursor{}
	c = &mark.where
	mark.where = tmp
}

/*
	returns true if Cursor is at the mark
*/
func (c Cursor) CursorAtMark(mark *Mark) bool {
	if CompareCursor(c, mark.where) == AEqualB {
		return true
	}

	return false
}

/*
	returns true if Cursor is after the mark
*/
func (c Cursor) CursorBeforeMark(mark *Mark) bool {
	if CompareCursor(c, mark.where) == ABeforeB {
		return true
	}

	return false
}

/*
	returns true if Cursor is before the mark
	TODO: need Cursor coordinates to calculate this
*/
func (c Cursor) CursorAfterMark(mark *Mark) bool {
	if CompareCursor(c, mark.where) == AAfterB {
		return true
	}

	return false
}

/*
	Some utility methods
*/

const (
	ABeforeB = -1
	AEqualB  = 0
	AAfterB  = 1
)

func CompareCursor(a Cursor, b Cursor) int {
	return compareLocation(a.loc, b.loc)
}

func compareLocation(a location, b location) int {
	if compareLine(a.line, b.line) == ABeforeB {
		return ABeforeB
	} else if compareLine(a.line, b.line) == AEqualB && compareCol(a.col, b.col) == ABeforeB {
		return ABeforeB
	}

	if compareLine(a.line, b.line) == AAfterB {
		return AAfterB
	} else if compareLine(a.line, b.line) == AEqualB && compareCol(a.col, b.col) == AAfterB {
		return AAfterB
	}

	// cursors must be equal
	return AEqualB
}

func compareLine(a line, b line) int {
	if a < b {
		return ABeforeB
	}

	if a > b {
		return AAfterB
	}

	return AEqualB
}

func compareCol(a col, b col) int {
	if a < b {
		return ABeforeB
	}

	if a > b {
		return AAfterB
	}

	return AEqualB
}
