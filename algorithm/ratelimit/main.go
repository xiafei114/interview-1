package main

import (
	"fmt"
	"sync"
	"time"
)

type rateLimit struct {
	begin time.Time
	cycle time.Duration
	limit int // 周期内的请求上限
	count int // 当前周期内的请求数
	lock sync.Mutex
}

func NewRateLimit(n int, cycle time.Duration) *rateLimit{
	return &rateLimit{
		begin: time.Now(),
		cycle: cycle,
		limit: n,
		count: 0,
		lock: sync.Mutex{},
	}
}

func (l *rateLimit) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	for {
		if l.count < l.limit { // 未达到请求上限
			l.count ++
			return true
		}
		// 判断是否还在时间周期内
		now := time.Now()
		if now.Sub(l.begin) > l.cycle { // 超过时间周期, 重新计数
			l.Reset(now)
			return true
		} else {
		}
	}
}

func (l *rateLimit) Reset(resetTime time.Time) {
	fmt.Println("reset ratelimit")
	l.begin = resetTime
	l.count = 0
}

func main() {
	var wg sync.WaitGroup
	rate := NewRateLimit(2, time.Second)

	for i:=0;i<10;i++ {
		wg.Add(1)

		fmt.Println("create req", i, time.Now())
		go func(i int) {
			if rate.Allow() {
				fmt.Println("response req", i, time.Now())
			}

			wg.Done()

			time.Sleep(200 * time.Millisecond)
		}(i)
	}

	wg.Wait()
}