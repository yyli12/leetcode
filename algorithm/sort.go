package algorithm

func MajorityElement(nums []int) int {
	adjust := func(l, r int) (int, int) {
		nl, nr := l-1, r
		pivot := nums[nr-1]
		for i := l; i < nr; {
			if nums[i] < pivot {
				nl++
				nums[i], nums[nl] = nums[nl], nums[i]
				i++
			} else if nums[i] > pivot {
				nr--
				nums[i], nums[nr] = nums[nr], nums[i]
			} else {
				i++
			}
		}
		return nl, nr
	}

	m := len(nums) / 2

	l, r := 0, len(nums)
	for {
		nl, nr := adjust(l, r)
		if nl >= m {
			l, r = l, nl+1
		} else if m > nr {
			l, r = nr, r
		} else {
			break
		}
	}
	return nums[m]
}
