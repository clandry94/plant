package raw_test

import (
	"testing"
	"github.com/clandry94/plant/edit/raw"
	"fmt"
)

type test struct {
	expected []rune
	start int
	runes []rune
}

func Test_PieceInsert(t *testing.T) {
	// expected : start, insert
	tests := []test{
		{[]rune("hello"), 0, []rune("hello")},
		{[]rune("hello world"), 5, []rune(" world")},
		{[]rune("hello, world"), 5, []rune(",")},
		{[]rune("Why hello, world"), 0, []rune("Why ")},
	}

	piece := &raw.Piece{}

	for _, test := range tests {
		piece.Insert(test.start, test.runes)

		fmt.Println(string(piece.Data()))
		if string(piece.Data()) != string(test.expected) {
			t.Errorf("actual %v != expected %v", string(piece.Data()), string(test.expected))
		}
	}
}

func Test_PieceDelete(t *testing.T) {

}