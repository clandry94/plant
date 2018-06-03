package raw

import (
	"container/list"
	"os"
	"bufio"
)

const (
	lineLength = 8 // 32 bytes of lines to start TODO: are pointers 32 bit?
	dataSize   = 4 // 16 bytes per line default length
)

type Contents struct {
	Lines *list.List
}

func NewContents() *Contents {
	return &Contents{
		Lines: list.New(),
	}
}

func NewContentsFromFile(file *os.File) (*Contents, error) {
	contents := &Contents {
		Lines: list.New(),
	}

	scanner := bufio.NewScanner(file)

	// currently 1 piece per line
	for scanner.Scan() {
		runes := []rune(scanner.Text())
		runes = append(runes, []rune("\n")...)
		piece := NewPiece()
		piece.Insert(0, runes)

		contents.Lines.PushBack(piece)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
}

	return contents, nil

}

type Piece struct {
	length int
	used   int
	data   []rune
	// marks
	// version
}

func NewPiece() *Piece{
	return &Piece{
		length: 0,
		used: 0,
		data: make([]rune, 0),
	}
}

// delete length characters from start
func (p *Piece) Delete(start int, length int) {
	p.data = append(p.data[:start], p.data[start+length:]...)
}

func (p *Piece) Insert(start int, runes []rune) {
	p.data = append(p.data[:start], append(runes, p.data[start:]...)...)
}

func (p *Piece) Data() []rune {
	return p.data
}
