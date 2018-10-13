package plant

import "github.com/gdamore/tcell"

const (
	Exit = iota
	CursorLeft
	CursorRight
	CursorDown
	CursorUp
	JumpCursorEnd
	Backspace
)

var Keys = map[tcell.Key]int{
	tcell.KeyEscape:     Exit,
	tcell.KeyCtrlH:      CursorLeft,
	tcell.KeyCtrlL:      CursorRight,
	tcell.KeyCtrlK:      CursorUp,
	tcell.KeyCtrlJ:      CursorDown,
	tcell.KeyCtrlE:      JumpCursorEnd,
	tcell.KeyBackspace2: Backspace,
}
