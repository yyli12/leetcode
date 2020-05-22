package main

import (
	"fmt"
	"github.com/leetcode/algorithm/dp"
)

func main() {
	fmt.Println(dp.CountSquares([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}))
}
