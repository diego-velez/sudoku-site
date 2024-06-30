package handler

import (
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"sudoku/pkg/sudoku"
	"testing"
)

func TestBoardHandler(t *testing.T) {
	expectedBoard := &[81]int{
		0, 5, 0, 7, 9, 1, 4, 0, 8,
		0, 9, 0, 3, 0, 4, 0, 6, 1,
		7, 0, 1, 0, 8, 6, 3, 9, 5,
		3, 2, 4, 5, 0, 8, 6, 1, 9,
		9, 8, 6, 0, 0, 0, 0, 7, 2,
		1, 7, 5, 6, 2, 9, 8, 3, 0,
		4, 1, 2, 8, 3, 0, 9, 0, 6,
		8, 3, 0, 1, 6, 5, 2, 4, 7,
		0, 6, 7, 9, 4, 0, 1, 8, 3,
	}

	// By default, tests change the working directory to the directory where the test resides
	// instead of the directory of the project like normal. This should change the working
	// directory to the normal one.
	// see:
	// https://stackoverflow.com/questions/23847003/golang-tests-and-working-directory
	// https://www.reddit.com/r/golang/comments/rquwhk/how_to_deal_with_changing_working_directory_while/
	err := os.Chdir("..")
	if err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/board/{difficulty}", HandleBoard)

	rand.Seed(1)

	request, err := http.NewRequest(http.MethodGet, "/board/medium", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	mux.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := responseRecorder.Body.String()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	doc.Find("#sudoku-expectedBoard").Children().Each(func(i int, s *goquery.Selection) {
		// Make sure all children are divs
		isDiv := s.Is("div")
		if !isDiv {
			elementType := s.Get(0).Data
			t.Errorf("Expected all elements in sudoku-expectedBoard to be divs, got %v", elementType)
		}

		// Make sure all children are the appropriate class
		value, exists := s.Attr("class")
		if !exists {
			t.Errorf("Expected all elements in sudoku-expectedBoard to have a class attribute")
		} else if value != "sudoku-cell" {
			t.Errorf("Expected all elements in sudoku-expectedBoard to have class=sudoku-cell, got %v", value)
		}

		// Get cell number from int
		var cellNumber int
		var err error
		if s.Text() == "" {
			// Should be editable if empty
			value, exists = s.Attr("contenteditable")
			if !exists {
				t.Errorf("Expected all empty elements in sudoku-expectedBoard to have a contenteditable attribute")
			} else if value != "true" {
				t.Errorf("Expected all empty elements in sudoku-expectedBoard to be editable")
			}

			// Should have appropriate oninput callback
			value, exists = s.Attr("oninput")
			if !exists {
				t.Errorf("Expected all empty elements in sudoku-expectedBoard to have an oninput attribute")
			} else if value != "inputHandler(this)" {
				t.Errorf("Expected all empty elements in sudoku-expectedBoard to call inputHandler(this) on input, got %v", value)
			}

			cellNumber = int(sudoku.EMPTY)
		} else {
			cellNumber, err = strconv.Atoi(s.Text())
		}

		if err != nil {
			t.Errorf("Error converting to int:%v", err)
		}

		// Ensure correct number
		if cellNumber != expectedBoard[i] {
			t.Errorf("Wrong cell number: got %v want %v", expectedBoard[i], cellNumber)
		}
	})
}
