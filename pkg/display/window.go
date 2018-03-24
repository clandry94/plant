package display

import (
	"errors"
	"github.com/gdamore/tcell"
	"github.com/clandry94/plant/pkg/edit"
	"os"
	"github.com/clandry94/plant/pkg/edit/raw"
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
						  Background(tcell.ColorBlack).
						  Foreground(tcell.ColorWhite)

	screen.SetStyle(defaultStyle)
	// screen.EnableMouse()
	screen.Clear()


	return &Window{
		screen: screen,
		style: tcell.StyleDefault,
	}, nil

}

/*
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
*/

func (w *Window) Sync() {
	w.screen.Sync()
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
func (w *Window) Load(file *os.File) error {
	return errors.New("not implemented")
}

/*
	Performs an incremental reloading of the display.
	If run, this should make sure the window accurately
	represents the buffer.
 */
func (w *Window) Redisplay() {
	w.screen.Show()
}

/*
	Performs a full window reload. This makes sure
	that the screen is correct no matter what
 */
func (w *Window) Refresh(data *raw.Contents) {
	screen := w.screen
	p := data.Lines.Front()
	i := 0

	for p != nil {
		line := p.Value.([]rune)
		j := 0
		for _, r := range line {
			if j > 80 {
				break
			}
			screen.SetContent(j, i, r, []rune{}, tcell.StyleDefault)
			j++
		}
		p = p.Next()
		i++
	}

	screen.Sync()
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

func (w *Window) SetCursor(){

}

func (w *Window) SetRow() {

}

func (w *Window) SetCol() {

}

func (w *Window) ClearLine() {

}

func (w *Window) ClearScreen() {

}


func (w *Window) PutRune(r rune) {

}

func (w *Window) PutRunes(r []rune) {

}

func (w *Window) DeleteRunes(i int) {

}

func(w *Window) InsertLines(i int) {

}

func (w *Window) DeleteLines(i int) {

}


