package main

import (
	"html/template"
	"net/http"
	"os"
	"src/pkg/sudoku"
)

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed:"+request.Method, http.StatusMethodNotAllowed)
		return
	}

	var unsolvedBoard = sudoku.NewBoardForDifficulty(sudoku.MEDIUM)

	templ, err := template.ParseFiles("web/template/index.gohtml", "web/template/com_sudoku_board.gohtml")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templ.ExecuteTemplate(writer, "index.gohtml", unsolvedBoard); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
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

	templ, err := template.ParseFiles("assets/com_sudoku_board.gohtml")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.ExecuteTemplate(writer, "sudoku_board", board)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
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
