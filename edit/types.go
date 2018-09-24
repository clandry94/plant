package edit

import (
	"container/list"
	"github.com/clandry94/plant/edit/raw"
	"os"
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

// the core of the sub editor. Only one context exists
// in the realm of a plant editor instance
type Context struct {
	// a circular chain with pointers to buffers
	// each buffer represents an open file (or new file?)
	buffers *list.List

	// the current buffer in focus
	currentBuffer *list.Element
}

// the representation of a location of a point
// in the subeditor
// TODO: implement Cursor coordinate system
type Cursor struct {
	loc location
}

// represents a mark made in the data
type Mark struct {
	name  string
	where Cursor

	// fixed means the location is the character after the actual location
	// instead of living before i.e. if looking at b then not fixed: a|bc
	// fixed: ab|c
	fixed bool
}
