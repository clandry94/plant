package edit

import (
	"container/ring"
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