package main

import (
	"github.com/a-h/templ"
	"html/template"
	"net/http"
	"os"
	"src/pkg/sudoku"
	my "src/web/template"
)

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed:"+request.Method, http.StatusMethodNotAllowed)
		return
	}

	var unsolvedBoard = sudoku.NewBoardForDifficulty(sudoku.MEDIUM)
	var templComponent = my.SudokuBoardTemplate(*unsolvedBoard)

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

	var templComponent = my.SudokuBoardTemplate(*board)
	if err := templComponent.Render(request.Context(), writer); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	var staticFilesystem = http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFilesystem))

	http.HandleFunc("/", HandleIndex)

	http.HandleFunc("/board/{difficulty}", HandleBoard)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
