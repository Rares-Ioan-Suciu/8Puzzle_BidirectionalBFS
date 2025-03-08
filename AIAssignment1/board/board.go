package board

const size = 3

type BoardState struct {
	Depth        int
	ZeroX, ZeroY int
	Board        [][]int
	Parent       *BoardState
}

var movesX = []int{0, 0, 1, -1}
var movesY = []int{1, -1, 0, 0}

var goal = [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
