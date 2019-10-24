package letterboard

func position(board Board, r rune) (int, Direction) {
	var li, ri int

	for i := 0; i < len(board); i++ {
		if board[i] == r {
			li = i
		}
	}

	for i := len(board) - 1; i >= 0; i-- {
		if board[i] == r {
			ri = i
		}
	}

	if len(board[0:li]) < len(board[ri+1:]) {
		return li, Left
	} else if len(board[0:li]) > len(board[ri+1:]) {
		return ri, Right
	}

	return li, Left
}

func SolveLetterboard(board Board, word string) Moves {
	var moves []Move

	var pos int
	var direction Direction

	for _, wr := range word {
		pos, direction = position(board, wr)

		if direction == Left {
			for i := 0; i < pos; i++ {
				moves = append(moves, Move{
					Direction: Left,
					Letter:    -1,
				})
			}

			board = append(board[pos+1:], board[0:pos]...)
		}

		if direction == Right {
			for i := len(board) - 1; i > len(board)-pos; i-- {
				moves = append(moves, Move{
					Direction: Right,
					Letter:    -1,
				})
			}

			board = append(board[len(board)-pos+1:], board[0:len(board)-pos]...)
		}

		moves = append(moves, Move{
			Direction: direction,
			Letter:    wr,
		})
	}

	return moves
}
