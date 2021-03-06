package plant

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/clandry94/plant/edit"
	"github.com/clandry94/plant/errors"
	"github.com/gdamore/tcell"
	"github.com/sirupsen/logrus"
	"os"
	"os/user"
)

var log = logrus.New()
var config = &Config{}

func init() {
	file, err := os.OpenFile("cmd.log", os.O_CREATE|os.O_WRONLY, 0777)
	if err == nil {
		log.Out = file
	} else {
		log.Info("failed to log to a file, using stderr")
	}

	initConfig()
	overrideKeys()
}

type Config struct {
	CBackward  string
	CForward   string
	CUp        string
	CDown      string
	CJumpToEnd string
	CDelete    string
}

func overrideKeys() {

}

func initConfig() {
	usr, err := user.Current()
	if err != nil {
		log.Error(err)
	}

	path := fmt.Sprintf("%s/%s", usr.HomeDir, ".plantrc")

	_, err = toml.DecodeFile(path, &config)
	if err != nil {
		log.Error(err)
	}
}

/*
	Holds the keybindings for basic cursor commands
*/
func Handle(ev tcell.Event, buffer *edit.Buffer) error {
	logger := log.WithField("module", "cmd")

	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch Keys[ev.Key()] {
		case Exit:
			// need to implement closing
			return errors.Exit()

		case CursorLeft: // move cursor left
			logger.Info("moving cursor move left")
			// window.SetCursor(window.CursorCol()-1, window.CursorRow())
			buffer.CursorMoveBack(1)

		case CursorRight: // move cursor right
			logger.Info("cursor move right")
			// window.SetCursor(window.CursorCol()+1, window.CursorRow())
			buffer.CursorMoveForward(1)

		case CursorDown: // move cursor down
			logger.Info("cursor move down")
			// window.SetCursor(window.CursorCol(), window.CursorRow()+1)
			buffer.CursorMoveDown(1)

		case CursorUp: // move cursor up
			logger.Info("cursor move up")
			// window.SetCursor(window.CursorCol(), window.CursorRow()-1)
			buffer.CursorMoveUp(1)

		case JumpCursorEnd: // set cursor end of line
			logger.Info("sending cursor to end of line!")
			buffer.SetCursorEndOfLine()

		case Backspace:
			logger.Info("backspace!")
			// window.DeleteRunes(1)
			buffer.Delete(1)

		default:
			// logger.Infof("inserting rune at x: %v y: %v | %v",
			//	window.CursorCol(), window.CursorRow(), string(ev.Rune()))
			// window.PutRune(ev.Rune())
			buffer.Insert(string(ev.Rune()))
		}
	}

	return nil
}
