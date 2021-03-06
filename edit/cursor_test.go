package edit_test

import (
	"container/list"
	"testing"

	"github.com/clandry94/plant/edit"
)

func TestCursor_Line(t *testing.T) {
	cursor := edit.NilCursor()
	if cursor.Line() != 0 {
		t.Errorf("expected %v got %v", 0, cursor.Line())
	}
}

func TestCursor_SetLine(t *testing.T) {
	cursor := edit.NilCursor()
	cursor.SetLine(1)

	if cursor.Line() != 1 {
		t.Errorf("expected %v got %v", 1, cursor.Line())
	}
}

func TestCursor_Col(t *testing.T) {
	cursor := edit.NilCursor()
	if cursor.Col() != 0 {
		t.Errorf("expected %v got %v", 0, cursor.Col())
	}
}

func TestCursor_SetCol(t *testing.T) {
	cursor := edit.NilCursor()

	cursor.SetCol(1)
	if cursor.Col() != 1 {
		t.Errorf("expected %v got %v", 1, cursor.Col())
	}
}

func TestCursor_SetToMark(t *testing.T) {
	markList := edit.Marks{
		Marks: list.New(),
	}

	mark, err := markList.Create("dummy", edit.NilCursor(), false)
	if err != nil {
		t.Error(err)
		return
	}

	cursor := edit.NilCursor()
	cursor.SetCol(5)
	cursor.SetLine(5)

	err = cursor.SetToMark(mark)
	if err != nil {
		t.Error(err)
		return
	}

	if cursor.Line() != 0 || cursor.Col() != 0 {
		t.Errorf("cursor not equal to mark. line: %d, col: %d", cursor.Line(), cursor.Col())
	}

}

func TestCursor_Swap(t *testing.T) {
	// TODO: needs implemented
}

func TestCursor_AtMark(t *testing.T) {
	markList := edit.Marks{
		Marks: list.New(),
	}

	mark, err := markList.Create("dummy", edit.NilCursor(), false)
	if err != nil {
		t.Error(err)
		return
	}

	cursor := edit.NilCursor()
	cursor.SetCol(5)
	cursor.SetLine(5)

	atMark := cursor.AtMark(mark)
	if atMark == true {
		t.Error("cursor at mark when it should not be")
	}

	cursor.SetCol(0)
	cursor.SetLine(0)
	atMark = cursor.AtMark(mark)
	if atMark == false {
		t.Error("cursor not at mark when it should be")
	}
}

func TestCursor_BeforeMark(t *testing.T) {
	markList := edit.Marks{
		Marks: list.New(),
	}

	mark, err := markList.Create("dummy", edit.NilCursor(), false)
	if err != nil {
		t.Error(err)
		return
	}

	cursor := edit.NilCursor()
	cursor.SetCol(5)
	cursor.SetLine(5)

	beforeMark := cursor.BeforeMark(mark)
	if beforeMark == true {
		t.Error("cursor should not be before mark")
	}

	cursor.SetCol(-5)
	cursor.SetLine(-5)
	beforeMark = cursor.BeforeMark(mark)
	if beforeMark == false {
		t.Error("cursor should be before mark")
	}
}

func TestCursor_AfterMark(t *testing.T) {
	markList := edit.Marks{
		Marks: list.New(),
	}

	mark, err := markList.Create("dummy", edit.NilCursor(), false)
	if err != nil {
		t.Error(err)
		return
	}

	cursor := edit.NilCursor()
	cursor.SetCol(5)
	cursor.SetLine(5)

	afterMark := cursor.AfterMark(mark)
	if afterMark == false {
		t.Error("cursor should be after mark")
	}

	cursor.SetCol(-5)
	cursor.SetLine(-5)
	afterMark = cursor.AfterMark(mark)
	if afterMark == true {
		t.Error("cursor should not be after mark")
	}
}
