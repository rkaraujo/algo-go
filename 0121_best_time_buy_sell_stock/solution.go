package solution0121

func maxProfit(prices []int) int {
	max_profit := 0

	prev := prices[0]

	cur_profit := 0
	for i := 1; i < len(prices); i++ {
		cur := prices[i]
		if cur > prev {
			cur_profit += cur - prev
		} else {
			cur_profit -= prev - cur
			if cur_profit < 0 {
				cur_profit = 0
			}
		}

		if cur_profit > max_profit {
			max_profit = cur_profit
		}

		prev = cur
	}

	return max_profit
}
