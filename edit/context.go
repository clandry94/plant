package edit

import (
	"container/list"
	"errors"
)

// the core of the sub editor. Only one context exists
// in the realm of a plant editor instance, but there could be
// more than one ecosystem in the future
type Context struct {
	// a circular chain with pointers to buffers
	// each buffer represents an open file (or new file?)
	buffers *list.List

	// the current buffer in focus
	currentBuffer *bufferElement
}

type bufferElement list.Element

func (e bufferElement) Buffer() *Buffer {
	return e.Value.(*Buffer)
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
func (e Context) BufferAdd(bufferName string) error {
	buffer := &Buffer{
		name:     bufferName,
		Cursor:   NilCursor(),
		contents: nil,
		file:     nil,
		dirty:    true,
		modes:    nil,
	}

	e.buffers.PushFront(buffer)

	return nil
}

// clear a buffer's contents (doesn't write to disk)
func (e Context) BufferClear(bufferName string) error {

	if e.currentBuffer.Buffer().name == bufferName {
		clearBuffer(e.currentBuffer.Value.(*Buffer))
		return nil
	}

	bufferElement, err := e.findBufferElement(bufferName)
	if err != nil {
		return err
	}

	// clear the buffer
	clearBuffer(bufferElement.Buffer())

	return nil
}

func clearBuffer(buffer *Buffer) {
	// TODO: clear the buffer
}

func (e Context) findBufferElement(bufferName string) (*bufferElement, error) {
	for b := e.buffers.Front(); b != nil; b = b.Next() {
		if b.Value.(*Buffer).name == bufferName {
			return (*bufferElement)(b), nil
		}
	}

	return nil, errors.New("buffer not found")
}

// delete a buffer from the context and lose all progress
func (e Context) BufferDelete(bufferName string) error {
	return nil
}

// set the current buffer to the buffer name
func (e Context) BufferSet(bufferName string) error {
	bufferElement, err := e.findBufferElement(bufferName)
	if err != nil {
		return err
	}

	// TODO: investigate if a race condition can happen here
	e.currentBuffer = bufferElement

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
