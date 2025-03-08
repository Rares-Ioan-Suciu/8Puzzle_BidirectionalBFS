package board

import (
	"fmt"
	"math/rand"
	"unicode"
)

func CanWeMove(x int, y int) bool {
	return 0 <= x && x < size && 0 <= y && y < size
}

func SeeCurrent(board [][]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(board[i][j])
			if j == size-1 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func getBoard(board [][]int) [][]int {
	newBoard := make([][]int, size)
	for i := range board {
		newBoard[i] = make([]int, size)
		copy(newBoard[i], board[i])
	}
	return newBoard
}

func RandomInitial() [][]int {
	ok := true
	var permutation []int

	for ok {
		permutation = rand.Perm(9)
		inversions := 0
		for i := 0; i < 8; i++ {
			for j := i + 1; j < 9; j++ {
				if permutation[i] > permutation[j] && permutation[i] != 0 && permutation[j] != 0 {
					inversions++
				}
			}
		}
		ok = !(inversions%2 == 0)
	}

	result := make([][]int, 3)
	for i := 0; i < 3; i++ {
		result[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			result[i][j] = permutation[i*3+j]
		}
	}
	return result
}

func FindZero(board [][]int) [2]int {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

func StrBoard(text string) [][]int {
	if len(text) != 9 {
		return nil
	}

	for i := 0; i < len(text); i++ {
		if !unicode.IsDigit(rune(text[i])) {
			return nil
		}

	}

	board := make([][]int, size)

	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
		for j := 0; j < size; j++ {

			board[i][j] = int(text[i*3+j] - '0')
		}
	}

	return board
}

func CheckBoard(board [][]int) bool {
	unidim := make([]int, 9)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			unidim[i*3+j] = board[i][j]
		}
	}
	inversions := 0
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if unidim[i] > unidim[j] && unidim[i] != 0 && unidim[j] != 0 {
				inversions++
			}
		}
	}

	return inversions%2 == 0
}

func StringState(board [][]int) string {
	result := ""
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result += fmt.Sprintf("%d,", board[i][j])
		}
	}
	return result
}
