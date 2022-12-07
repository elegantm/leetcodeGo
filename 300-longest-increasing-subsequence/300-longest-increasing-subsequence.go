package leetcode

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	res := 1
	dp := make([]int, length)
	for k := range dp {
		dp[k] = 1
	}

	for k := range dp {
		for j := 0; j < k; j++ {
			if nums[k] > nums[j] {
				dp[k] = max(dp[k], dp[j]+1)
			}
		}
		res = max(res, dp[k])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
