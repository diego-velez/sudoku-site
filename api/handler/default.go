package handler

import (
	"github.com/a-h/templ"
	"html/template"
	"net/http"
	"sudoku/pkg/sudoku"
	customTemplates "sudoku/web/template"
)

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	var unsolvedBoard = sudoku.NewBoardForDifficulty(sudoku.MEDIUM)
	var templComponent = customTemplates.SudokuBoardTemplate(*unsolvedBoard)

	var html, err = templ.ToGoHTML(request.Context(), templComponent)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	myTempl, err := template.ParseFiles("web/template/index.gohtml")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := myTempl.Execute(writer, html); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func HandleBoard(writer http.ResponseWriter, request *http.Request) {
	var difficulty = request.PathValue("difficulty")

	var board *sudoku.Board
	switch difficulty {
	case "easy":
		board = sudoku.NewBoardForDifficulty(sudoku.EASY)
	case "medium":
		board = sudoku.NewBoardForDifficulty(sudoku.MEDIUM)
	case "hard":
		board = sudoku.NewBoardForDifficulty(sudoku.HARD)
	default:
		http.Error(writer, "Invalid difficulty:"+difficulty, http.StatusBadRequest)
		return
	}

	var templComponent = customTemplates.SudokuBoardTemplate(*board)
	if err := templComponent.Render(request.Context(), writer); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
