func main() {
	n := 10
	statisticPlot(solveNQueens(n), n)
}

func solveNQueens(n int) [][]string {
	var chessBoard []string
	result, ok := createRowPermutations(chessBoard, n)
	if !ok {
		return nil
	}
	return result
}

func statisticPlot(permutations [][]string, n int) {
	var counts [][]int
	for i := 0; i < n; i++ {
		counts = append(counts, make([]int, n))
	}
	for _, permutation := range permutations {
		for i, row := range permutation {
			for j, column := range row {
				if column == 'Q' {
					counts[i][j]++
				}
			}
		}
	}
	for _, row := range counts {
		sum := 0
		for _, column := range row {
			fmt.Print(column, " ")
			sum += column
		}
		fmt.Println()
		fmt.Println(sum)
	}
}

func createRowPermutations(chessBoard []string, n int) ([][]string, bool) {
	if len(chessBoard) == n {
		return [][]string{}, true
	}
	var rows [][]string
	for j := 0; j < n; j++ {
		if isPlaceValid(chessBoard, len(chessBoard), j) {
			row := []byte(createEmptyRow(n))
			row[j] = 'Q'
			result, ok := createRowPermutations(append(chessBoard, string(row)), n)
			if ok {
				if len(result) > 0 {
					for _, r := range result {
						newRow := []string{string(row)}
						newRow = append(newRow, r...)
						rows = append(rows, newRow)
					}
				} else {
					newRow := []string{string(row)}
					rows = append(rows, newRow)
				}
			}
		}
	}
	if len(rows) == 0 {
		return rows, false
	}
	return rows, true
}

func createEmptyRow(n int) string {
	var row string
	for j := 0; j < n; j++ {
		row += "."
	}
	return row
}

func isPlaceValid(chessBoard []string, i, j int) bool {
	if i == 0 {
		return true
	}
	n := len(chessBoard[0])
	for k := 0; k < i; k++ {
		if chessBoard[k][j] == 'Q' {
			return false
		}
	}
	left, right := j-1, j+1
	for row := i - 1; row >= 0; row-- {
		if (left >= 0 && chessBoard[row][left] == 'Q') || (right < n && chessBoard[row][right] == 'Q') {
			return false
		}
		left--
		right++
		if left < 0 && right >= n {
			break
		}
	}
	return true
}
