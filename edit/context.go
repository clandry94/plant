package edit

import (
	"container/list"
	"container/ring"
	"os"
)

// the core of the sub editor. Only one context exists
// in the realm of a plant editor instance, but there could be
// more than one ecosystem in the future
type Context struct {
	// a circular chain with pointers to buffers
	// each buffer represents an open file (or new file?)
	buffers *ring.Ring

	// the current buffer in focus
	currentBuffer *Buffer
}

func NewContext() (Context, error) {
	return Context{}, nil
}

// save context state/buffer states to a file
func (e Context) Save(filename string) error {
	return nil
}

// load a context state from a file
func (e Context) Load(filename string) error {
	return nil
}

// create a new buffer with no file info
func (e Context) BufferCreate(bufferName string) error {
	return nil
}

// clear a buffer's contents (doesn't write to disk)
func (e Context) BufferClear(bufferName string) error {
	return nil
}

// delete a buffer from the context and lose all progress
func (e Context) BufferDelete(bufferName string) error {
	return nil
}

// set the current buffer to the buffer name
func (e Context) BufferSet(bufferName string) error {
	return nil
}

// switch to the next buffer and return the buffer name
func (e Context) BufferSetNext() string {
	return "the buffer name"
}

// switch to the previous buffer in the ring
func (e Context) BufferSetPrev() string {
	return "the prev buffer"
}

// the core of file represented in the sub editor
type Buffer struct {
	name string

	// current location of the cursor
	cursor Location

	// maintains a set of marks
	marks *list.List

	// holds the raw content of the buffer
	contents contents

	fileInfo os.FileInfo

	dirty bool

	// individual modes appended to the buffer
	modes *list.List
}

func (b *Buffer) SetName(name string) error {
	return nil
}

func (b *Buffer) GetName() string {
	return b.name
}

func (b *Buffer) CursorSet(loc Location) error {
	return nil
}

func (b *Buffer) CursorMoveForward(i int) error {
	return nil
}

func (b *Buffer) CursorMoveBack(i int) error {
	return nil
}

func (b *Buffer) GetCursor() Location {
	return b.cursor
}

// gets the location at the start of the buffer (might not be 0)
func (b *Buffer) BufferStart() Location {
	return Location{}
}

// gets the location at the end of the buffer (might not be the end)
func (b *Buffer) BufferEnd() Location {
	return Location{}
}

// the representation of a location of a point
// in the subeditor
type Location struct {
	name  string
	procs []func(v interface{}) error
}

// represents a mark made in the data
type mark struct {
	name  string
	where Location

	// fixed means the location is the character after the actual location
	// instead of living before i.e. if looking at b then not fixed: a|bc
	// fixed: ab|c
	fixed bool
}

// the data of the buffer, represented as a rectangle
type contents struct {
	data [][]byte
}
