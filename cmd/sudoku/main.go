package main

import (
	"net/http"
	"os"
	"src/api/handler"
)

func main() {
	var staticFilesystem = http.FileServer(http.Dir("web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFilesystem))

	http.HandleFunc("/", handler.HandleIndex)

	http.HandleFunc("/board/{difficulty}", handler.HandleBoard)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
