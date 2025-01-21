package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	r := 0

	bs := []byte(s)

	for i := range bs {

		us := map[byte]struct{}{
			bs[i]: {},
		}

		for j := i + 1; j < len(s); j++ {
			if _, ok := us[bs[j]]; ok {
				break
			}
			us[bs[j]] = struct{}{}
		}

		if len(us) > r {
			r = len(us)
		}
	}
	return r
}

// @lc code=end
