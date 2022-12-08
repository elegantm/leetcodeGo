package leetcode

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0] // 最小值
	max := 0         // 最大利润
	for i := 1; i < len(prices); i++ {
		if prices[i] < min { // 如果当前值小于最小值，更新最小值
			min = prices[i] // 更新最小值
		}
		if prices[i]-min > max { // 如果当前值减去最小值大于最大利润，更新最大利润
			max = prices[i] - min // 更新最大利润
		}
	}
	return max
}
