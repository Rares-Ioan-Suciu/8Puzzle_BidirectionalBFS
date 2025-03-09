# 8Puzzle_BidirectionalBFS

## Description

**8Puzzle_BidirectionalBFS** is a Go-based application that solves the classic 8-puzzle problem using **bidirectional breadth-first search (BFS)**.

## Features

- **Efficient 8-Puzzle Solver:** Uses bidirectional BFS to quickly find the shortest solution path.
- **Graphical User Interface:** Built with Fyne, allowing users to visualize the puzzle-solving process.
- **Custom & Random Puzzles:** Users can input their own puzzle or generate a random solvable configuration.
- **Step-by-Step Visualization:** Animates each step of the solution.
- **Pause & Resume Feature:** Allows users to control the game execution.

## Installation

### Prerequisites

- Go installed (version 1.16 or later)
- Fyne library (`fyne.io/fyne/v2`)

### Clone the Repository

```sh
git clone https://github.com/Rares-Ioan-Suciu/8Puzzle_BidirectionalBFS.git
cd 8Puzzle_BidirectionalBFS
```

### Install Dependencies

```sh
go mod tidy
```

## Running the Application

```sh
go run gui.go
```

## Usage

1. Choose to **start a new game** (random or custom board).
2. Click **Start** to begin solving.
3. Pause, reset, or step through the solution manually.

## How It Works

### Bidirectional BFS Algorithm

The program:
- Starts two simultaneous BFS searches using go routines (threads): one from the initial board and one from the goal state.
- If the two searches meet, the solution path is reconstructed.
- The search reduces time complexity compared to traditional BFS.

### Board Representation

The puzzle board is represented as a 3x3 matrix, where `0` represents the empty tile.

Example:

```
1 2 3
4 5 6
7 8 0
```

###
[Youtube Demo](https://youtu.be/4SW85icXgYs)
