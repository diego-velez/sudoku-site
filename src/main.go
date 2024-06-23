package main

import (
	"html/template"
	"net/http"
	"os"
)

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed:"+request.Method, http.StatusMethodNotAllowed)
		return
	}

	var unsolvedBoard = NewBoardForDifficulty(MEDIUM)

	templ, err := template.ParseFiles("assets/index.gohtml", "assets/com_sudoku_board.gohtml")
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

	var board *Board
	switch difficulty {
	case "easy":
		board = NewBoardForDifficulty(EASY)
	case "medium":
		board = NewBoardForDifficulty(MEDIUM)
	case "hard":
		board = NewBoardForDifficulty(HARD)
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
	var fs = http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

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
