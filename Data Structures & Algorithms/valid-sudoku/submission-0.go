func isValidSudoku(board [][]byte) bool {
	row := make([]map[byte]bool, 9)
	column := make([]map[byte]bool, 9)
	box := make([]map[byte]bool, 9)

	for i:= 0;i < 9 ; i++ {
		row[i] = make(map[byte]bool, 9)
		column[i] = make(map[byte]bool, 9)
		box[i] = make(map[byte]bool, 9)
	}

	for r := 0 ; r < 9; r++ {
		for c := 0; c < 9; c++ {
			num := board[r][c]
			if num == '.' {
				continue
			}
			boxNo := (r / 3) * 3 + (c / 3)

			if row[r][num] || column[c][num] || box[boxNo][num] {
				return false
			}
			row[r][num] = true
			column[c][num] = true
			box[boxNo][num] = true 
		}
	}
	return true
}
