package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 示例1: 基本的扇出-扇入模式
func example1_basic_fan_out_fan_in() {
	fmt.Println("\n=== 示例1: 基本扇出-扇入 ===")
	
	// 数据源
	source := func() <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 1; i <= 10; i++ {
				out <- i
			}
		}()
		return out
	}
	
	// 工作函数（耗时操作）
	worker := func(id int, in <-chan int) <-chan string {
		out := make(chan string)
		go func() {
			defer close(out)
			for n := range in {
				time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
				result := fmt.Sprintf("Worker %d: %d -> %d", id, n, n*n)
				out <- result
			}
		}()
		return out
	}
	
	// 合并多个 channel
	merge := func(channels ...<-chan string) <-chan string {
		var wg sync.WaitGroup
		out := make(chan string)
		
		// 为每个输入 channel 启动一个 goroutine
		output := func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}
		
		wg.Add(len(channels))
		for _, c := range channels {
			go output(c)
		}
		
		// 等待所有 goroutine 完成后关闭输出 channel
		go func() {
			wg.Wait()
			close(out)
		}()
		
		return out
	}
	
	// 构建流水线
	input := source()
	
	// 扇出：启动3个 worker
	c1 := worker(1, input)
	c2 := worker(2, input)
	c3 := worker(3, input)
	
	// 扇入：合并结果
	results := merge(c1, c2, c3)
	
	// 消费结果
	for result := range results {
		fmt.Println(result)
	}
}

// 示例2: 扇出到多个独立的处理器
func example2_fan_out_to_multiple_processors() {
	fmt.Println("\n=== 示例2: 扇出到多个处理器 ===")
	
	// 数据源
	source := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}
	
	// 复制 channel 的值到多个目标
	fanOut := func(in <-chan int, n int) []<-chan int {
		// 创建可写的 channel
		writableOuts := make([]chan int, n)
		readableOuts := make([]<-chan int, n)
		
		for i := 0; i < n; i++ {
			ch := make(chan int)
			writableOuts[i] = ch
			readableOuts[i] = ch
		}
		
		go func() {
			defer func() {
				for _, out := range writableOuts {
					close(out)
				}
			}()
			
			for v := range in {
				// 发送到所有输出 channel
				for _, out := range writableOuts {
					out <- v
				}
			}
		}()
		
		return readableOuts
	}
	
	// 构建流水线
	input := source(1, 2, 3, 4, 5)
	channels := fanOut(input, 3)
	
	var wg sync.WaitGroup
	for i, ch := range channels {
		wg.Add(1)
		go func(id int, c <-chan int) {
			defer wg.Done()
			for v := range c {
				fmt.Printf("处理器 %d 收到: %d\n", id, v)
			}
		}(i+1, ch)
	}
	
	wg.Wait()
}

// 示例3: 带负载均衡的扇出
func example3_load_balanced_fan_out() {
	fmt.Println("\n=== 示例3: 负载均衡扇出 ===")
	
	type Task struct {
		ID   int
		Data int
	}
	
	// 任务生成器
	taskGenerator := func(n int) <-chan Task {
		out := make(chan Task)
		go func() {
			defer close(out)
			for i := 1; i <= n; i++ {
				out <- Task{ID: i, Data: i * 10}
			}
		}()
		return out
	}
	
	// Worker
	worker := func(id int, tasks <-chan Task) <-chan string {
		out := make(chan string)
		go func() {
			defer close(out)
			count := 0
			for task := range tasks {
				count++
				// 模拟处理
				time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
				result := fmt.Sprintf("Worker %d 完成任务 %d (第%d个任务)", id, task.ID, count)
				out <- result
			}
			fmt.Printf("Worker %d 总共处理了 %d 个任务\n", id, count)
		}()
		return out
	}
	
	// 合并结果
	merge := func(channels ...<-chan string) <-chan string {
		var wg sync.WaitGroup
		out := make(chan string)
		
		wg.Add(len(channels))
		for _, c := range channels {
			go func(ch <-chan string) {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}(c)
		}
		
		go func() {
			wg.Wait()
			close(out)
		}()
		
		return out
	}
	
	// 构建流水线
	tasks := taskGenerator(20)
	
	// 扇出到5个 worker（负载自动均衡）
	workers := make([]<-chan string, 5)
	for i := 0; i < 5; i++ {
		workers[i] = worker(i+1, tasks)
	}
	
	// 扇入结果
	results := merge(workers...)
	
	// 消费结果
	for result := range results {
		fmt.Println(result)
	}
}

