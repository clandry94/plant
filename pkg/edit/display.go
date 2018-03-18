package edit

import "os"

/*
	This describes the interface that must be implemented by a
	redisplay module such that it can be used by the editor.

	This lets us use different editing libraries and UIs as long as
	they implement a specific set of public methods
 */
type Redisplay interface{
	// Perform cleanup operations and close the interface
	Close() error

	// Save the current state of the editor in the .plant_state file format
	Save() error

	// Load an editor state from a .plant_state file
	Load(*os.File) error

	// Performs an incremental refresh of the display. If run,
	// this should make sure the window accurately represents the buffer.
	Redisplay()

	// Performs a full window reload. This makes sure that the screen is correct
	// no matter what.
	Refresh()


	// Current row that the cursor is on in the window. This might be different from
	// the cursor point in the editor due to linewrap
	CursorRow() int

	// Same as CursorRow but for columns
	CursorCol() int

	// Set the location of the cursor in the window
	SetCursor()

	// Set the row of the cursor in the window
	SetRow()

	// Set the col of the cursor in the window
	SetCol()

	// SetAttr()
	// Attr()

	// Clears the current line
	ClearLine()

	// Clears the screen
	ClearScreen()

	// Place a rune at the cursor location
	PutRune(rune)

	// Place a slice of runes at the current location
	PutRunes([]rune)

	// Delete n runes from the start of the cursor
	DeleteRunes(int)

	// Insert n blank lines
	InsertLines(int)

	// Delete n lines
	DeleteLines(int)
}
