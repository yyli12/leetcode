package algorithm

import "fmt"

/*
https://leetcode.com/problems/integer-to-english-words/
*/
func numberSmallerThan100ToWords(num int) string {
	if num > 99 {
		panic("num > 99")
	}
	mapper := map[int]string{
		0:  "Zero",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}

	if words, ok := mapper[num]; ok {
		return words
	}

	tyMapper := map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
	ty := num / 10
	one := num % 10
	if one == 0 {
		return tyMapper[ty]
	}
	return fmt.Sprintf("%s %s", tyMapper[ty], mapper[one])
}
func numberSmallerThan1000ToWords(num int) string {
	if num > 999 {
		panic("num > 999")
	}
	if num < 100 {
		return numberSmallerThan100ToWords(num)
	}
	mapper := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}
	hundred := num / 100
	ten := num % 100
	words := mapper[hundred] + " Hundred"
	if ten != 0 {
		words += " " + numberSmallerThan100ToWords(ten)
	}
	return words
}
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	i := 0
	suffix := []string{
		"",
		"Thousand",
		"Million",
		"Billion",
	}
	result := ""
	for num > 0 {
		k := num % 1000

		if k != 0 {
			words := numberSmallerThan1000ToWords(k)
			if i == 0 {
				result = words
			} else {
				if result == "" {
					result = fmt.Sprintf("%s %s", words, suffix[i])
				} else {
					result = fmt.Sprintf("%s %s %s", words, suffix[i], result)
				}
			}
		}
		num = num / 1000
		i += 1
	}
	return result
}
func NumberToWords(num int) string {
	return numberToWords(num)
}

/*
https://leetcode.com/problems/word-break/
*/
func wordBreak(s string, wordDict []string) bool {
	can := make([]bool, len(s)+1, len(s)+1)
	can[0] = true
	wordSet := map[string]struct{}{}
	wordLenSet := map[int]struct{}{}
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
		wordLenSet[len(word)] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for wordLen := range wordLenSet {
			if i-wordLen >= 0 {
				fmt.Println(i-wordLen, i)
				word := s[i-wordLen : i]
				fmt.Println(word)
				if _, ok := wordSet[word]; ok {
					can[i] = can[i-wordLen]
				}
				if can[i] {
					break
				}
			}
		}
	}
	fmt.Println(can)
	return can[len(s)]
}

func WordBreak(s string, wordDict []string) bool {
	return wordBreak(s, wordDict)
}

/*
https://leetcode.com/problems/backspace-string-compare/
*/

const EOF = '\n'

type iterator struct {
	str string
	loc int
	bs  int
}

func newIterator(s string) *iterator {
	return &iterator{
		str: s,
		loc: len(s) - 1,
		bs:  0,
	}
}

func (i *iterator) last() byte {
	for i.loc >= 0 {
		if i.str[i.loc] == '#' {
			i.bs++
		} else {
			if i.bs != 0 {
				i.bs--
			} else {
				ret := i.str[i.loc]
				i.loc--
				return ret
			}
		}
		i.loc--
	}
	return EOF
}

func backspaceCompare(S string, T string) bool {
	sIter, tIter := newIterator(S), newIterator(T)
	for {
		s, t := sIter.last(), tIter.last()
		if s != t {
			return false
		} else {
			if s == EOF {
				return true
			}
		}
	}
}

func BackspaceCompare(S string, T string) bool {
	return backspaceCompare(S, T)
}
