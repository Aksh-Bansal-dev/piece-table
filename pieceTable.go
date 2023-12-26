package main

import (
	"fmt"
)

const (
	ADD_BUFFER_TYPE      = true
	ORIGINAL_BUFFER_TYPE = false
)

type Piece struct {
	bufferType bool
	offset     int
	length     int
}

type PieceTable struct {
	ogBuffer  string // Original buffer
	addBuffer string // Add buffer
	table     []Piece
}

func newPieceTable(str string) PieceTable {
	rootPiece := Piece{bufferType: ORIGINAL_BUFFER_TYPE, offset: 0, length: len(str)}
	return PieceTable{ogBuffer: str, addBuffer: "", table: []Piece{rootPiece}}
}

func (pt *PieceTable) find(offset int) (int, int, error) {
	curOffset := 0
	for i, piece := range pt.table {
		curOffset += piece.length
		if curOffset > offset {
			return i, offset - (curOffset - piece.length), nil
		}
	}
	if curOffset == offset {
		return len(pt.table), 0, nil
	}
	return -1, -1, fmt.Errorf("Index out of bounds")
}

func (pt *PieceTable) toString() string {
	res := ""
	for _, piece := range pt.table {
		if piece.bufferType == ADD_BUFFER_TYPE {
			res += pt.addBuffer[piece.offset : piece.offset+piece.length]
		} else {
			res += pt.ogBuffer[piece.offset : piece.offset+piece.length]
		}
	}
	return res
}

func (pt *PieceTable) delete(offset, length int) error {
	for length > 0 {
		idx, leftLen, err := pt.find(offset)
		if err != nil {
			return err
		}
		if idx == len(pt.table) {
			return nil
		}
		leftPiece := Piece{
			bufferType: pt.table[idx].bufferType,
			offset:     pt.table[idx].offset,
			length:     leftLen}
		rightPiece := Piece{
			bufferType: pt.table[idx].bufferType,
			offset:     pt.table[idx].offset + leftLen + length,
			length:     pt.table[idx].length - leftLen - length}

		insertPieces := []Piece{}
		if leftPiece.length > 0 {
			insertPieces = append(insertPieces, leftPiece)
		}
		if rightPiece.length > 0 {
			insertPieces = append(insertPieces, rightPiece)
		}

		length -= (pt.table[idx].length - leftLen)
		pt.table = replaceAndInsert(pt.table, idx, insertPieces)

	}
	return nil
}

func (pt *PieceTable) add(str string, offset int) error {
	idx, leftLen, err := pt.find(offset)
	addPiece := Piece{bufferType: ADD_BUFFER_TYPE, length: len(str), offset: len(pt.addBuffer)}
	pt.addBuffer += str

	if err != nil {
		return err
	}

	if idx == len(pt.table) {
		pt.table = append(pt.table, addPiece)
		return nil
	}

	leftPiece := Piece{
		bufferType: pt.table[idx].bufferType,
		offset:     pt.table[idx].offset,
		length:     leftLen}
	rightPiece := Piece{
		bufferType: pt.table[idx].bufferType,
		offset:     pt.table[idx].offset + leftLen,
		length:     pt.table[idx].length - leftLen}

	insertPieces := []Piece{}
	if leftPiece.length > 0 {
		insertPieces = append(insertPieces, leftPiece)
	}
	insertPieces = append(insertPieces, addPiece)
	if rightPiece.length > 0 {
		insertPieces = append(insertPieces, rightPiece)
	}
	pt.table = replaceAndInsert(pt.table, idx, insertPieces)
	return nil
}

func replaceAndInsert[T any](a []T, index int, valueArr []T) []T {
	for i := len(valueArr) - 1; i >= 0; i-- {
		value := valueArr[i]
		if len(a) == index {
			return append(a, value)
		}
		a = append(a[:index+1], a[index:]...)
		a[index] = value
	}

	return append(a[:index+len(valueArr)], a[index+len(valueArr)+1:]...)
}
