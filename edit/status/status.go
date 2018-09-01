package status

/*
	Provides status values for the editor. Done this way
	to abstract away modification of the editor.
 */

type Status struct {
	Lines			int
	CurrentLine		int
	Cols			int
	CurrentCol		int
}




