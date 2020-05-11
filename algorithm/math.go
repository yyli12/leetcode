package algorithm

func CheckStraightLine(coordinates [][]int) bool {

	p0, p1 := coordinates[0], coordinates[1]
	dx, dy := p1[0]-p0[0], p1[1]-p0[1]

	for _, p := range coordinates[2:] {
		if (p[0]-p1[0])*dy != (p[1]-p1[1])*dx {
			return false
		}
	}

	return true
}

func CheckStraightLine2(coordinates [][]int) bool {

	a := coordinates[0][1] - coordinates[1][1]
	b := coordinates[1][0] - coordinates[0][0]
	c := coordinates[0][1]*coordinates[1][0] - coordinates[1][1]*coordinates[0][0]
	for _, p := range coordinates[2:] {
		if a*p[0]+b*p[1] != c {
			return false
		}
	}

	return true
}
