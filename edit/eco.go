package edit

import (
	"container/ring"
	"container/list"
	"os"
)

// the core of the sub editor. Only one ecosystem exists
// in the realm of a plant editor instance, but there could be
// more than one ecosystem in the future
type Ecosystem struct {
	// a circular chain with pointers to buffers
	// each buffer represents an open file (or new file?)
	buffers *ring.Ring

	// the current buffer in focus
	currentBuffer *buffer
}

// the core of file represented in the sub editor
type buffer struct {
	// current location of the cursor
	cursor location

	// maintains a set of marks
	marks *list.List

	// holds the raw content of the buffer
	contents contents

	fileInfo os.FileInfo

	dirty bool

	// individual modes appended to the buffer
	modes *list.List
}

// the representation of a location of a point
// in the subeditor
type location struct {
	name string
	procs []func(v interface{}) error
}

// represents a mark made in the data
type mark struct {
	name string
	where location

	// fixed means the location is the character after the actual location
	// instead of living before i.e. if looking at b then not fixed: a|bc
	// fixed: ab|c
	fixed bool
}

// the data of the buffer, represented as a rectangle
type contents struct {
	data [][]byte
}