// 示例4: 多阶段扇出-扇入
func example4_multi_stage_fan_out_fan_in() {
	fmt.Println("\n=== 示例4: 多阶段扇出-扇入 ===")
	
	// 第一阶段：数据生成
	generate := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}
	
	// 第二阶段：并行计算平方
	square := func(id int, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				result := n * n
				fmt.Printf("Square Worker %d: %d -> %d\n", id, n, result)
				out <- result
				time.Sleep(100 * time.Millisecond)
			}
		}()
		return out
	}
	
	// 第三阶段：并行加倍
	double := func(id int, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				result := n * 2
				fmt.Printf("Double Worker %d: %d -> %d\n", id, n, result)
				out <- result
				time.Sleep(50 * time.Millisecond)
			}
		}()
		return out
	}
	
	// 合并函数
	merge := func(channels ...<-chan int) <-chan int {
		var wg sync.WaitGroup
		out := make(chan int)
		
		wg.Add(len(channels))
		for _, c := range channels {
			go func(ch <-chan int) {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}(c)
		}
		
		go func() {
			wg.Wait()
			close(out)
		}()
		
		return out
	}
	
	// 构建多阶段流水线
	input := generate(1, 2, 3, 4, 5, 6, 7, 8)
	
	// 阶段1：扇出到3个平方计算 worker
	s1 := square(1, input)
	s2 := square(2, input)
	s3 := square(3, input)
	
	// 扇入
	stage1Output := merge(s1, s2, s3)
	
	// 阶段2：扇出到2个加倍 worker
	d1 := double(1, stage1Output)
	d2 := double(2, stage1Output)
	
	// 扇入
	finalOutput := merge(d1, d2)
	
	// 收集结果
	fmt.Println("\n最终结果:")
	var results []int
	for result := range finalOutput {
		results = append(results, result)
	}
	
	// 排序显示
	fmt.Println(results)
}

// 示例5: 带优先级的扇入
func example5_priority_fan_in() {
	fmt.Println("\n=== 示例5: 带优先级的扇入 ===")
	
	type Message struct {
		Priority int
		Content  string
	}
	
	// 生成不同优先级的消息
	highPriority := func() <-chan Message {
		out := make(chan Message)
		go func() {
			defer close(out)
			for i := 1; i <= 5; i++ {
				time.Sleep(200 * time.Millisecond)
				out <- Message{Priority: 1, Content: fmt.Sprintf("高优先级-%d", i)}
			}
		}()
		return out
	}
	
	lowPriority := func() <-chan Message {
		out := make(chan Message)
		go func() {
			defer close(out)
			for i := 1; i <= 10; i++ {
				time.Sleep(100 * time.Millisecond)
				out <- Message{Priority: 3, Content: fmt.Sprintf("低优先级-%d", i)}
			}
		}()
		return out
	}
	
	// 优先级合并
	priorityMerge := func(high, low <-chan Message) <-chan Message {
		out := make(chan Message)
		go func() {
			defer close(out)
			for high != nil || low != nil {
				select {
				case msg, ok := <-high:
					if !ok {
						high = nil
					} else {
						out <- msg
					}
				default:
					select {
					case msg, ok := <-high:
						if !ok {
							high = nil
						} else {
							out <- msg
						}
					case msg, ok := <-low:
						if !ok {
							low = nil
						} else {
							out <- msg
						}
					}
				}
			}
		}()
		return out
	}
	
	// 构建流水线
	high := highPriority()
	low := lowPriority()
	merged := priorityMerge(high, low)
	
	// 消费结果
	for msg := range merged {
		fmt.Printf("[优先级 %d] %s\n", msg.Priority, msg.Content)
	}
}

