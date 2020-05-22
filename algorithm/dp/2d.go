package dp

func CountSquares(matrix [][]int) int {
	return countSquares(matrix)
}

func countSquares(matrix [][]int) int {
	R := len(matrix)
	if R == 0 {
		return 0
	}
	C := len(matrix[0])
	if C == 0 {
		return 0
	}

	count := 0
	checkSquareCorner := func(isSquareCorner []bool, r, c int) bool {
		return isSquareCorner[r*C+c]
	}
	setSquareCorner := func(squareCorners [][2]int, isSquareCorner []bool, r, c int) [][2]int {
		count++
		squareCorners = append(squareCorners, [2]int{r, c})
		isSquareCorner[r*C+c] = true
		return squareCorners
	}

	isSquareCorner := make([]bool, R*C)
	squareCorners := make([][2]int, 0, R*C)
	for r, row := range matrix {
		for c, elem := range row {
			if elem == 1 {
				squareCorners = setSquareCorner(squareCorners, isSquareCorner, r, c)
			}
		}
	}

	for len(squareCorners) > 0 {
		nextIsSquareCorner := make([]bool, R*C)
		nextSquareCorners := make([][2]int, 0, R*C)
		for _, rc := range squareCorners {
			r, c := rc[0], rc[1]
			if r+1 < R && c+1 < C && matrix[r+1][c+1] == 1 {
				up := checkSquareCorner(isSquareCorner, r, c+1)
				left := checkSquareCorner(isSquareCorner, r+1, c)
				if up && left {
					nextSquareCorners = setSquareCorner(nextSquareCorners, nextIsSquareCorner, r+1, c+1)
				}
			}
		}
		isSquareCorner = nextIsSquareCorner
		squareCorners = nextSquareCorners
	}
	return count
}
