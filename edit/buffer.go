package edit

import (
	"container/list"
	"fmt"
	"os"
	"github.com/clandry94/plant/edit/raw"
)

// the core of file represented in the sub editor
type Buffer struct {
	name string

	// current location of the Cursor
	Cursor Cursor

	// maintains a set of marks
	Marks Marks

	// holds the raw content of the buffer
	contents *raw.Contents

	file *os.File

	dirty bool

	// individual modes appended to the buffer
	modes *list.List
}

/*
	Sets the name of the buffer
	TODO: handle name collisions
*/
func (b *Buffer) SetName(name string) error {
	b.name = name
	return nil
}

/*
	Returns the name of the buffer
*/
func (b *Buffer) Name() string {
	return b.name
}

/*
	Sets Cursor in a buffer to the location provided
*/
func (b *Buffer) SetCursor(loc Cursor) error {
	b.Cursor = loc
	return nil
}

/*
	Sets the Cursor to i characters from the buffer start
*/
func (b *Buffer) SetCursorToCount(i int) error {
	loc, err := b.CountToLocation(i)
	if err != nil {
		return err
	}

	return b.SetCursor(loc)
}

/*
	Move buffer Cursor forward i characters
*/
func (b *Buffer) CursorMoveForward(i int) error {
	return b.Cursor.SetCol(b.Cursor.Col() + i)
}

/*
	Move buffer Cursor backward i characters
*/
func (b *Buffer) CursorMoveBack(i int) error {
	return b.Cursor.SetCol(b.Cursor.Col() - i)
}

/*
	Retrieve the location of the buffer's Cursor
*/
func (b *Buffer) GetCursor() Cursor {
	return b.Cursor
}

/*
	Number of characters between the start of the buffer
	and the location provided
*/
func (b *Buffer) CountFromBufStart(loc Cursor) int {
	return 0
}

/*
	Percentage from the start of the buffer that the location arg is at
*/
func (b *Buffer) PercentFromBufStart(loc Cursor) float64 {
	// return (b.CountFromBufStart(b.GetCursor()) * 100) / GetNumChars()
	return 0.0
}

/*
	Accepts a non negative count and converts it to a location in the buffer
*/
func (b *Buffer) CountToLocation(i int) (Cursor, error) {
	if i < 0 {
		return Cursor{}, fmt.Errorf("negative count")
	}

	return Cursor{}, nil
}

// gets the location at the start of the buffer (might not be 0)
func (b *Buffer) BufferStart() Cursor {
	return Cursor{}
}

// gets the location at the end of the buffer (might not be the end)
func (b *Buffer) BufferEnd() Cursor {
	return Cursor{}
}

/*
	Returns the rune after the Cursor. Error if end of buffer
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



