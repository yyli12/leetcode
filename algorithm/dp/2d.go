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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	R, C := len(word1), len(word2)

	dp := make([]int, (R+1)*(C+1))
	set := func(r, c int, v int) {
		dp[r*(C+1)+c] = v
	}
	get := func(r, c int) int {
		return dp[r*(C+1)+c]
	}

	for r := 0; r <= R; r++ {
		set(r, 0, r)
	}
	for c := 0; c <= C; c++ {
		set(0, c, c)
	}

	for r := 1; r <= R; r++ {
		cr := word1[r-1]
		for c := 1; c <= C; c++ {
			cc := word2[c-1]
			v := 0
			if cr == cc {
				v = min(get(r-1, c-1), min(get(r-1, c), get(r, c-1))+1)
			} else {
				v = min(get(r-1, c-1), min(get(r-1, c), get(r, c-1))) + 1
			}
			set(r, c, v)
		}
	}
	return get(R, C)
}

func MinDistance(word1 string, word2 string) int {
	return minDistance(word1, word2)
}
