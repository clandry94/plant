package edit

import (
	"container/list"
	"os"
	"fmt"
)

// the core of file represented in the sub editor
type Buffer struct {
	name string

	// current location of the cursor
	cursor cursor

	// maintains a set of marks
	Marks Marks

	// holds the raw content of the buffer
	contents contents

	fileInfo os.FileInfo

	dirty bool

	// individual modes appended to the buffer
	modes *list.List
}

/*
	Sets the name of the buffer
 */
func (b *Buffer) SetName(name string) error {
	return nil
}

/*
	Returns the name of the buffer
 */
func (b *Buffer) GetName() string {
	return b.name
}

/*
	Sets cursor in a buffer to the location provided
 */
func (b *Buffer) SetCursor(loc cursor) error {
	return nil
}

/*
	Sets the cursor to i characters from the buffer start
 */
func (b *Buffer) SetCursorToCount(i int) error {
	loc, err := b.CountToLocation(i)
	if err != nil {
		return err
	}

	return b.SetCursor(loc)
}

/*
	Move buffer cursor forward i characters
 */
func (b *Buffer) CursorMoveForward(i int) error {
	return nil
}

/*
	Move buffer cursor backward i characters
 */
func (b *Buffer) CursorMoveBack(i int) error {
	return nil
}

/*
	Retrieve the location of the buffer's cursor
 */
func (b *Buffer) GetCursor() cursor {
	return b.cursor
}

/*
	Number of characters between the start of the buffer
	and the location provided
 */
func (b *Buffer) CountFromBufStart(loc cursor) int {
	return 0
}

/*
	Percentage from the start of the buffer that the location arg is at
 */
func (b *Buffer) PercentFromBufStart(loc cursor) float64 {
	// return (b.CountFromBufStart(b.GetCursor()) * 100) / GetNumChars()
	return 0.0
}

/*
	Accepts a non negative count and converts it to a location in the buffer
 */
func (b *Buffer) CountToLocation(i int) (cursor, error) {
	if i < 0 {
		return cursor{}, fmt.Errorf("negative count")
	}

	return cursor{}, nil
}

// gets the location at the start of the buffer (might not be 0)
func (b *Buffer) BufferStart() cursor {
	return cursor{}
}

// gets the location at the end of the buffer (might not be the end)
func (b *Buffer) BufferEnd() cursor {
	return cursor{}
}

/*
	Returns the rune after the cursor. Error if end of buffer
 */
func (b *Buffer) GetRune() (rune, error) {
	return 'a', nil
}

/*
	returns the string of runes after a point.
	Errors if the end of the buffer is reached, but still returns the partial string
 */
func (b *Buffer) GetString(count int) (string, error) {
	return "", nil
}

/*
	Returns the number of runes in the buffer (length of buffer)
	TODO: keep track of number of characters because counting num
	characters is O(n^2)
 */
func (b *Buffer) NumRunes() int {
	return 0
}

/*
	Returns number of lines in the buffer. Counts the last line
 */
func (b *Buffer) NumLines() int {
	return 0
}

// the data of the buffer, represented as a rectangle
type contents struct {
	data [][]rune
}

