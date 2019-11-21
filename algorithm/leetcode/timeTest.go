package leetcode

import (
	"fmt"
	"time"
)

func CalcDays() int {
	now := time.Now()
	n := time.Date(2019, 10, 20, now.Hour(), now.Minute(), 0, 0, now.Location())
	after := time.Date(2019, 10, 20, 0, 0, 0 ,0, now.Location())
	hours := n.Sub(after).Hours()
	fmt.Println(now.AddDate(0, 0, -13), now.AddDate(0, 0, -26))
	fmt.Println(int(hours))
	return int(hours)
}