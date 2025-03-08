package board

import (
	"container/list"
	"fmt"
	"sync"
)

func NextBoardStates(currentBoard *BoardState) []*BoardState {
	nextBoards := []*BoardState{}
	for i := 0; i < 4; i++ {
		newX, newY := currentBoard.ZeroX+movesX[i], currentBoard.ZeroY+movesY[i]
		if CanWeMove(newX, newY) {
			nextBoard := getBoard(currentBoard.Board)
			nextBoard[currentBoard.ZeroX][currentBoard.ZeroY], nextBoard[newX][newY] = nextBoard[newX][newY], nextBoard[currentBoard.ZeroX][currentBoard.ZeroY]

			nextBoardState := &BoardState{
				Depth:  currentBoard.Depth + 1,
				ZeroX:  newX,
				ZeroY:  newY,
				Board:  nextBoard,
				Parent: currentBoard,
			}

			nextBoards = append(nextBoards, nextBoardState)

		}
	}

	return nextBoards
}

func BidirectionalBFS(startBoard [][]int) ([]*BoardState, int) {
	startBoardState := &BoardState{
		Depth:  0,
		ZeroX:  FindZero(startBoard)[0],
		ZeroY:  FindZero(startBoard)[1],
		Board:  startBoard,
		Parent: nil,
	}

	finalBoardState := &BoardState{
		Depth:  0,
		ZeroX:  FindZero(goal)[0],
		ZeroY:  FindZero(goal)[1],
		Board:  goal,
		Parent: nil,
	}

	startQueue := list.New()
	finalQueue := list.New()

	var visitedStatesStart sync.Map
	var visitedStatesFinal sync.Map

	meetChannel := make(chan *BoardState, 1)
	var wg sync.WaitGroup

	wg.Add(2)

	go searchWorker(startQueue, &visitedStatesStart, &visitedStatesFinal, meetChannel, startBoardState, &wg)
	go searchWorker(finalQueue, &visitedStatesFinal, &visitedStatesStart, meetChannel, finalBoardState, &wg)

	go func() {
		wg.Wait()
		close(meetChannel)
	}()

	for match := range meetChannel {
		if match != nil {
			startState, foundStart := visitedStatesStart.Load(StringState(match.Board))
			finalState, foundFinal := visitedStatesFinal.Load(StringState(match.Board))
			if foundStart && foundFinal {
				return getPath(startState.(*BoardState), finalState.(*BoardState))
			}
		}
	}

	return nil, 0
}

func searchWorker(queue *list.List, visited *sync.Map, otherVisited *sync.Map, meetChannel chan<- *BoardState, start *BoardState, wg *sync.WaitGroup) {
	defer wg.Done()

	queue.PushBack(start)
	visited.Store(StringState(start.Board), start)

	for queue.Len() > 0 {
		head := queue.Remove(queue.Front()).(*BoardState)

		if _, found := otherVisited.Load(StringState(head.Board)); found {
			meetChannel <- head
			return
		}

		for _, child := range NextBoardStates(head) {
			childString := StringState(child.Board)
			if _, found := visited.Load(childString); !found {
				visited.Store(childString, child)
				queue.PushBack(child)
			}
		}
	}
}

func getPath(forward *BoardState, backwards *BoardState) ([]*BoardState, int) {
	path := []*BoardState{}

	for forward != nil {
		path = append(path, forward)
		forward = forward.Parent
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	meetIndex := len(path)
	backwards = backwards.Parent

	for backwards != nil {

		path = append(path, backwards)
		backwards = backwards.Parent
	}

	fmt.Println("Final path:")
	for i, step := range path {
		fmt.Printf("Step %d:\n", i)
		SeeCurrent(step.Board)
	}

	return path, meetIndex
}
