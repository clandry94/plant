package edit

import (
)

// the representation of a location of a point
// in the subeditor
// TODO: implement cursor coordinate system
type cursor struct {
	// create location coordinate system
}

/*
	Sets the cursors location to the mark provided
 */
func (c *cursor) SetToMark(mark *Mark) error {
	return nil
}

/*
	Swaps the location of the cursor and the mark
 */
func (c *cursor) Swap(mark *Mark) {
	// TODO: set to internal coordinates of the cursor
	// this probably won't work right due to pointer things
	tmp := cursor{}
	c = &mark.where
	mark.where = tmp
}

/*
	returns true if cursor is at the mark
 */
func (c cursor) PointAtMark(mark *Mark) bool {
	return c == mark.where
}

/*
	returns true if cursor is before the mark
	TODO: need cursor coordinates to calculate this
 */
func (c cursor) PointAfterMark(mark *Mark) bool {
	return false
}


/*
	returns true if cursor is after the mark
 */
func (c cursor) PointBeforeMark(mark *Mark) bool {
	return false
}

/*
	Some utility methods
 */

const (
	Loc1Before = -1
	Loc1Equal = 0
	Loc1After = 1
)

func CompareLocation(loc1 cursor, loc2 cursor) int {
	// if loc1 before loc 2 return Loc1Before

	// if loc1 same as loc2 return Loc1Equal

	// if loc1 after loc2 return Loc1After

	return Loc1Equal
}


