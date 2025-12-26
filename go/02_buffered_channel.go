package main

import (
	"fmt"
	"time"
)

// 示例1: 缓冲 Channel 的基本特性
func example1_basic_buffered() {
	fmt.Println("\n=== 示例1: 缓冲 Channel 基础 ===")
	
	// 创建缓冲大小为3的 channel
	ch := make(chan int, 3)
	
	// 发送3个值不会阻塞
	fmt.Println("发送3个值到缓冲 channel...")
	ch <- 1
	fmt.Println("  发送: 1")
	ch <- 2
	fmt.Println("  发送: 2")
	ch <- 3
	fmt.Println("  发送: 3")
	fmt.Println("所有值已发送，没有阻塞！")
	
	// 接收值
	fmt.Println("\n接收值:")
	fmt.Println("  接收:", <-ch)
	fmt.Println("  接收:", <-ch)
	fmt.Println("  接收:", <-ch)
}

// 示例2: 无缓冲 vs 缓冲 Channel
func example2_unbuffered_vs_buffered() {
	fmt.Println("\n=== 示例2: 无缓冲 vs 缓冲 Channel ===")
	
	// 无缓冲 channel
	fmt.Println("无缓冲 Channel:")
	unbuffered := make(chan int)
	go func() {
		fmt.Println("  Goroutine: 发送到无缓冲 channel")
		unbuffered <- 1 // 会阻塞直到被接收
		fmt.Println("  Goroutine: 发送完成（说明已被接收）")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("  Main: 接收值")
	<-unbuffered
	time.Sleep(500 * time.Millisecond)
	
	// 缓冲 channel
	fmt.Println("\n缓冲 Channel:")
	buffered := make(chan int, 1)
	go func() {
		fmt.Println("  Goroutine: 发送到缓冲 channel")
		buffered <- 1 // 不会阻塞（缓冲区有空间）
		fmt.Println("  Goroutine: 发送完成（立即返回）")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("  Main: 接收值")
	<-buffered
}

// 示例3: 缓冲区满时的阻塞行为
func example3_buffer_full() {
	fmt.Println("\n=== 示例3: 缓冲区满时阻塞 ===")
	
	ch := make(chan int, 2)
	
	ch <- 1
	ch <- 2
	fmt.Println("缓冲区已满（2/2）")
	
	// 尝试发送第3个值（会阻塞）
	go func() {
		fmt.Println("Goroutine: 尝试发送第3个值（会阻塞）...")
		ch <- 3
		fmt.Println("Goroutine: 第3个值发送成功（缓冲区有空间了）")
	}()
	
	time.Sleep(1 * time.Second)
	fmt.Println("Main: 接收一个值，释放缓冲区空间")
	<-ch
	
	time.Sleep(1 * time.Second)
	fmt.Println("Main: 接收剩余值")
	fmt.Println("  ", <-ch)
	fmt.Println("  ", <-ch)
}

// 示例4: 使用缓冲 Channel 解耦生产者和消费者
func example4_producer_consumer() {
	fmt.Println("\n=== 示例4: 生产者-消费者模式 ===")
	
	// 缓冲区允许生产者和消费者以不同速度运行
	ch := make(chan int, 10)
	
	// 快速生产者
	go func() {
		for i := 1; i <= 20; i++ {
			ch <- i
			fmt.Printf("生产: %d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
		close(ch)
	}()
	
	// 慢速消费者
	time.Sleep(500 * time.Millisecond) // 延迟启动
	for val := range ch {
		fmt.Printf("  消费: %d\n", val)
		time.Sleep(150 * time.Millisecond)
	}
}

// 示例5: 获取 Channel 长度和容量
func example5_len_and_cap() {
	fmt.Println("\n=== 示例5: Channel 的 len 和 cap ===")
	
	ch := make(chan int, 5)
	
	fmt.Printf("初始状态: len=%d, cap=%d\n", len(ch), cap(ch))
	
	ch <- 10
	ch <- 20
	ch <- 30
	fmt.Printf("发送3个值后: len=%d, cap=%d\n", len(ch), cap(ch))
	
	<-ch
	fmt.Printf("接收1个值后: len=%d, cap=%d\n", len(ch), cap(ch))
	
	<-ch
	<-ch
	fmt.Printf("接收所有值后: len=%d, cap=%d\n", len(ch), cap(ch))
}

// 示例6: 非阻塞发送和接收
func example6_non_blocking() {
	fmt.Println("\n=== 示例6: 非阻塞操作（使用 select + default）===")
	
	ch := make(chan int, 2)
	
	// 非阻塞发送
	for i := 1; i <= 3; i++ {
		select {
		case ch <- i:
			fmt.Printf("发送成功: %d\n", i)
		default:
			fmt.Printf("缓冲区满，发送失败: %d\n", i)
		}
	}
	
	// 非阻塞接收
	for i := 1; i <= 3; i++ {
		select {
		case val := <-ch:
			fmt.Printf("接收成功: %d\n", val)
		default:
			fmt.Println("Channel 为空，接收失败")
		}
	}
}

// 示例7: 选择合适的缓冲区大小
func example7_buffer_size_tuning() {
	fmt.Println("\n=== 示例7: 缓冲区大小的影响 ===")
	
	testWithBufferSize := func(size int) time.Duration {
		ch := make(chan int, size)
		start := time.Now()
		
		// 生产者
		go func() {
			for i := 1; i <= 100; i++ {
				ch <- i
			}
			close(ch)
		}()
		
		// 消费者
		for range ch {
			time.Sleep(1 * time.Millisecond)
		}
		
		return time.Since(start)
	}
	
	sizes := []int{0, 1, 10, 50}
	for _, size := range sizes {
		duration := testWithBufferSize(size)
		fmt.Printf("缓冲区大小 %3d: 耗时 %v\n", size, duration)
	}
}

// 示例8: 缓冲 Channel 的实际应用 - 限制并发数
func example8_concurrency_limit() {
	fmt.Println("\n=== 示例8: 使用缓冲 Channel 限制并发数 ===")
	
	// 最多允许3个并发任务
	semaphore := make(chan struct{}, 3)
	
	// 启动10个任务
	for i := 1; i <= 10; i++ {
		taskID := i
		
		// 获取信号量（如果满了会阻塞）
		semaphore <- struct{}{}
		
		go func() {
			defer func() { <-semaphore }() // 释放信号量
			
			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("任务 %d 完成\n", taskID)
		}()
	}
	
	// 等待所有任务完成
	for i := 0; i < 3; i++ {
		semaphore <- struct{}{}
	}
	fmt.Println("所有任务已完成")
}

func main() {
	fmt.Println("======================================")
	fmt.Println("    Go 缓冲 Channel 示例")
	fmt.Println("======================================")
	
	example1_basic_buffered()
	example2_unbuffered_vs_buffered()
	example3_buffer_full()
	example4_producer_consumer()
	example5_len_and_cap()
	example6_non_blocking()
	example7_buffer_size_tuning()
	example8_concurrency_limit()
	
	fmt.Println("\n所有示例运行完成！")
}
