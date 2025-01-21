package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
type Index = int
type Value = int

func twoSum(nums []int, target int) []int {

	store := map[Value]Index{}

	for i, n := range nums {
		if j, ok := store[n]; ok {
			return []int{i, j}
		}
		if _, ok := store[target-n]; !ok {
			store[target-n] = i
		}
	}
	return []int{}
}

// @lc code=end
