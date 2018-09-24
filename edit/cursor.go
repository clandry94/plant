package edit

import (
	editErrors "github.com/clandry94/plant/edit/errors"
)

type location struct {
	line int
	col  int
}

func NilCursor() Cursor {
	return Cursor{
		loc: location{
			line: 0,
			col:  0,
		},
	}
}

/*
	Gets current line
*/
func (c Cursor) Line() int {
	return int(c.loc.line)
}

/*
	Gets current col
*/
func (c Cursor) Col() int {
	return int(c.loc.col)
}

/*
	Sets the current line of the cursor
*/
func (c *Cursor) SetLine(line int) error {
	if line < 0 {
		return editErrors.IndexOutOfRangeError{1, line}
	}

	c.loc.line = line

	return nil
}

/*
	Sets the current col of the cursor
*/
func (c *Cursor) SetCol(column int) error {
	if column < 0 {
		return editErrors.IndexOutOfRangeError{1, column}
	}

	c.loc.col = column

	return nil
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
	//tmp := Cursor{}
	//c = &mark.where
	//mark.where = tmp
}

/*
	returns true if Cursor is at the mark
*/
func (c Cursor) AtMark(mark *Mark) bool {
	if compareCursor(c, mark.where) == AEqualB {
		return true
	}

	return false
}

/*
	returns true if Cursor is after the mark
*/
func (c Cursor) BeforeMark(mark *Mark) bool {
	if compareCursor(c, mark.where) == ABeforeB {
		return true
	}

	return false
}

/*
	returns true if Cursor is before the mark
	TODO: need Cursor coordinates to calculate this
*/
func (c Cursor) AfterMark(mark *Mark) bool {
	if compareCursor(c, mark.where) == AAfterB {
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

func compareCursor(a Cursor, b Cursor) int {
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

func compareLine(a int, b int) int {
	if a < b {
		return ABeforeB
	}

	if a > b {
		return AAfterB
	}

	return AEqualB
}

func compareCol(a int, b int) int {
	if a < b {
		return ABeforeB
	}

	if a > b {
		return AAfterB
	}

	return AEqualB
}
