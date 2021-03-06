package edit

import (
	"container/list"
	"fmt"
	"github.com/clandry94/plant/edit/raw"
	"github.com/clandry94/plant/edit/status"
)

func (b Buffer) Status() status.Status {
	return status.Status{
		Lines:       b.NumLines(),
		CurrentLine: b.Cursor.Line(),
		Cols:        b.NumCols(),
		CurrentCol:  b.Cursor.Col(),
	}
}

func (b Buffer) GetContents() *raw.Contents {
	return b.contents
}

// Inserts a string at the current cursor location
func (b *Buffer) Insert(str string) {
	p := b.contents.Lines.Front()

	i := 0
	for p != nil {
		if b.Cursor.Line() == i {
			err := p.Value.(*raw.Piece).Insert(b.Cursor.Col(), []rune(str))
			if err != nil {
				log.WithField("module", "buffer").Error(err.Error())
			}

			b.CursorMoveForward(len(str))
		}

		p = p.Next()
		i++
	}
}

// delete a n characters at the current cursor location
func (b *Buffer) Delete(length int) {
	p := b.contents.Lines.Front()

	i := 0
	for p != nil {
		if b.Cursor.Line() == i {
			err := p.Value.(*raw.Piece).Delete(b.Cursor.Col(), length)
			if err != nil {
				log.Warn(err)
			}

			b.CursorMoveBack(length)
		}
		p = p.Next()
		i++
	}
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
	Move buffer cursor up i rows and to the end of the line
	if current post is outside of the length of the above line
*/
func (b *Buffer) CursorMoveUp(i int) error {
	err := b.Cursor.SetLine(b.Cursor.Line() - i)
	if err != nil {
		return err
	}

	b.reelCursor()

	return nil
}

/*
	move buffer cursor down i rows and to the end of the line
	if current pos is outside of the length of the below line
*/
func (b *Buffer) CursorMoveDown(i int) error {
	err := b.Cursor.SetLine(b.Cursor.Line() + i)
	if err != nil {
		return err
	}

	b.reelCursor()

	return nil
}

func (b *Buffer) SetCursorEndOfLine() {
	// for some reason there are 2 extra character at the end of each line.
	// This is probably the newline
	// update: this is from the \n
	b.Cursor.SetCol(b.endOfLineCol() - 2)
}

// reeling the cursor means to bring it back to the end of the current line
// after a line change
func (b *Buffer) reelCursor() {
	curLine := b.currentLine()

	// Handles empty lines
	// TODO: make this nicer
	if curLine.Value.(*raw.Piece).Len() == 1 {
		b.Cursor.SetCol(b.endOfLineCol() - 1)
		return
	}

	if b.Cursor.Col() > curLine.Value.(*raw.Piece).Len()-2 {
		b.SetCursorEndOfLine()
	}
}

// is a *raw.Piece in the value. Returning element
// so we have access to the prev and next
func (b *Buffer) currentLine() *list.Element {
	p := b.contents.Lines.Front()

	i := 0
	for p != nil {
		if b.Cursor.Line() == i {
			return p
		}
		p = p.Next()
		i++
	}

	return nil
}

func (b Buffer) endOfLineCol() int {
	p := b.contents.Lines.Front()

	i := 0
	for p != nil {
		if b.Cursor.Line() == i {
			return len(p.Value.(*raw.Piece).Data())
		}
		p = p.Next()
		i++
	}

	return 0
}

/*
	Move buffer Cursor forward i characters
*/
func (b *Buffer) CursorMoveForward(i int) error {
	err := b.Cursor.SetCol(b.Cursor.Col() + i)
	if err != nil {
		return err
	}

	return nil
}

/*
	Move buffer Cursor backward i characters
*/
func (b *Buffer) CursorMoveBack(i int) error {
	err := b.Cursor.SetCol(b.Cursor.Col() - i)
	if err != nil {
		return err
	}

	return nil
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
	return 1337
}

/*
	Returns number of columns in the current line. Does not count newlie
*/
func (b *Buffer) NumCols() int {
	return 1337
}
