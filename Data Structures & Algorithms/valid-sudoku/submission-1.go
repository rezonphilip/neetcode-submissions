func isValidSudoku(board [][]byte) bool {
    seen := make(map[int]bool)

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            num := board[r][c]

            if num == '.' {
                continue
            }
            boxNo := ( r / 3 ) * 3 + ( c / 3 )
            rowStr := (100 * (r+1)) + int(num) 
            colStr := (10000 * (c+1)) + int(num)
            boxStr := (1000000 * (boxNo+1)) + int(num)
            if seen[rowStr] || seen[colStr] || seen[boxStr] {
                return false
            }

            seen[rowStr] = true
            seen[colStr] = true
            seen[boxStr] = true
        } 
    }

    return true
}