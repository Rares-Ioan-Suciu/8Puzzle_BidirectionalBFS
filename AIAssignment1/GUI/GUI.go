package GUI

import (
	"AIAssignment1/board"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"time"
)

func getSteps(steps []*board.BoardState, meetIndex int) ([][][]int, int) {
	result := make([][][]int, 0)
	for _, step := range steps {
		result = append(result, step.Board)
	}
	return result, meetIndex
}

func StartWindow() {
	EightPuzzle := app.New()
	Screen := EightPuzzle.NewWindow("8 Puzzle Bidirectional BFS")
	var boardStates [][][]int
	var meetIndex int
	boardStates = append(boardStates, [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}})

	var gameContainer *fyne.Container

	moveIndex := 0
	isRunning := false

	statusLabel := widget.NewLabelWithStyle("Welcome! Choose an option to start.", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	customButton := func(label string, color color.RGBA, icon fyne.Resource, action func()) *fyne.Container {
		btn := widget.NewButtonWithIcon(label, icon, action)
		btnRect := canvas.NewRectangle(color)
		btnContainer := container.NewStack(btnRect, btn)
		btnRect.Resize(btn.MinSize())
		return btnContainer
	}

	updateGame := func(board [][]int) fyne.CanvasObject {
		var gameObject []fyne.CanvasObject
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				val := board[i][j]

				rect := canvas.NewRectangle(color.RGBA{R: 50, G: 180, B: 55, A: 255})
				rect.SetMinSize(fyne.NewSize(100, 100))

				text := canvas.NewText(string(val), color.White)
				if val != 0 {
					text = canvas.NewText(string(rune('0'+val)), theme.TextColor())
				} else {
					text = canvas.NewText("", theme.TextColor())
					rect.FillColor = color.RGBA{R: 100, G: 50, B: 55, A: 255}
				}
				text.TextSize = 32
				text.Alignment = fyne.TextAlignCenter

				gameObject = append(gameObject, container.NewStack(rect, text))
			}
		}
		return container.NewGridWithColumns(3, gameObject...)
	}

	startButton := customButton("Start Game", color.RGBA{R: 0, G: 200, B: 0, A: 255}, theme.MediaPlayIcon(), func() {
		if isRunning {
			statusLabel.SetText("Already Running")
			return
		}
		isRunning = true
		statusLabel.SetText("Game started!")

		go func() {
			for i := 0; i < len(boardStates); i++ {
				time.Sleep(1 * time.Second)
				gameContainer.Objects = []fyne.CanvasObject{updateGame(boardStates[i])}
				gameContainer.Refresh()
				moveIndex++

				if moveIndex == meetIndex {
					isRunning = false
					dialog.ShowInformation("Meet Point", "The bidirectional searches meet at this step!", Screen)

					time.Sleep(4000 * time.Millisecond)
					isRunning = true

				}

				for !isRunning {
					time.Sleep(100 * time.Millisecond)
				}
			}
			statusLabel.SetText("Game Over!")
		}()
	})

	pauseButton := widget.NewButtonWithIcon("Pause Game", theme.MediaStopIcon(), func() {
		if !isRunning {
			statusLabel.SetText("Game starts again")
			isRunning = true
			return
		}
		isRunning = false
		statusLabel.SetText("Game Paused!")
	})

	stepButton := widget.NewButtonWithIcon("Step", theme.NavigateNextIcon(), func() {
		if moveIndex < len(boardStates)-1 {
			gameContainer.Objects = []fyne.CanvasObject{updateGame(boardStates[moveIndex+1])}
			gameContainer.Refresh()
			moveIndex++
			statusLabel.SetText(fmt.Sprintf("Move %d", moveIndex))
		}
	})

	resetButton := widget.NewButtonWithIcon("Reset", theme.ViewRefreshIcon(), func() {
		moveIndex = 0
		isRunning = false
		gameContainer.Objects = []fyne.CanvasObject{updateGame(boardStates[0])}
		gameContainer.Refresh()
		statusLabel.SetText("Game Reset!")

	})

	randomGameButton := widget.NewButton("Random Game", func() {
		statusLabel.SetText("Generating random game...")
		boardStates, meetIndex = getSteps(board.BidirectionalBFS(board.RandomInitial()))
		statusLabel.SetText(fmt.Sprintf("Game generated !"))
		gameContainer.Objects = []fyne.CanvasObject{updateGame(boardStates[0])}
	})

	userInput := widget.NewEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Input board:", Widget: userInput},
		},
	}

	submitButton := widget.NewButton("Submit board!", func() {
		statusLabel.SetText("Generating random game...")
		inputBoard := board.StrBoard(userInput.Text)
		if inputBoard == nil {
			dialog.ShowError(errors.New("The input is empty"), Screen)
			return
		}

		if board.CheckBoard(inputBoard) {
			boardStates, meetIndex = getSteps(board.BidirectionalBFS(inputBoard))
			statusLabel.SetText("Game generated!")
			gameContainer.Objects = []fyne.CanvasObject{updateGame(boardStates[0])}
		} else {
			fmt.Println("Input board cannot be solved!")
		}
	})

	formContainer := container.NewVBox(
		form,
		container.NewCenter(submitButton),
	)

	form.SubmitText = "Submit board!"

	gameContainer = container.NewCenter(updateGame(boardStates[0]))

	menu := container.NewVBox(
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Select Game Mode:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		randomGameButton,
		formContainer,
	)

	statusBar := container.NewHBox(
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Game Status:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		statusLabel,
		layout.NewSpacer(),
	)

	gamePanel := container.NewVBox(
		statusBar,
		container.NewGridWithColumns(2, startButton, pauseButton),
		container.NewGridWithColumns(2, stepButton, resetButton),
	)

	background := canvas.NewRectangle(color.RGBA{R: 198, G: 226, B: 44, A: 200})
	background.Resize(fyne.NewSize(800, 600))

	mainLayout := container.NewStack(
		background,
		container.NewBorder(menu, gamePanel, nil, nil, gameContainer),
	)

	Screen.Resize(fyne.NewSize(800, 600))
	Screen.CenterOnScreen()
	Screen.SetContent(mainLayout)

	Screen.ShowAndRun()
}
