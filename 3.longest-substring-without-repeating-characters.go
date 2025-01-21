package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
type LastIndex = int

func lengthOfLongestSubstring(s string) int {
	start := 0
	longest := 0

	used := map[byte]LastIndex{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		// has been used and current is less than used index
		if _, ok := used[c]; ok && start <= used[c] {
			start = used[c] + 1
		}
		used[c] = i

		longest = max(longest, i-start+1)
	}
	return longest
}

// @lc code=end
