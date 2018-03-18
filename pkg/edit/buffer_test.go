package edit

import (
	"github.com/clandry94/plant/edit/raw"
	"testing"
)

func newBuffer(name string) *Buffer {

	return &Buffer{
		name:     name,
		Cursor:   NilCursor(),
		contents: raw.NewContents(),
		file:     nil,
		dirty:    false,
		modes:    nil,
	}
}

func TestBuffer_Insert(t *testing.T) {
	buf := newBuffer("test")
	buf.contents.Lines.PushFront(raw.NewPiece())
	buf.Insert("Hello, world!")

	results := string(buf.contents.Lines.Front().Value.(*raw.Piece).Data())

	if results != "Hello, world!" {
		t.Errorf("expected %v != actual %v", "Hello, World!", results)
	}
}

func TestBuffer_Delete(t *testing.T) {
	buf := newBuffer("test")
	buf.contents.Lines.PushFront(raw.NewPiece())
	buf.Insert("Hello, world!")
	buf.Delete(7)

	results := string(buf.contents.Lines.Front().Value.(*raw.Piece).Data())

	if results != "world!" {
		t.Errorf("expected %v != actual %v", "world!", results)
	}
}

func TestBuffer_SetName(t *testing.T) {
	buf := newBuffer("test")
	buf.SetName("test2")

	if buf.Name() != "test2" {
		t.Errorf("buffer name should be %v and not %v", "test2", "test")
	}
}

func TestBuffer_Name(t *testing.T) {
	buf := newBuffer("test")
	if buf.Name() != "test" {
		t.Errorf("buffer name should be %v and not %v", "test", buf.Name())
	}
}

func TestBuffer_SetCursor(t *testing.T) {
	buf := newBuffer("test")
	cursor := NilCursor()
	cursor.SetCol(1)
	buf.SetCursor(cursor)

	if buf.Cursor.Col() != 1 {
		t.Errorf("col should be %v and not %v", 1, buf.Cursor.Col())
	}
}

func TestBuffer_SetCursorToCount(t *testing.T) {
	// todo: implement
}

func TestBuffer_CursorMoveForward(t *testing.T) {
	buf := newBuffer("test")
	buf.CursorMoveForward(1)
	if buf.Cursor.Col() != 1 {
		t.Errorf("col should be %v and not %v", 1, buf.Cursor.Col())
	}
}

func TestBuffer_CursorMoveBack(t *testing.T) {
	buf := newBuffer("test")
	buf.Cursor.SetCol(1)
	buf.CursorMoveBack(1)
	if buf.Cursor.Col() != 0 {
		t.Errorf("col should be %v and not %v", 0, buf.Cursor.Col())
	}
}

func TestBuffer_GetCursor(t *testing.T) {
	buf := newBuffer("test")
	buf.Cursor.SetCol(1)
	cursor := buf.GetCursor()
	if cursor.Col() != 1 {
		t.Errorf("gotten cursor was not the same as buffer cursor")
	}
}
