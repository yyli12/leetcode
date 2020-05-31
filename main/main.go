package main

import (
	"fmt"
	"github.com/leetcode/algorithm"
)

func main() {
	//fmt.Println(algorithm.FindMaxLength([]int{0,0,0,0}))
	fmt.Println(algorithm.CanFinish(6, [][]int{
		{0, 1},
		{0, 2},
		{2, 3},
		{3, 1},
		{4, 0},
		{5, 4},
	}))
}
