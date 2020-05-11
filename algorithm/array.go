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
