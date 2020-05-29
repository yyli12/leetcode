package algorithm

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
