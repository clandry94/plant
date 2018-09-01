package edit

import (
	"container/list"
	"errors"
	"github.com/clandry94/plant/edit/raw"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
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

func NewContext() (Context, error) {

	file, err := os.OpenFile("editor.log", os.O_CREATE|os.O_WRONLY, 0777)
	if err == nil {
		log.Out = file
	} else {
		log.Info("failed to log to a file, using stderr")
	}

	log.WithFields(logrus.Fields{
		"module" : "context",
	}).Info("created context")

	return Context{
		buffers: list.New(),
	}, nil
}

func (e Context) CurrentBuffer() *Buffer {
	log.Debug("retrieving current buffer, it is a:")
	log.Debug(e.buffers.Front().Value.(*Buffer))
	return e.buffers.Front().Value.(*Buffer)
}

// save context/buffer states to a file
func (e Context) Save(filename string) error {
	return nil
}

// load a context state from a file
func (e Context) Load(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	err = e.NewBuffer(Name(file.Name()), File(file))
	if err != nil {
		log.Panic("could not load buffer!")
		return err
	}

	return nil
}

/*
	Functional options for creating buffers
 */
type BufferOptions struct {
	Name string
	File *os.File
}

type BufferOption func(*BufferOptions)

func Name(name string) BufferOption {
	return func(args *BufferOptions) {
		args.Name = name
	}
}

func File(file *os.File) BufferOption {
	return func(args *BufferOptions) {
		args.File = file
	}
}

// create a new buffer with no file info and push to the front
// of the list
func (e *Context) NewBuffer(args ...BufferOption) error {
	options := &BufferOptions{
		Name: "test", // will have issues if more than two with the default are made
		File: nil,
	}

	for _, arg := range args {
		arg(options)
	}

	logger := log.WithFields(logrus.Fields{
		"filename" : options.File.Name(),
		"bufname" : options.Name,
	})

	logger.Info("creating buffer")

	buffer := &Buffer{
		name:     options.Name,
		Cursor:   NilCursor(),
		contents: &raw.Contents{},
		file:     options.File,
		dirty:    true,
		modes:    nil,
	}

	contents, err := raw.NewContentsFromFile(buffer.file)
	if err != nil {
		logger.Errorf("unable to create buffer: %v", err)
		return err
	}

	buffer.contents = contents

	logger.Debug("adding buffer to front of list")
	e.buffers.PushFront(buffer)
	logger.Debugf("new buffer list size %v", e.buffers.Len())
	e.currentBuffer = e.buffers.Front()

	return nil
}

// clear a buffer's contents (doesn't write to disk)
func (e Context) BufferClear(bufferName string) error {

	if e.currentBuffer.Value.(*Buffer).name == bufferName {
		clearBuffer(e.currentBuffer.Value.(*Buffer))
		return nil
	}

	bufferElement, err := e.findBufferElement(bufferName)
	if err != nil {
		return err
	}

	// clear the buffer
	clearBuffer(bufferElement.Value.(*Buffer))

	return nil
}

func clearBuffer(buffer *Buffer) {
	// TODO: clear the buffer
}

func (e Context) findBufferElement(bufferName string) (*list.Element, error) {
	for b := e.buffers.Front(); b != nil; b = b.Next() {
		if b.Value.(*Buffer).Name() == bufferName {
			return b, nil
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
func (e *Context) BufferSetNext() string {
	e.currentBuffer = e.currentBuffer.Next()
	return e.currentBuffer.Value.(*Buffer).Name()
}

// switch to the previous buffer in the ring
func (e *Context) BufferSetPrev() string {
	e.currentBuffer = e.currentBuffer.Prev()
	return e.currentBuffer.Value.(*Buffer).Name()
}
