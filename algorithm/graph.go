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
