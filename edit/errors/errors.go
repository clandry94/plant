package errors

import (
	"fmt"
)

type IndexOutOfRangeError struct {
	PieceLength int
	IndexPos    int
}

func (e IndexOutOfRangeError) Error() string {
	return fmt.Sprintf("index out of range. length: %v, indexPos: %v", e.PieceLength, e.IndexPos)
}
