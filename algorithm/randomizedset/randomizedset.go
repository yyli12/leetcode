package randomizedset

import "math/rand"

type RandomizedSet struct {
	size     int
	store    []int
	indexMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		size:     0,
		store:    make([]int, 5),
		indexMap: map[int]int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (s *RandomizedSet) Insert(val int) bool {
	if _, exist := s.indexMap[val]; exist {
		return false
	}

	if s.size == len(s.store) {
		// extend the store
		s.store = append(s.store, val)
		s.store = s.store[:cap(s.store)]
	} else {
		s.store[s.size] = val
	}
	s.indexMap[val] = s.size
	s.size++
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (s *RandomizedSet) Remove(val int) bool {
	if oldIndex, exist := s.indexMap[val]; exist {
		if oldIndex != s.size-1 {
			s.store[oldIndex] = s.store[s.size-1]
			s.indexMap[s.store[oldIndex]] = oldIndex
		}
		s.size--
		delete(s.indexMap, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (s *RandomizedSet) GetRandom() int {
	if s.size > 0 {
		return s.store[rand.Intn(s.size)]
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
