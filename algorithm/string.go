package algorithm

import (
	"fmt"
	"sort"
	"strings"
)

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

/*
https://leetcode.com/problems/regular-expression-matching/
*/

func IsMatch(s string, p string) bool {
	return isMatch(s, p)
}

func isMatch(str string, p string) bool {
	type token struct {
		char byte
		star bool
	}

	tokens := make([]token, 0, len(p))
	for i := 0; i < len(p); i++ {
		token := token{
			char: p[i],
		}
		if i+1 < len(p) && p[i+1] == '*' {
			token.star = true
			i++
		}
		tokens = append(tokens, token)
	}

	match := make([][]bool, len(tokens)+1)
	for i := 0; i <= len(tokens); i++ {
		match[i] = make([]bool, len(str)+1)
	}

	match[0][0] = true
	for t := 1; t <= len(tokens); t++ {
		token := tokens[t-1]
		match[t][0] = match[t-1][0] && token.star

		for s := 1; s <= len(str); s++ {
			char := str[s-1]
			if token.star {
				// token is star, can match or skip
				if char == token.char || token.char == '.' {
					match[t][s] = match[t-1][s-1] || match[t][s-1] || match[t-1][s]
				} else {
					match[t][s] = match[t-1][s]
				}
			} else {
				// token is not star, must match
				match[t][s] = (char == token.char || token.char == '.') && match[t-1][s-1]
			}
		}
	}

	//fmt.Println(tokens)
	return match[len(tokens)][len(str)]
}

func FrequencySort(s string) string {
	return frequencySort(s)
}

func frequencySort(s string) string {
	ct := make([]int, 1<<8)
	cs := make([]byte, 0, 1<<8)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if ct[c] == 0 {
			cs = append(cs, c)
		}
		ct[c]++
	}

	sort.SliceStable(cs, func(i, j int) bool {
		return ct[cs[i]] > ct[cs[j]]
	})

	result := make([]byte, 0, len(s))
	for _, c := range cs {
		for ct[c] > 0 {
			result = append(result, c)
			ct[c]--
		}
	}
	return string(result)
}

func validIPv4Segment(segment string) bool {
	if len(segment) > 3 || len(segment) == 0 {
		return false
	}
	if segment == "0" {
		return true
	}
	leadingZero := true
	result := 0
	multiplier := 1
	switch len(segment) {
	case 3:
		multiplier = 100
		break
	case 2:
		multiplier = 10
		break
	}
	for _, digit := range segment {
		if digit > '9' || digit < '0' {
			return false
		}
		if leadingZero && digit == '0' {
			return false
		}
		leadingZero = false
		result += int(digit-'0') * multiplier
		multiplier /= 10
	}
	fmt.Println(result)
	return result < 256
}

func isIPv4(ip string) bool {
	segments := strings.Split(ip, ".")
	if len(segments) != 4 {
		return false
	}
	for _, segment := range segments {
		if !validIPv4Segment(segment) {
			return false
		}
	}
	return true
}

func validIPv6Segment(segment string) bool {
	if len(segment) > 4 || len(segment) == 0 {
		return false
	}
	for _, digit := range segment {
		if (digit > '9' || digit < '0') && (digit > 'f' || digit < 'a') && (digit > 'F' || digit < 'A') {
			return false
		}
	}
	return true
}

func isIPv6(ip string) bool {
	segments := strings.Split(ip, ":")
	if len(segments) != 8 {
		return false
	}
	for _, segment := range segments {
		if !validIPv6Segment(segment) {
			return false
		}
	}
	return true
}

func validIPAddress(ip string) string {
	if isIPv4(ip) {
		return "IPv4"
	}
	if isIPv6(ip) {
		return "IPv6"
	}
	return "Neither"
}

func longestDupSubstring(S string) string {
	N := len(S)
	low := 0
	high := N
	for low < high {
		l := (low + high) / 2
		ok := false
		result := ""
		seenBefore := map[string]struct{}{}
		for start := 0; start <= N-l; start++ {
			subS := S[start : start+l]
			if _, seen := seenBefore[subS]; seen {
				ok = true
				result = subS
				break
			}
			seenBefore[subS] = struct{}{}
		}
		if ok {
			if low == high-1 {
				return result
			}
			low = l
		} else {
			high = l
		}
	}
	return ""
}

func LongestDupSubstring(S string) string {
	return longestDupSubstring(S)
}

func getPermutation(n int, k int) string {
	k--
	s := 1
	for i := 1; i < n; i++ {
		s *= i
	}
	result := ""
	digits := "123456789"
	digits = digits[:n]
	for {
		n--
		i, nextK := k/s, k%s
		result += digits[i : i+1]
		digits = digits[:i] + digits[i+1:]
		if n == 0 {
			break
		}
		s /= n
		k = nextK
	}
	return result
}

func GetPermutation(n int, k int) string {
	return getPermutation(n, k)
}
