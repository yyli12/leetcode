package algorithm

import (
	"fmt"
	"sort"
)

func ThreeEqualParts(A []int) []int {
	return threeEqualParts(A)
}

func threeEqualParts(A []int) []int {
	N := len(A)
	jValidBits := 1
	for j := N - 1; j >= 2; j-- {
		if A[j] == 1 {
			jValidBits = N - j
		}
		if jValidBits*3 > N {
			break
		}

		// check A[j:] == A[j-jValidBits:j]
		eq2 := true
		for bit := 1; bit <= jValidBits && eq2; bit++ {
			if A[N-bit] != A[j-bit] {
				eq2 = false
			}
		}
		if !eq2 {
			continue
		}

		possible := true
		for i := j - jValidBits - 1; i >= 0 && possible; i-- {
			eq1 := true
			for bit := 1; bit <= jValidBits && eq1; bit++ {
				if A[N-bit] != A[i-bit+1] {
					eq1 = false
				}
			}
			if !eq1 {
				if A[i] == 1 {
					possible = false
				} else {
					possible = true
				}
			} else {
				// the rest should be all zero
				allZero := true
				for z := i - jValidBits; z >= 0 && allZero; z-- {
					if A[z] != 0 {
						allZero = false
					}
				}
				if allZero {
					return []int{i, j}
				} else {
					if A[i] == 1 {
						possible = false
					} else {
						possible = true
					}
				}
			}
		}
	}
	return []int{-1, -1}
}

func FindMaxLength(nums []int) int {
	return findMaxLength(nums)
}

func findMaxLength(nums []int) int {
	n := len(nums)
	n1 := 0
	for _, num := range nums {
		n1 += num
	}
	n0 := n - n1
	if n0 == n1 {
		return n
	}
	if n0 == 0 || n1 == 0 {
		return 0
	}

	result := 0
	for i := 0; i < n; i++ {
		count := 0
		for j := i; j < n; j++ {
			count += nums[j]
			if (j-i)%2 == 0 && count == (j-i)/2 && j-i > result {
				result = j - i
			}
		}
	}

	return result
}

func PossibleBipartition(N int, dislikes [][]int) bool {
	return possibleBipartition(N, dislikes)
}

func possibleBipartition(N int, dislikes [][]int) bool {
	group := make([]int, N+1)
	dislikeTable := make([][]int, N+1)
	for _, dislike := range dislikes {
		i, j := dislike[0], dislike[1]
		dislikeTable[i] = append(dislikeTable[i], j)
		dislikeTable[j] = append(dislikeTable[j], i)
	}

	checkAndMark := func(start int, groupID int) bool {
		thisGroup := map[int]struct{}{start: {}}
		for len(thisGroup) > 0 {
			nextGroup := map[int]struct{}{}
			for i := range thisGroup {
				if group[i] != 0 {
					if group[i] != groupID {
						return false
					}
				} else {
					group[i] = groupID
					for _, j := range dislikeTable[i] {
						nextGroup[j] = struct{}{}
					}
				}
			}
			thisGroup = nextGroup
			groupID = -groupID
		}
		return true
	}

	for i := 1; i <= N; i++ {
		if group[i] == 0 && !checkAndMark(i, 1) {
			return false
		}
	}
	return true
}

func twoCitySchedCost(costs [][]int) int {
	N := len(costs) >> 1
	sum := 0
	diff := make([]int, N<<1)
	for i, cost := range costs {
		sum += cost[0]
		diff[i] = cost[1] - cost[0]
	}
	sort.Ints(diff)
	for i := 0; i < N; i++ {
		sum += diff[i]
	}
	return sum
}

func TwoCitySchedCost(costs [][]int) int {
	return twoCitySchedCost(costs)
}

func solve(board [][]byte) {
	old := make([][]byte, len(board))
	for index, row := range board {
		old[index] = make([]byte, len(row))
		copy(old[index], row)
		for i := 0; i < len(row); i++ {
			row[i] = 'X'
		}
	}
	var paint func(i, j int)
	paint = func(i, j int) {
		if 0 <= i && i < len(board) && 0 <= j && j < len(board[0]) {
			if old[i][j] == 'O' {
				fmt.Println(i, j, 'O')
				old[i][j] = 'X'
				board[i][j] = 'O'
				paint(i-1, j)
				paint(i+1, j)
				paint(i, j-1)
				paint(i, j+1)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		paint(i, 0)
		paint(i, len(board[0])-1)
	}
	for j := 0; j < len(board[0]); j++ {
		paint(0, j)
		paint(len(board)-1, j)
	}
}

func Solve(board [][]byte) {
	solve(board)
}

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	N := len(citations)
	if citations[N-1] == 0 {
		return 0
	}
	if citations[0] >= N {
		return N
	}
	l := 1
	r := N
	if citations[N-1] < r {
		r = citations[N-1]
	}
	r++
	for r-l > 1 {
		h := (l + r) / 2
		if N-1 >= h {
			b := citations[N-h]
			a := citations[N-h-1]
			if a > h {
				l = h + 1
			} else if b < h {
				r = h
			} else {
				l = h
			}
		} else {
			b := citations[N-h]
			if b < h {
				r = h
			} else {
				l = h
			}
		}
	}
	return l
}

func HIndex(citations []int) int {
	return hIndex(citations)
}
