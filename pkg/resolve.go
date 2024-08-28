package tetrino

// This struct is for the position of the tetrino, given the x and y values of them.
type P struct {
	X int
	Y int
}

// This struct is a save for the position of the tetrinos.
type S struct {
	board [][]string
	pos   [][]P
}

// Solves the given problem after all the checks were complete.
func Solve(t [][]string, board [][]string) [][]string {
	bs := []S{}
	for i := 0; i < len(t); i++ {
		placed, dupe, canplace, lastpos := false, false, false, true
		pos := [][]P{}
		for h := 0; h < len(board[0]) && !placed; h++ {
			for v := 0; v < len(board) && !placed; v++ {
				if Possible(t[i], board, h, v) && !placed {
					canplace = true
					hasDuplicate := false
					if len(bs) > 0 && len(bs) == i+1 {
						for _, pos := range bs[len(bs)-1].pos[len(bs[len(bs)-1].pos)-1] {
							if pos.X == h && pos.Y == v {
								hasDuplicate = true
								dupe = true
							}
						}
					}
					if len(bs) > 0 {
						pos = bs[len(bs)-1].pos
					}
					if hasDuplicate {
						continue
					}
					if !dupe {
						pos = append(pos, []P{})
					}
					placed = true
					lb := ReturnBoard(board)
					pos[len(pos)-1] = append(pos[len(pos)-1], P{h, v})
					if !dupe {
						lastpos = false
						bs = append(bs, S{lb, pos})
					}
					board = PlaceTetromino(t[i], board, h, v, i)
				}
			}
		}

		if !canplace {
			lastpos = false
		}

		if !placed {
			i--
			if len(bs) <= 1 {
				i = -1 
				boardSize := len(board) + 1
				board = make([][]string, boardSize)
				for i := range board {
					board[i] = make([]string, boardSize)
					for j := range board[i] {
						board[i][j] = "."
					}
				}
				bs = []S{}
			}
			if len(bs) > 1 {
				i-- 
				if lastpos {
					{
						bs = bs[:len(bs)-1]
						board = ReturnBoard(bs[len(bs)-1].board)
					}
				} else {
					board = ReturnBoard(bs[len(bs)-1].board)
				}
			}
		}
	}

	return board
}

// Checks if placing the given tetrino is even possible.
func Possible(t []string, b [][]string, iv int, ih int) bool {
	for v, line := range t {
		for h, char := range line {
			if char == '#' {
				if iv+v > len(b)-1 || ih+h > len(b)-1 {
					return false
				}
				if b[iv+v][ih+h] != "." {
					return false
				}
			}
		}
	}
	return true
}

// Places the tetrino within the board.
func PlaceTetromino(tr []string, b [][]string, iv int, ih int, ti int) [][]string {
	for ivt, line := range tr {
		for iht, char := range line {
			if char == '#' {
				b[iv+ivt][ih+iht] = string(rune(65 + ti))
			}
		}
	}
	return b
}

// Simply returns the board that is entered.
func ReturnBoard(b [][]string) [][]string {
	br := [][]string{}
	for _, row := range b {
		newRow := []string{}
		newRow = append(newRow, row...)
		br = append(br, newRow)
	}
	return br
}
