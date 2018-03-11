package edit

import (
	"container/list"
	"errors"
)

// represents a mark made in the data
type Mark struct {
	name  string
	where Cursor

	// fixed means the location is the character after the actual location
	// instead of living before i.e. if looking at b then not fixed: a|bc
	// fixed: ab|c
	fixed bool
}

func (m *Mark) SetCursor(location Cursor) {
	m.where = location
}

func (m *Mark) GetCursor() Cursor {
	return m.where
}

type Marks struct {
	Marks *list.List
}

/*
	Creates a new mark and places it at the front of the mark list

 	TODO: keep the marks sorted in the order that they are in the buffer
 	TODO: check if mark with the same name exists in the list
*/
func (m *Marks) Create(name string, location Cursor, fixed bool) error {
	if name == "" {
		return errors.New("mark must have a name")
	}

	mark := Mark{
		name:  name,
		where: location,
		fixed: fixed,
	}

	m.Marks.PushFront(mark)

	return nil
}

/*
	Deletes a mark by its name
	TODO: Traverse the mark list and delete the mark. Error a const if doesn't exist
*/
func (m *Marks) Delete(name string) error {

	return nil
}

/*
	Receives a Cursor and sets that mark to the Cursor's location
	TODO: traverse mark list to correct mark, then change its location
*/
func (m *Marks) MarkToPoint(name string, location Cursor) error {
	return nil
}
