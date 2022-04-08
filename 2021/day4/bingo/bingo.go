package bingo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Cell struct {
	Row    int
	Column int
	Value  int
	Marked bool
}

type Board struct {
	Cells []Cell
}

func (b Board) CalculateScore(aValue int) int {
	score := 0
	for _, cell := range b.Cells {
		if !cell.Marked {
			score += cell.Value
		}
	}
	return score * aValue
}

func (b Board) Play(aValue int) int {
	for idx, cell := range b.Cells {
		if cell.Value == aValue {
			newCell := Cell{cell.Row, cell.Column, cell.Value, true}
			b.Cells[idx] = newCell
			// cell.Marked = true
			if b.Wins() {
				v := b.CalculateScore(aValue)
				return v
			}
		}
	}
	return 0
}

func (b Board) GetCell(row int, column int) (*Cell, error) {
	for _, cell := range b.Cells {
		if cell.Row == row && cell.Column == column {
			return &cell, nil
		}
	}
	return nil, errors.New("not found")
}

func (b Board) Wins() bool {
	rowCounts := []int{0, 0, 0, 0, 0}
	columnCounts := []int{0, 0, 0, 0, 0}
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			cell, _ := b.GetCell(row, col)
			if cell.Marked {
				rowCounts[row] += 1
				columnCounts[col] += 1
				if rowCounts[row] == 5 || columnCounts[col] == 5 {
					return true
				}
			}
		}
	}
	return false
}

func MakeBoard(inputRows []string) (*Board, error) {
	newBoard := Board{}
	for idx, val := range inputRows {
		clean := strings.TrimSpace(val)
		data := strings.ReplaceAll(clean, "  ", " ")
		values := strings.Split(data, " ")
		for vidx, vval := range values {
			cellValue, _ := strconv.Atoi(vval)
			newBoard.Cells = append(newBoard.Cells, Cell{
				Row:    idx,
				Column: vidx,
				Value:  cellValue,
				Marked: false,
			})
		}
	}
	if len(newBoard.Cells) != 25 {
		fmt.Println(inputRows)
		return nil, errors.New("invalid board")
	}
	return &newBoard, nil
}

func GetResult(input []string) int {
	pinput := strings.Split(input[0], ",")
	plays := []int{}
	for _, value := range pinput {
		v, _ := strconv.Atoi(value)
		plays = append(plays, v)
	}
	boards := []*Board{}
	buf := []string{}
	for _, rowString := range input[1:] {
		if len(rowString) > 0 {
			buf = append(buf, rowString)
		}
		if len(buf) == 5 {
			newBoard, err := MakeBoard(buf)
			if err == nil {
				boards = append(boards, newBoard)
				buf = []string{}
			}
		}
	}
	for _, play := range plays {
		for _, board := range boards {
			result := board.Play(play)
			if result > 0 {
				return result
			}

		}
	}
	return 0
}
func GetLastWinner(input []string) int {
	pinput := strings.Split(input[0], ",")
	plays := []int{}
	for _, value := range pinput {
		v, _ := strconv.Atoi(value)
		plays = append(plays, v)
	}
	boards := []*Board{}
	buf := []string{}
	for _, rowString := range input[1:] {
		if len(rowString) > 0 {
			buf = append(buf, rowString)
		}
		if len(buf) == 5 {
			newBoard, err := MakeBoard(buf)
			if err == nil {
				boards = append(boards, newBoard)
				buf = []string{}
			}
		}
	}
	initialBoardCount := len(boards)
	winningBoards := []*Board{}
	for _, play := range plays {
		for _, board := range boards {
			if !BoardIsInList(winningBoards, board) {
				result := board.Play(play)
				if result > 0 {
					winningBoards = append(winningBoards, board)
					if len(winningBoards) == initialBoardCount {
						return result
					}
				}
			}

		}
	}
	return 0
}

func BoardIsInList(aList []*Board, board *Board) bool {
	for _, val := range aList {
		if val == board {
			return true
		}
	}
	return false
}
