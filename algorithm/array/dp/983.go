package dp

// 983. 最低票价
import "math"
// ans = 0
// dp[i] 是当天需要花费的金额
// 如果dp[i]买了第j种车票, 则dp[i]
// func mincostTickets(days []int, costs []int) int {
//     data := [366]int{}
//     dayM := make(map[int]bool)

//     // 标记365天哪天要进行旅行
//     for _, day := range days {
//         dayM[day] = true
//     }

//     var dp func(day int) int
//     dp = func(day int) int {
//         if day > 365 {
//             return 0
//         }

//         if data[day] > 0 {
//             return data[day]
//         }

//         if dayM[day] {
//             data[day] = min(min(dp(day + 1) + costs[0], dp(day + 7) + costs[1]), dp(day + 30) + costs[2])
//         } else {
//             data[day] = dp(day + 1)
//         }
//         return data[day]
//     }

//     return dp(1)
// }

func mincostTickets(days []int, costs []int) int {
	// data 表示在第i天时的最小花费
	data := [366]int{}
	durations := [3]int{1, 7, 30}

	var dp func(idx int) int
	dp = func(idx int) int {
		if idx >= len(days) {
			return 0
		}

		if data[idx] > 0 {
			return data[idx]
		}

		data[idx] = math.MaxInt32
		j := idx
		// 计算3种在当前时间下,3种方式的最小值
		for i:=0;i<3;i++ {
			// 在days数组i, i + 1之间的数据,如果不存在与days数组当中,
			// 则data[i] = data[i+1]
			// 直到days[i + 1] >= days[i] + durations[i]
			for ;j < len(days) && days[j] < days[idx] + durations[i];j++ {}
			data[idx] = min(data[idx], dp(j) + costs[i])
		}

		return data[idx]
	}

	return dp(0)
}


//func mincostTickets(days []int, costs []int) int {
//	durations := [3]int{1, 7, 30}
//	dp := [366]int{}
//	var start int
//	i := 1
//	for _, day := range days {
//		for ; i < day; i++ {
//			dp[i] = dp[i-1]
//		}
//
//		dp[i] = math.MaxInt32
//		// 计算3中通行证售价的最低消费
//		for j := 0; j < 3; j++ {
// 			假设在第i天
//			start = i - durations[j]
//			if start < 0 {
//				start = 0
//			}
//			dp[i] = min(dp[i], dp[start]+costs[j])
//		}
//
//		i++
//	}
//
//	return dp[days[len(days)-1]]
//}

func MincostTickets3(days []int, costs []int) int {
	durations := [3]int{1, 7, 30}
	data := [366]int{}

	var dp func(idx int) int
	dp = func(idx int) int {
		if idx >= len(days) {
			return 0
		}
		if data[idx] > 0 {
			return data[idx]
		}
		data[idx] = math.MaxInt32
		j := idx
		for i := 0; i < 3; i++ {
			for ; j < len(days) && days[j] < days[idx]+durations[i]; j++ {
			}
			data[idx] = min(data[idx], dp(j)+costs[i])
		}
		return data[idx]
	}

	return dp(0)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}