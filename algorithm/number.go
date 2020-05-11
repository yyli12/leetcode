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
