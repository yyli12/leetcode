package algorithm

func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 1 {
		return true
	}
	pres := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		this, pre := prerequisite[0], prerequisite[1]
		pres[this] = append(pres[this], pre)
	}

	taken := make([]bool, numCourses)
	inCircle := make([]bool, numCourses)
	var dfs func(start int) bool

	dfs = func(start int) bool {
		taken[start] = true
		inCircle[start] = true
		for _, pre := range pres[start] {
			if !taken[pre] {
				if !dfs(pre) {
					return false
				}
			} else {
				if inCircle[pre] {
					return false
				}
			}
		}
		inCircle[start] = false
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !taken[i] && !dfs(i) {
			return false
		}
	}
	return true
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if src == dst {
		return 0
	}
	if len(flights) == 0 {
		return -1
	}
	flightTable := make([][][2]int, n)
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		flightTable[from] = append(flightTable[from], [2]int{to, price})
	}

	K++
	result := -1
	nextStops := map[int]int{src: 0}
	cheapestPricesTo := map[int]int{src: 0}
	for K >= 0 {
		nextNextStops := map[int]int{}
		for thisStop, priceSoFar := range nextStops {
			if thisStop == dst {
				if result == -1 || priceSoFar < result {
					result = priceSoFar
				}
			}
			for _, flight := range flightTable[thisStop] {
				newPrice := priceSoFar + flight[1]
				if cheapestPrice, exist := cheapestPricesTo[flight[0]]; !exist || newPrice < cheapestPrice {
					nextNextStops[flight[0]] = newPrice
					cheapestPricesTo[flight[0]] = newPrice
				}
			}
		}
		nextStops = nextNextStops
		K--
	}
	return result
}

func FindCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	return findCheapestPrice(n, flights, src, dst, K)
}
