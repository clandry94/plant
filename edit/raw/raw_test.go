package raw

import (
	"fmt"
	"testing"
)

const (
	testString = "thisisastringthatthebenchmarktestwillactupon"
)

type insertTest struct {
	expected []rune
	start    int
	runes    []rune
}

type deleteTest struct {
	startWord []rune
	expected  []rune
	start     int
	length    int
}

func Test_PieceInsert(t *testing.T) {
	// expected : start, insert
	tests := []insertTest{
		{[]rune("hello"), 0, []rune("hello")},
		{[]rune("hello world"), 5, []rune(" world")},
		{[]rune("hello, world"), 5, []rune(",")},
		{[]rune("Why hello, world"), 0, []rune("Why ")},
	}

	piece := &Piece{}

	for _, test := range tests {
		piece.Insert(test.start, test.runes)

		fmt.Println(string(piece.Data()))
		if string(piece.Data()) != string(test.expected) {
			t.Errorf("actual %v != expected %v", string(piece.Data()), string(test.expected))
		}
	}
}

func Benchmark_PieceInsertDifferentString10000(b *testing.B) {
	piece := &Piece{}

	testRune := []rune(testString)

	for n := 0; n < b.N; n++ {
		piece.data = testRune

		piece.Insert(4, testRune)
	}
}

func Benchmark_PieceInsertSameString100(b *testing.B) {
	piece := &Piece{}

	testRune := []rune(testString)
	piece.data = testRune

	for n := 0; n < b.N; n++ {
		piece.Insert(4, testRune)
	}
}


func Test_PieceDelete(t *testing.T) {
	// expected : start, length
	tests := []deleteTest{
		{[]rune("bonjour"), []rune(""), 0, 7},
		{[]rune("bonjour"), []rune("jour"), 0, 3},
		{[]rune("bonjour"), []rune("bonour"), 3, 1},
		{[]rune("bonjour"), []rune("b"), 1, 6},
		{[]rune("bonjour"), []rune("bonjou"), 6, 1},
		{[]rune("bonjour"), []rune("bonjour"), 0, 0},
	}

	piece := &Piece{}

	for _, test := range tests {
		piece.data = test.startWord

		piece.Delete(test.start, test.length)
		fmt.Println(string(piece.Data()))
		if string(piece.Data()) != string(test.expected) {
			t.Errorf("actual %v != expected %v", string(piece.Data()), string(test.expected))
		}
	}
}

func Benchmark_PieceDelete10000(b *testing.B) {
	piece := &Piece{}

	testRune := []rune(testString)

	for n := 0; n < b.N; n++ {
		piece.data = testRune

		piece.Delete(4, 8)
	}
}