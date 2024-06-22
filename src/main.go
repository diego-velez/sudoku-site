package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// Boards represent a JSON response with 2 variations of the same Sudoku board, one that is
// completely filled, and one that has empty cells.
type Boards struct {
	SolvedBoard   *Board `json:"solved_board"`
	UnsolvedBoard *Board `json:"unsolved_board"`
}

func HandleIndex(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "assets/index.html")
}

func HandleBoards(writer http.ResponseWriter, request *http.Request) {
	var difficulty = request.PathValue("difficulty")
	var amountToRemove int

	switch difficulty {
	case "easy":
		amountToRemove = 10
	case "medium":
		amountToRemove = 20
	case "hard":
		amountToRemove = 30
	}

	var board = NewBoard()

	var boards = Boards{
		SolvedBoard:   board,
		UnsolvedBoard: NewBoardRemoveNumbers(*board, amountToRemove),
	}

	json.NewEncoder(writer).Encode(boards)
}

func main() {
	var fs = http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HandleIndex)

	http.HandleFunc("/boards/{difficulty}", HandleBoards)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
