package cmd

import (
	"github.com/gdamore/tcell"
	"github.com/clandry94/plant/pkg/display"
)

/*
	Holds the keybindings for basic cursor commands
 */


func Handle(ev tcell.Event, window *display.Window) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			// need to implement closing
			panic("exited!")
		case tcell.KeyCtrlH: // move cursor left
			screen := window.Screen()
			window.SetCursor(window.CursorCol()-1, window.CursorRow())
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		case tcell.KeyCtrlL: // move cursor right
			screen := window.Screen()
			window.SetCursor(window.CursorCol()+1, window.CursorRow())
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		case tcell.KeyCtrlJ: // move cursor down
			screen := window.Screen()
			window.SetCursor(window.CursorCol(), window.CursorRow()+1)
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		case tcell.KeyCtrlK: // move cursor up
			screen := window.Screen()
			window.SetCursor(window.CursorCol(), window.CursorRow()-1)
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		}
	}
}
