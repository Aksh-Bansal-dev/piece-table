package main

type Piece struct {
	bufferType bool
	offset     int
	length     int
}

type PieceTable struct {
	og        string
	addBuffer string
	table     []Piece
}

func newPieceTable(str string) PieceTable {
	rootPiece := Piece{bufferType: false, offset: 0, length: len(str)}
	return PieceTable{og: str, addBuffer: "", table: []Piece{rootPiece}}
}
