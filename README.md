# 8Puzzle_BidirectionalBFS

## Description

**8Puzzle_BidirectionalBFS** is a Go-based application that solves the classic 8-puzzle problem using **bidirectional breadth-first search (BFS)**. The application provides both a command-line solver and a graphical user interface (GUI) using the **Fyne** framework.

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
git clone https://github.com/yourusername/8Puzzle_BidirectionalBFS.git
cd 8Puzzle_BidirectionalBFS
```

### Install Dependencies

```sh
go mod tidy
```

## Running the Application

### Command-Line Execution

To run the solver in terminal mode:

```sh
go run main.go
```

### Running the GUI

To launch the graphical interface:

```sh
go run gui.go
```

## Usage

### Command-Line Mode

1. The program initializes a starting board.
2. Uses **bidirectional BFS** to find the shortest path to the goal state.
3. Outputs the sequence of moves.

### GUI Mode

1. Choose to **start a new game** (random or custom board).
2. Click **Start** to begin solving.
3. Pause, reset, or step through the solution manually.

## How It Works

### Bidirectional BFS Algorithm

The program:
- Starts two simultaneous BFS searches: one from the initial board and one from the goal state.
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

### GUI Implementation

The graphical interface:
- Uses **Fyne** to display the puzzle.
- Animates each move.
- Allows interaction through buttons.

## Future Improvements

- Implement A* search for comparison.
- Optimize performance further.
- Add support for different board sizes (e.g., 15-puzzle).

## License

This project is licensed under the **MIT License**.

