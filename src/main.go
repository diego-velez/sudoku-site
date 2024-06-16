package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Boards struct {
	CompleteBoard [9][9]int `json:"complete_board"`
	PlayableBoard [9][9]int `json:"playable_board"`
}

func handling(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "assets/index.html")
}

func getBoard(writer http.ResponseWriter, request *http.Request) {
	var board = *generateBoard(nil, getSequence(), 0, 0)

	var boards = Boards{
		CompleteBoard: board,
		PlayableBoard: removeNumbers(board, 5),
	}

	json.NewEncoder(writer).Encode(boards)
}

func main() {
	var fs = http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handling)

	http.HandleFunc("/boards", getBoard)

	var port = os.Getenv("PORT")

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
