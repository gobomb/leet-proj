package leetcode

import (
	"log"
	"sort"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	var is = new(intervalSlice)
	for i := range intervals {
		is.slice = append(is.slice, interval{intervals[i][0], intervals[i][1]})
	}

	sort.Sort(is)
	rs := [][]int{}

	for i := 0; i+1 < is.Len(); i++ {
		for i+1 < is.Len() {
			if is.slice[i+1].start <= is.slice[i].end {
				is.slice[i].end = max(is.slice[i+1].end, is.slice[i].end)
				is.slice = append(is.slice[:i+1], is.slice[i+2:]...)
				continue
			} else {
				rs = append(rs, []int{is.slice[i].start, is.slice[i].end})
				break
			}
		}
	}
	rs = append(rs, []int{is.slice[is.Len()-1].start, is.slice[is.Len()-1].end})

	return rs
}

type interval struct {
	start, end int
}

type intervalSlice struct {
	slice []interval
}

func (is *intervalSlice) Less(i, j int) bool {
	return is.slice[i].start < is.slice[j].start
}

func (is *intervalSlice) Len() int {
	return len(is.slice)
}

func (is *intervalSlice) Swap(i, j int) {
	is.slice[i], is.slice[j] = is.slice[j], is.slice[i]
}
