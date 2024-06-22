package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Boards struct {
	CompleteBoard *Board `json:"complete_board"`
	PlayableBoard *Board `json:"playable_board"`
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
		CompleteBoard: board,
		PlayableBoard: NewBoardRemoveNumbers(*board, amountToRemove),
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
