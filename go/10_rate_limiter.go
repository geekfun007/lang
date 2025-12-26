package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 示例1: 令牌桶算法（Token Bucket）
func example1_token_bucket() {
	fmt.Println("\n=== 示例1: 令牌桶限流器 ===")
	
	type TokenBucket struct {
		capacity   int
		tokens     chan struct{}
		refillRate time.Duration
	}
	
	NewTokenBucket := func(capacity int, refillRate time.Duration) *TokenBucket {
		tb := &TokenBucket{
			capacity:   capacity,
			tokens:     make(chan struct{}, capacity),
			refillRate: refillRate,
		}
		
		// 初始填满令牌
		for i := 0; i < capacity; i++ {
			tb.tokens <- struct{}{}
		}
		
		// 定期补充令牌
		go func() {
			ticker := time.NewTicker(refillRate)
			defer ticker.Stop()
			for range ticker.C {
				select {
				case tb.tokens <- struct{}{}:
					// 成功添加令牌
				default:
					// 桶已满
				}
			}
		}()
		
		return tb
	}
	
	Allow := func(tb *TokenBucket) bool {
		select {
		case <-tb.tokens:
			return true
		default:
			return false
		}
	}
	
	// 创建限流器：容量5，每200ms补充1个令牌（5 QPS）
	limiter := NewTokenBucket(5, 200*time.Millisecond)
	
	// 模拟请求
	for i := 1; i <= 15; i++ {
		if Allow(limiter) {
			fmt.Printf("[%s] ✅ 请求 %d 通过\n", time.Now().Format("15:04:05.000"), i)
		} else {
			fmt.Printf("[%s] ❌ 请求 %d 被限流\n", time.Now().Format("15:04:05.000"), i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// 示例2: 漏桶算法（Leaky Bucket）
func example2_leaky_bucket() {
	fmt.Println("\n=== 示例2: 漏桶限流器 ===")
	
	type LeakyBucket struct {
		capacity   int
		rate       time.Duration
		requests   chan struct{}
		processing chan struct{}
	}
	
	NewLeakyBucket := func(capacity int, rate time.Duration) *LeakyBucket {
		lb := &LeakyBucket{
			capacity:   capacity,
			rate:       rate,
			requests:   make(chan struct{}, capacity),
			processing: make(chan struct{}),
		}
		
		// 启动处理器（以固定速率处理）
		go func() {
			ticker := time.NewTicker(rate)
			defer ticker.Stop()
			for range ticker.C {
				select {
				case <-lb.requests:
					lb.processing <- struct{}{}
				default:
				}
			}
		}()
		
		return lb
	}
	
	TryAdd := func(lb *LeakyBucket) bool {
		select {
		case lb.requests <- struct{}{}:
			return true
		default:
			return false // 桶满
		}
	}
	
	Wait := func(lb *LeakyBucket) {
		<-lb.processing
	}
	
	// 创建漏桶：容量10，每300ms处理一个请求
	bucket := NewLeakyBucket(10, 300*time.Millisecond)
	
	var wg sync.WaitGroup
	
	// 快速发送15个请求
	for i := 1; i <= 15; i++ {
		if TryAdd(bucket) {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				fmt.Printf("请求 %d 进入桶，等待处理...\n", id)
				Wait(bucket)
				fmt.Printf("[%s] ✅ 请求 %d 处理完成\n", time.Now().Format("15:04:05.000"), id)
			}(i)
		} else {
			fmt.Printf("❌ 请求 %d 被拒绝（桶满）\n", i)
		}
		time.Sleep(50 * time.Millisecond)
	}
	
	wg.Wait()
}

// 示例3: 固定窗口计数器
func example3_fixed_window_counter() {
	fmt.Println("\n=== 示例3: 固定窗口计数器 ===")
	
	type FixedWindowCounter struct {
		limit      int
		window     time.Duration
		count      int
		windowEnd  time.Time
		mu         sync.Mutex
	}
	
	NewFixedWindowCounter := func(limit int, window time.Duration) *FixedWindowCounter {
		return &FixedWindowCounter{
			limit:     limit,
			window:    window,
			count:     0,
			windowEnd: time.Now().Add(window),
		}
	}
	
	Allow := func(fwc *FixedWindowCounter) bool {
		fwc.mu.Lock()
		defer fwc.mu.Unlock()
		
		now := time.Now()
		
		// 检查是否需要重置窗口
		if now.After(fwc.windowEnd) {
			fwc.count = 0
			fwc.windowEnd = now.Add(fwc.window)
			fmt.Printf("[%s] 🔄 窗口重置\n", now.Format("15:04:05.000"))
		}
		
		// 检查是否超过限制
		if fwc.count >= fwc.limit {
			return false
		}
		
		fwc.count++
		return true
	}
	
	// 创建限流器：每秒最多5个请求
	limiter := NewFixedWindowCounter(5, 1*time.Second)
	
	// 模拟请求
	for i := 1; i <= 20; i++ {
		if Allow(limiter) {
			fmt.Printf("[%s] ✅ 请求 %d 通过\n", time.Now().Format("15:04:05.000"), i)
		} else {
			fmt.Printf("[%s] ❌ 请求 %d 被限流\n", time.Now().Format("15:04:05.000"), i)
		}
		time.Sleep(150 * time.Millisecond)
	}
}

// 示例4: 滑动窗口计数器
func example4_sliding_window_counter() {
	fmt.Println("\n=== 示例4: 滑动窗口计数器 ===")
	
	type SlidingWindowCounter struct {
		limit     int
		window    time.Duration
		requests  []time.Time
		mu        sync.Mutex
	}
	
	NewSlidingWindowCounter := func(limit int, window time.Duration) *SlidingWindowCounter {
		return &SlidingWindowCounter{
			limit:    limit,
			window:   window,
			requests: make([]time.Time, 0),
		}
	}
	
	Allow := func(swc *SlidingWindowCounter) bool {
		swc.mu.Lock()
		defer swc.mu.Unlock()
		
		now := time.Now()
		cutoff := now.Add(-swc.window)
		
		// 移除过期请求
		validRequests := make([]time.Time, 0)
		for _, t := range swc.requests {
			if t.After(cutoff) {
				validRequests = append(validRequests, t)
			}
		}
		swc.requests = validRequests
		
		// 检查是否超过限制
		if len(swc.requests) >= swc.limit {
			return false
		}
		
		swc.requests = append(swc.requests, now)
		return true
	}
	
	GetCount := func(swc *SlidingWindowCounter) int {
		swc.mu.Lock()
		defer swc.mu.Unlock()
		return len(swc.requests)
	}
	
	// 创建限流器：1秒窗口内最多5个请求
	limiter := NewSlidingWindowCounter(5, 1*time.Second)
	
	for i := 1; i <= 15; i++ {
		if Allow(limiter) {
			fmt.Printf("[%s] ✅ 请求 %d 通过 (当前窗口: %d/5)\n", 
				time.Now().Format("15:04:05.000"), i, GetCount(limiter))
		} else {
			fmt.Printf("[%s] ❌ 请求 %d 被限流\n", time.Now().Format("15:04:05.000"), i)
		}
		time.Sleep(150 * time.Millisecond)
	}
}

// 示例5: 并发限流器（信号量）
func example5_concurrency_limiter() {
	fmt.Println("\n=== 示例5: 并发限流器 ===")
	
	type ConcurrencyLimiter struct {
		semaphore chan struct{}
	}
	
	NewConcurrencyLimiter := func(maxConcurrent int) *ConcurrencyLimiter {
		return &ConcurrencyLimiter{
			semaphore: make(chan struct{}, maxConcurrent),
		}
	}
	
	Acquire := func(cl *ConcurrencyLimiter, ctx context.Context) error {
		select {
		case cl.semaphore <- struct{}{}:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	
	Release := func(cl *ConcurrencyLimiter) {
		<-cl.semaphore
	}
	
	// 创建限流器：最多3个并发
	limiter := NewConcurrencyLimiter(3)
	
	var wg sync.WaitGroup
	
	// 启动10个任务
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			ctx := context.Background()
			if err := Acquire(limiter, ctx); err != nil {
				fmt.Printf("任务 %d 获取许可失败: %v\n", id, err)
				return
			}
			defer Release(limiter)
			
			fmt.Printf("[%s] 任务 %d 开始执行\n", time.Now().Format("15:04:05"), id)
			time.Sleep(1 * time.Second)
			fmt.Printf("[%s] 任务 %d 完成\n", time.Now().Format("15:04:05"), id)
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	
	wg.Wait()
}

// 示例6: 自适应限流器
func example6_adaptive_limiter() {
	fmt.Println("\n=== 示例6: 自适应限流器 ===")
	
	type AdaptiveLimiter struct {
		minLimit     int
		maxLimit     int
		currentLimit int
		successCount int
		failureCount int
		mu           sync.Mutex
		semaphore    chan struct{}
	}
	
	// 定义 adjust 函数
	var adjust func(*AdaptiveLimiter)
	adjust = func(al *AdaptiveLimiter) {
		al.mu.Lock()
		defer al.mu.Unlock()
		
		total := al.successCount + al.failureCount
		if total == 0 {
			return
		}
		
		successRate := float64(al.successCount) / float64(total)
		
		if successRate > 0.9 && al.currentLimit < al.maxLimit {
			al.currentLimit++
			fmt.Printf("📈 增加限流: %d\n", al.currentLimit)
		} else if successRate < 0.5 && al.currentLimit > al.minLimit {
			al.currentLimit--
			fmt.Printf("📉 降低限流: %d\n", al.currentLimit)
		}
		
		al.successCount = 0
		al.failureCount = 0
	}
	
	NewAdaptiveLimiter := func(minLimit, maxLimit int) *AdaptiveLimiter {
		al := &AdaptiveLimiter{
			minLimit:     minLimit,
			maxLimit:     maxLimit,
			currentLimit: minLimit,
			semaphore:    make(chan struct{}, maxLimit),
		}
		
		// 定期调整限流
		go func() {
			ticker := time.NewTicker(2 * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				adjust(al)
			}
		}()
		
		return al
	}
	
	TryAcquire := func(al *AdaptiveLimiter) bool {
		al.mu.Lock()
		currentCount := len(al.semaphore)
		limit := al.currentLimit
		al.mu.Unlock()
		
		if currentCount >= limit {
			return false
		}
		
		select {
		case al.semaphore <- struct{}{}:
			return true
		default:
			return false
		}
	}
	
	Release := func(al *AdaptiveLimiter, success bool) {
		<-al.semaphore
		
		al.mu.Lock()
		if success {
			al.successCount++
		} else {
			al.failureCount++
		}
		al.mu.Unlock()
	}
	
	// 创建自适应限流器
	limiter := NewAdaptiveLimiter(2, 10)
	
	var wg sync.WaitGroup
	
	// 模拟请求
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			if !TryAcquire(limiter) {
				fmt.Printf("❌ 请求 %d 被限流\n", id)
				return
			}
			
			// 模拟处理
			time.Sleep(200 * time.Millisecond)
			
			// 模拟成功/失败
			success := id%5 != 0 // 80%成功率
			
			Release(limiter, success)
			
			if success {
				fmt.Printf("✅ 请求 %d 成功\n", id)
			} else {
				fmt.Printf("⚠️  请求 %d 失败\n", id)
			}
		}(i)
		time.Sleep(50 * time.Millisecond)
	}
	
	wg.Wait()
}

// 示例7: 分布式限流器（模拟）
func example7_distributed_limiter() {
	fmt.Println("\n=== 示例7: 分布式限流器（模拟）===")
	
	type DistributedLimiter struct {
		nodeID    string
		limit     int
		window    time.Duration
		localCount map[string]int
		mu         sync.Mutex
	}
	
	NewDistributedLimiter := func(nodeID string, limit int, window time.Duration) *DistributedLimiter {
		dl := &DistributedLimiter{
			nodeID:     nodeID,
			limit:      limit,
			window:     window,
			localCount: make(map[string]int),
		}
		
		// 定期清理过期数据
		go func() {
			ticker := time.NewTicker(window)
			defer ticker.Stop()
			for range ticker.C {
				dl.mu.Lock()
				dl.localCount = make(map[string]int)
				dl.mu.Unlock()
			}
		}()
		
		return dl
	}
	
	Allow := func(dl *DistributedLimiter, userID string) bool {
		dl.mu.Lock()
		defer dl.mu.Unlock()
		
		count := dl.localCount[userID]
		if count >= dl.limit {
			return false
		}
		
		dl.localCount[userID]++
		return true
	}
	
	// 模拟3个节点
	nodes := []*DistributedLimiter{
		NewDistributedLimiter("node-1", 5, 2*time.Second),
		NewDistributedLimiter("node-2", 5, 2*time.Second),
		NewDistributedLimiter("node-3", 5, 2*time.Second),
	}
	
	var wg sync.WaitGroup
	
	// 模拟用户请求分散到不同节点
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(reqID int) {
			defer wg.Done()
			
			userID := fmt.Sprintf("user-%d", (reqID%3)+1)
			node := nodes[reqID%3]
			
			if Allow(node, userID) {
				fmt.Printf("[%s] ✅ 请求 %d (%s on %s) 通过\n", 
					time.Now().Format("15:04:05"), reqID, userID, node.nodeID)
			} else {
				fmt.Printf("[%s] ❌ 请求 %d (%s on %s) 被限流\n", 
					time.Now().Format("15:04:05"), reqID, userID, node.nodeID)
			}
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	
	wg.Wait()
}

func main() {
	fmt.Println("======================================")
	fmt.Println("    速率限制器实战")
	fmt.Println("======================================")
	
	example1_token_bucket()
	time.Sleep(1 * time.Second)
	
	example2_leaky_bucket()
	time.Sleep(1 * time.Second)
	
	example3_fixed_window_counter()
	time.Sleep(1 * time.Second)
	
	example4_sliding_window_counter()
	time.Sleep(1 * time.Second)
	
	example5_concurrency_limiter()
	time.Sleep(1 * time.Second)
	
	example6_adaptive_limiter()
	time.Sleep(1 * time.Second)
	
	example7_distributed_limiter()
	
	fmt.Println("\n所有示例运行完成！")
}
