package dp

import "fmt"

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

func change(amount int, coins []int) int {
	solutions := make([][]int, len(coins)+1)
	solutions[0] = make([]int, amount+1)

	for i := 1; i <= len(coins); i++ {
		solutions[i] = make([]int, amount+1)
		solutions[i][0] = 1
		coin := coins[i-1]
		for a := 1; a <= amount; a++ {
			if coin > a {
				solutions[i][a] = solutions[i-1][a]
			} else {
				solutions[i][a] =
					solutions[i-1][a] + // no use of coin
						solutions[i][a-coin] // use this coin
			}
		}
	}
	return solutions[len(coins)][amount]
}

func Change(amount int, coins []int) int {
	return change(amount, coins)
}

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	prevRow := make([]bool, len(t)+1)
	for tIdx := 1; tIdx <= len(t); tIdx++ {
		prevRow[tIdx] = true
	}

	var thisRow []bool

	for sIdx, sChar := range s {
		thisRow = make([]bool, len(t)+1)
		for tIdx, tChar := range t {
			if tIdx < sIdx {
				continue
			}
			if sChar == tChar {
				thisRow[tIdx+1] = prevRow[tIdx] || thisRow[tIdx]
			} else {
				thisRow[tIdx+1] = thisRow[tIdx]
			}
			fmt.Println(string([]rune{sChar}), string([]rune{tChar}), thisRow[tIdx+1])
		}
		fmt.Println(thisRow)
		prevRow = thisRow
	}

	return thisRow[len(t)]
}

func IsSubsequence(s string, t string) bool {
	return isSubsequence(s, t)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	R, C := len(dungeon), len(dungeon[0])

	dungeon[R-1][C-1] = max(1, 1-dungeon[R-1][C-1])

	for r := R - 1; r >= 0; r-- {
		for c := C - 1; c >= 0; c-- {
			if r == R-1 && c == C-1 {
				continue
			}

			m := 1 << 31
			if r+1 < R {
				m = min(m, dungeon[r+1][c])
			}
			if c+1 < C {
				m = min(m, dungeon[r][c+1])
			}

			dungeon[r][c] = max(1, m-dungeon[r][c])
		}
	}

	return dungeon[0][0]

}

func CalculateMinimumHP(dungeon [][]int) int {
	return calculateMinimumHP(dungeon)
}
