package main

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	// imin and imax are the boundaries of the search space

	imin, imax, half := 0, m, (m+n+1)/2

	var i, j int
	for imin <= imax {
		i = (imin + imax) / 2
		// i + j is total length of left half
		j = half - i

		if i < m && nums2[j-1] > nums1[i] {
			// num2[:j] has greater elements than num1[i:]. imin is too small
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// num1[:i] has greater elements than num2[j:]. imax is too big
			imax = i - 1
		} else {
			// i is perfect
			break
		}
	}

	var maxLeft int
	if i == 0 {
		maxLeft = nums2[j-1]
	} else if j == 0 {
		maxLeft = nums1[i-1]
	} else {
		maxLeft = max(nums1[i-1], nums2[j-1])
	}

	// if odd, median is maxLeft
	if (m+n)%2 == 1 {
		return float64(maxLeft)
	}

	var minRight int
	if i == m {
		minRight = nums2[j]
	} else if j == n {
		minRight = nums1[i]
	} else {
		minRight = min(nums1[i], nums2[j])
	}

	// if even, median is average of maxLeft and minRight
	return float64(maxLeft+minRight) / 2
}

// @lc code=end
