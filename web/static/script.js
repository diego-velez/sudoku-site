const EMPTY = 0;

// Handles div input for each editable div within the Sudoku board
function inputHandler(div) {
    const boardElement = document.getElementById('sudoku-board');

    // Limit input to 1 character
    if (div.textContent.length > 1) {
        div.textContent = div.textContent.slice(0, 1);
    }

    // Limit input to numbers only
    div.textContent = div.textContent.replace(/[^\d]/g, '');

    // Check if the Sudoku board is complete and correct by converting it to its 2-dimensional array representation, and
    // validating each cell as it is being converted.
    const board = []
    for (let i = 0; i < boardElement.childElementCount; i++) {
        const childElement = boardElement.children[i];
        const rowIndex = Math.floor(i / 9);

        // The number in the cell, will be converted to the EMPTY representation if needed
        const num = childElement.textContent === "" ? EMPTY : parseInt(childElement.textContent);

        // Checks if there is still an empty cell in the board
        if (num === EMPTY) {
            console.log("Board is not complete")
            return;
        }

        // Checks if the row exists in the array representation, add it if not
        let row = board[rowIndex];
        if (row === undefined) {
            row = [num];
            board.push(row);

            // TODO: Maybe we could check if cell is valid this iteration
            continue;
        }

        // Checks if the cell is valid with the construction of the board so far
        if (!validCell(board, num, rowIndex, row.indexOf(num))) {
            console.log("Board invalid :(");
            return;
        }

        row.push(num);
    }

    alert("muy bien jodio bastardo");
}

// Reports if the cell is valid, works with a board that hasn't been fully constructed yet.
function validCell(board, num, row, column) {
    // Check if the number was played in the same row or column
    for (let i = 0; i < board.length; i++) {
        if (board[row][i] === num || board[i][column] === num) {
            return false;
        }
    }

    // Search for the 3x3 grid that the number is in
    let rowStart = null, rowEnd = null;
    let columnStart = null, columnEnd = null;

    for (let start = 0; start < 9; start += 3) {
        const end = start + 2;

        if (start <= row <= end) {
            rowStart = start;

            if (board[end] === undefined) {
                rowEnd = board.length - 1;
            } else {
                rowEnd = end;
            }
        }

        if (start <= column <= end) {
            columnStart = start;

            if (board[row][end] === undefined) {
                columnEnd = board[row].length - 1;
            } else {
                columnEnd = end;
            }
        }
    }

    // Search the 3x3 grid for the number
    for (let thisRow = rowStart; thisRow <= rowEnd; thisRow++) {
        for (let thisColumn = columnStart; thisColumn <= columnEnd; thisColumn++) {
            if (board[thisRow][thisColumn] === num) {
                return false;
            }
        }
    }

    return true;
}
