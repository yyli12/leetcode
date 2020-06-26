package algorithm

import "fmt"

/*
https://leetcode.com/problems/maximum-swap/
*/

func maximumSwap(num int) int {
	if num == 0 {
		return 0
	}

	leastDigitIndex := map[int]int{}
	for i := 0; i < 10; i++ {
		leastDigitIndex[i] = -1
	}
	digits := make([]int, 0, 15)
	maxDigit := -1
	index := -1
	n := num
	for n > 0 {
		index++
		digit := n % 10
		n /= 10

		if leastDigitIndex[digit] == -1 {
			leastDigitIndex[digit] = index
		}
		if digit > maxDigit {
			maxDigit = digit
		}
		digits = append(digits, digit)
	}
	fmt.Println(digits)

	index = len(digits) - 1
	done := false
	for index > 0 && !done {
		digit := digits[index]
		if digit < maxDigit {
			for lDigit := maxDigit; lDigit > digit; lDigit-- {
				if leastDigitIndex[lDigit] != -1 && leastDigitIndex[lDigit] < index {
					digits[index] = lDigit
					digits[leastDigitIndex[lDigit]] = digit
					done = true
					break
				}
			}
			maxDigit = digit
		}
		index--
	}
	fmt.Println(digits)

	newNum := 0
	for index = len(digits) - 1; index > -1; index-- {
		newNum = 10*newNum + digits[index]
	}
	return newNum
}

func MaximumSwap(num int) int {
	return maximumSwap(num)
}

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	l, r := 1, n+1

	for l+1 < r {
		less, equal, greater := 0, 0, 0
		m := (l + r) / 2
		for _, num := range nums {
			if num == m {
				equal++
				if equal > 1 {
					break
				}
			} else if num < m {
				less++
			} else {
				greater++
			}
		}

		if equal > 1 {
			return m
		} else if less > m-1 {
			r = m
		} else if greater > n-m {
			l = m + 1
		}
	}
	return l
}

func FindDuplicate(nums []int) int {
	return findDuplicate2(nums)
}

func findDuplicate2(nums []int) int {
	var result int
	for i := uint(0); i < 31; i++ {
		bit := 1 << i
		expected, actual := 0, 0
		for j := 0; j < len(nums); j++ {
			if j&bit > 0 {
				expected++
			}
			if nums[j]&bit > 0 {
				actual++
			}
		}
		fmt.Println(expected, actual)
		if actual > expected {
			result ^= bit
		}
	}
	return result
}