// 示例6: 带错误处理的扇出-扇入
func example6_fan_out_with_error_handling() {
	fmt.Println("\n=== 示例6: 带错误处理的扇出-扇入 ===")
	
	type Result struct {
		Value int
		Err   error
	}
	
	// 任务生成器
	taskGen := func() <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 1; i <= 15; i++ {
				out <- i
			}
		}()
		return out
	}
	
	// Worker（可能失败）
	worker := func(id int, tasks <-chan int) <-chan Result {
		out := make(chan Result)
		go func() {
			defer close(out)
			for task := range tasks {
				time.Sleep(50 * time.Millisecond)
				
				// 模拟随机失败
				if rand.Float32() < 0.3 {
					out <- Result{
						Value: 0,
						Err:   fmt.Errorf("worker %d: 任务 %d 失败", id, task),
					}
				} else {
					out <- Result{
						Value: task * task,
						Err:   nil,
					}
				}
			}
		}()
		return out
	}
	
	// 合并结果
	merge := func(channels ...<-chan Result) <-chan Result {
		var wg sync.WaitGroup
		out := make(chan Result)
		
		wg.Add(len(channels))
		for _, c := range channels {
			go func(ch <-chan Result) {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}(c)
		}
		
		go func() {
			wg.Wait()
			close(out)
		}()
		
		return out
	}
	
	// 构建流水线
	tasks := taskGen()
	
	// 扇出到4个 worker
	workers := make([]<-chan Result, 4)
	for i := 0; i < 4; i++ {
		workers[i] = worker(i+1, tasks)
	}
	
	// 扇入结果
	results := merge(workers...)
	
	// 收集统计
	successCount := 0
	errorCount := 0
	
	for result := range results {
		if result.Err != nil {
			fmt.Printf("❌ %v\n", result.Err)
			errorCount++
		} else {
			fmt.Printf("✅ 结果: %d\n", result.Value)
			successCount++
		}
	}
	
	fmt.Printf("\n统计: 成功 %d, 失败 %d\n", successCount, errorCount)
}

// 示例7: 动态扇出（根据负载调整 worker 数量）
func example7_dynamic_fan_out() {
	fmt.Println("\n=== 示例7: 动态扇出 ===")
	
	type WorkerPool struct {
		tasks   chan int
		results chan int
		workers int
		wg      sync.WaitGroup
	}
	
	// 定义 addWorker 函数
	var addWorker func(*WorkerPool)
	addWorker = func(wp *WorkerPool) {
		wp.workers++
		wp.wg.Add(1)
		id := wp.workers
		
		go func() {
			defer wp.wg.Done()
			fmt.Printf("Worker %d 启动\n", id)
			
			for task := range wp.tasks {
				time.Sleep(100 * time.Millisecond)
				wp.results <- task * task
			}
			
			fmt.Printf("Worker %d 停止\n", id)
		}()
	}
	
	NewWorkerPool := func(initialWorkers int) *WorkerPool {
		wp := &WorkerPool{
			tasks:   make(chan int, 100),
			results: make(chan int, 100),
			workers: 0,
		}
		
		// 启动初始 worker
		for i := 0; i < initialWorkers; i++ {
			addWorker(wp)
		}
		
		return wp
	}
	
	wp := NewWorkerPool(2)
	
	// 提交任务
	go func() {
		for i := 1; i <= 20; i++ {
			wp.tasks <- i
			
			// 根据队列长度动态添加 worker
			if len(wp.tasks) > 10 && wp.workers < 5 {
				fmt.Printf("队列长度 %d，添加新 worker\n", len(wp.tasks))
				addWorker(wp)
			}
		}
		close(wp.tasks)
	}()
	
	// 收集结果
	go func() {
		wp.wg.Wait()
		close(wp.results)
	}()
	
	for result := range wp.results {
		fmt.Printf("结果: %d\n", result)
	}
	
	fmt.Printf("最终 worker 数量: %d\n", wp.workers)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("======================================")
	fmt.Println("    Go 扇出-扇入模式示例")
	fmt.Println("======================================")
	
	example1_basic_fan_out_fan_in()
	example2_fan_out_to_multiple_processors()
	example3_load_balanced_fan_out()
	example4_multi_stage_fan_out_fan_in()
	example5_priority_fan_in()
	example6_fan_out_with_error_handling()
	example7_dynamic_fan_out()
	
	fmt.Println("\n所有示例运行完成！")
}
