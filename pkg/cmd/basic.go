package cmd

import (
	"github.com/gdamore/tcell"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/clandry94/plant/pkg/edit"
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

func Handle(ev tcell.Event, buffer *edit.Buffer) {
	logger := log.WithField("module", "cmd")

	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			// need to implement closing
			panic("exited!")
		case tcell.KeyCtrlH: // move cursor left
			logger.Info("moving cursor move left")
			// window.SetCursor(window.CursorCol()-1, window.CursorRow())
			buffer.CursorMoveBack(1)
		case tcell.KeyCtrlL: // move cursor right
			logger.Info("cursor move right")
			// window.SetCursor(window.CursorCol()+1, window.CursorRow())
			buffer.CursorMoveForward(1)
		case tcell.KeyCtrlJ: // move cursor down
			logger.Info("cursor move down")
			// window.SetCursor(window.CursorCol(), window.CursorRow()+1)
			buffer.CursorMoveDown(1)
		case tcell.KeyCtrlK: // move cursor up
			logger.Info("cursor move up")
			// window.SetCursor(window.CursorCol(), window.CursorRow()-1)
			buffer.CursorMoveUp(1)
		case tcell.KeyBackspace2:
			logger.Info("backspace!")
			// window.DeleteRunes(1)
			buffer.Delete(1)
		case tcell.KeyRune:
			// logger.Infof("inserting rune at x: %v y: %v | %v",
			//	window.CursorCol(), window.CursorRow(), string(ev.Rune()))
			// window.PutRune(ev.Rune())
			buffer.Insert(string(ev.Rune()))
		}
	}
}
