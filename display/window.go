package display

import (
	"errors"
	"github.com/gdamore/tcell"
	"github.com/clandry94/plant/edit"
)

type Window struct {
	screen tcell.Screen
	style  tcell.Style
}

type Pane struct {

}

/*
	Grow the current pane by n lines
 */
func (p *Pane) Grow(amount int) error {
	return nil
}

func (p *Pane) TopLine() int {
	return 0
}

func (p *Pane) BottomLine() int {
	return 0
}

func (p *Pane) Top() *edit.Cursor {
	return nil
}

func (p *Pane) Bottom() *edit.Cursor{
	return nil
}

func NewWindow() (*Window, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}

	err = screen.Init()
	if err != nil {
		return nil, err
	}

	defaultStyle := tcell.StyleDefault.
						  Background(tcell.ColorWhite).
						  Foreground(tcell.ColorBlack)

	screen.SetStyle(defaultStyle)
	screen.EnableMouse()
	screen.Clear()


	return &Window{
		screen: screen,
		style: tcell.StyleDefault,
	}, nil
}

func (w *Window) Poll() error {
	for {
		event := w.screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				panic("temporary")
			case tcell.KeyCtrlL:
				w.screen.Sync()
			}
		case *tcell.EventResize:
			w.screen.Sync()
		}
	}

	return nil
}

/*
	creates a new window depending on the config provided
 */
func (w *Window) Create() error {
	// TODO: accept options for things like splits, ordering, new buffer, etc
	return nil
}

/*
	destroy the pane supplied
 */
func (w *Window) Destroy(pane *Pane) error {
	return nil
}

/*
	Exit the window and do cleanup
 */
func (w *Window) Close() error {
	return nil
}

/*
	Saves a context (all state info, current open buffers, etc)
 */
func (w *Window) Save() error {
	return errors.New("not implemented")
}

/*
	Loads a saved context
 */
func (w *Window) Load() error {
	return errors.New("not implemented")
}

/*
	Performs an incremental reloading of the display.
	If run, this should make sure the window accurately
	represents the buffer.
 */
func (w *Window) Redisplay() {
}

/*
	Performs a full window reload. This makes sure
	that the screen is correct no matter what
 */
func (w *Window) Refresh() {
}

/*
	Gets current row that the cursor is on in
	the window. This might be different from
	the cursor point in the editor due to line wrap
 */
func (w *Window) CursorRow() int {
	return 0
}

/*
	Same as CursorRow() but for col
 */
func (w *Window) CursorCol() int {
	return 0
}
