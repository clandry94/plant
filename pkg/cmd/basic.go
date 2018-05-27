package cmd

import (
	"github.com/gdamore/tcell"
	"github.com/clandry94/plant/pkg/display"
	"os"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	file, err := os.OpenFile("cmd.log", os.O_CREATE|os.O_WRONLY, 0777)
	if err == nil {
		log.Out = file
	} else {
		log.Info("failed to log to a file, using stderr")
	}
}

/*
	Holds the keybindings for basic cursor commands
 */

func Handle(ev tcell.Event, window *display.Window) {
	logger := log.WithField("module", "cmd")

	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			// need to implement closing
			panic("exited!")
		case tcell.KeyCtrlH: // move cursor left
			logger.Info("moving cursor move left")
			screen := window.Screen()
			window.SetCursor(window.CursorCol()-1, window.CursorRow())
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		case tcell.KeyCtrlL: // move cursor right
			logger.Info("cursor move right")
			screen := window.Screen()
			window.SetCursor(window.CursorCol()+1, window.CursorRow())
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		case tcell.KeyCtrlJ: // move cursor down
			logger.Info("cursor move down")
			screen := window.Screen()
			window.SetCursor(window.CursorCol(), window.CursorRow()+1)
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		case tcell.KeyCtrlK: // move cursor up
			logger.Info("cursor move up")
			screen := window.Screen()
			window.SetCursor(window.CursorCol(), window.CursorRow()-1)
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		case tcell.KeyRune:
			logger.Infof("inserting rune at x: %v y: %v | %v",
				window.CursorCol(), window.CursorRow(), string(ev.Rune()))
			screen := window.Screen()
			screen.SetContent(window.CursorCol(), window.CursorRow(), ev.Rune(), []rune{}, tcell.StyleDefault)
			window.SetCursor(window.CursorCol()+1, window.CursorRow())
			screen.ShowCursor(window.CursorCol(), window.CursorRow())
		}
	}
}
