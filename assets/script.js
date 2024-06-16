document.addEventListener('DOMContentLoaded', function() {
    const boardElement = document.getElementById('sudoku-board');

    // Get the boards from the server
    fetch("boards")
        .then((response) => {
            if (response.ok !== true) {
                throw Error("Could not get board :(");
            }

            return response.json();
        })
        .then((boards) => {
            drawBoard(boards);
        })
        .catch((error) => {
            console.error(error);
        });

    // Draws the Sudoku board
    function drawBoard(boards) {
        console.log(boards.complete_board);

        for (let i = 0; i < 9; i++) {
            for (let j = 0; j < 9; j++) {
                const cell = document.createElement('div');
                cell.classList.add('sudoku-cell');
                cell.contentEditable = true;
                if (boards.playable_board[i][j] !== 0) {
                    cell.textContent = boards.playable_board[i][j];
                    cell.classList.add('initial');
                    cell.contentEditable = false;
                }

                cell.oninput = function(event) {
                    if (this.textContent.length > 1) {
                        this.textContent = this.textContent.slice(0, 1);
                    }

                    this.textContent = this.textContent.replace(/[^\d]/g, '');

                    // Convert the Sudoku board from HTML to JSON
                    const board = []
                    for (let i = 0; i < boardElement.childElementCount; i++) {
                        const child = boardElement.children[i];
                        const cellNum = child.textContent === "" ? 0 : parseInt(child.textContent);
                        const rowIndex = Math.floor(i / 9);

                        let row = board.at(rowIndex);
                        if (row === undefined) {
                            row = [cellNum];
                            board.push(row);
                            continue;
                        }

                        row.push(cellNum);
                    }

                    if (arraysEqual(board, boards.complete_board)) {
                        alert("muy bien jodio bastardo");
                    }
                };

                boardElement.appendChild(cell);
            }
        }
    }

    function arraysEqual(arr1, arr2) {
        // Check if arrays are of different length
        if (arr1.length !== arr2.length) {
            return false;
        }

        // Deep comparison using JSON.stringify
        for (let i = 0; i < arr1.length; i++) {
            if (JSON.stringify(arr1[i]) !== JSON.stringify(arr2[i])) {
                return false;
            }
        }

        return true;
    }
});
