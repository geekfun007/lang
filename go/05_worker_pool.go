package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task 代表一个工作任务
type Task struct {
	ID     int
	Data   string
	Result chan<- string
}

// 示例1: 简单的 Worker Pool
func example1_simple_worker_pool() {
	fmt.Println("\n=== 示例1: 简单的 Worker Pool ===")
	
	const numWorkers = 3
	const numTasks = 10
	
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)
	
	// 启动 workers
	for w := 1; w <= numWorkers; w++ {
		go func(id int) {
			for task := range tasks {
				fmt.Printf("Worker %d 开始处理任务 %d\n", id, task)
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				result := task * 2
				fmt.Printf("Worker %d 完成任务 %d，结果: %d\n", id, task, result)
				results <- result
			}
		}(w)
	}
	
	// 发送任务
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)
	
	// 收集结果
	for i := 1; i <= numTasks; i++ {
		<-results
	}
}

// 示例2: 使用 WaitGroup 的 Worker Pool
func example2_worker_pool_with_waitgroup() {
	fmt.Println("\n=== 示例2: 使用 WaitGroup 的 Worker Pool ===")
	
	const numWorkers = 4
	tasks := make(chan int, 20)
	var wg sync.WaitGroup
	
	// 启动 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				fmt.Printf("Worker %d: 处理任务 %d\n", id, task)
				time.Sleep(100 * time.Millisecond)
			}
			fmt.Printf("Worker %d: 完成所有任务\n", id)
		}(w)
	}
	
	// 发送任务
	for i := 1; i <= 20; i++ {
		tasks <- i
	}
	close(tasks)
	
	// 等待所有 workers 完成
	wg.Wait()
	fmt.Println("所有任务已完成")
}

// 示例3: Worker Pool 处理结构化任务
type Job struct {
	ID       int
	Priority int
	Data     string
}

type Result struct {
	JobID  int
	Output string
	Error  error
}

func processJob(job Job) Result {
	// 模拟处理
	time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
	return Result{
		JobID:  job.ID,
		Output: fmt.Sprintf("处理结果: %s (优先级: %d)", job.Data, job.Priority),
		Error:  nil,
	}
}

func example3_structured_worker_pool() {
	fmt.Println("\n=== 示例3: 处理结构化任务的 Worker Pool ===")
	
	const numWorkers = 3
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	var wg sync.WaitGroup
	
	// 启动 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				fmt.Printf("Worker %d: 开始处理 Job %d\n", id, job.ID)
				result := processJob(job)
				results <- result
			}
		}(w)
	}
	
	// 发送任务
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- Job{
				ID:       i,
				Priority: rand.Intn(5),
				Data:     fmt.Sprintf("任务数据-%d", i),
			}
		}
		close(jobs)
	}()
	
	// 收集结果
	go func() {
		wg.Wait()
		close(results)
	}()
	
	for result := range results {
		fmt.Printf("收到结果 [Job %d]: %s\n", result.JobID, result.Output)
	}
}

// 示例4: 可动态调整大小的 Worker Pool
type WorkerPool struct {
	tasks      chan Task
	numWorkers int
	wg         sync.WaitGroup
	quit       chan struct{}
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		tasks:      make(chan Task, 100),
		numWorkers: numWorkers,
		quit:       make(chan struct{}),
	}
}

func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case task, ok := <-wp.tasks:
			if !ok {
				fmt.Printf("Worker %d: 任务队列已关闭\n", id)
				return
			}
			wp.processTask(id, task)
		case <-wp.quit:
			fmt.Printf("Worker %d: 收到退出信号\n", id)
			return
		}
	}
}

func (wp *WorkerPool) processTask(workerID int, task Task) {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	result := fmt.Sprintf("Worker %d 处理了任务 %d: %s", workerID, task.ID, task.Data)
	if task.Result != nil {
		task.Result <- result
	}
}

func (wp *WorkerPool) Submit(task Task) {
	wp.tasks <- task
}

func (wp *WorkerPool) Stop() {
	close(wp.tasks)
	wp.wg.Wait()
}

func (wp *WorkerPool) ForceStop() {
	close(wp.quit)
	wp.wg.Wait()
}

func example4_dynamic_worker_pool() {
	fmt.Println("\n=== 示例4: 可控制的 Worker Pool ===")
	
	pool := NewWorkerPool(3)
	pool.Start()
	
	results := make(chan string, 10)
	
	// 提交任务
	for i := 1; i <= 10; i++ {
		pool.Submit(Task{
			ID:     i,
			Data:   fmt.Sprintf("数据-%d", i),
			Result: results,
		})
	}
	
	// 收集结果
	go func() {
		for i := 0; i < 10; i++ {
			result := <-results
			fmt.Println(result)
		}
		close(results)
	}()
	
	// 等待所有任务完成
	time.Sleep(2 * time.Second)
	pool.Stop()
	fmt.Println("Worker Pool 已停止")
}

// 示例5: 带超时控制的 Worker Pool
func example5_worker_pool_with_timeout() {
	fmt.Println("\n=== 示例5: 带超时控制的 Worker Pool ===")
	
	const numWorkers = 3
	tasks := make(chan int, 10)
	results := make(chan string, 10)
	timeout := time.After(2 * time.Second)
	
	// 启动 workers
	for w := 1; w <= numWorkers; w++ {
		go func(id int) {
			for task := range tasks {
				// 模拟长时间任务
				processingTime := time.Duration(rand.Intn(500)) * time.Millisecond
				time.Sleep(processingTime)
				result := fmt.Sprintf("Worker %d 完成任务 %d", id, task)
				
				select {
				case results <- result:
					// 成功发送结果
				case <-timeout:
					fmt.Printf("Worker %d: 超时，停止发送结果\n", id)
					return
				}
			}
		}(w)
	}
	
	// 发送任务
	go func() {
		for i := 1; i <= 20; i++ {
			select {
			case tasks <- i:
			case <-timeout:
				fmt.Println("超时，停止发送任务")
				close(tasks)
				return
			}
		}
		close(tasks)
	}()
	
	// 收集结果直到超时
	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("达到超时时间，停止接收结果")
			return
		}
	}
}

// 示例6: 限流的 Worker Pool（速率限制）
func example6_rate_limited_worker_pool() {
	fmt.Println("\n=== 示例6: 限流的 Worker Pool ===")
	
	const numWorkers = 5
	const rateLimit = 2 // 每秒最多处理2个任务
	
	tasks := make(chan int, 20)
	results := make(chan string, 20)
	rateLimiter := time.Tick(time.Second / rateLimit)
	
	// 启动 workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				<-rateLimiter // 等待速率限制器
				fmt.Printf("Worker %d: 处理任务 %d\n", id, task)
				time.Sleep(100 * time.Millisecond)
				results <- fmt.Sprintf("任务 %d 完成", task)
			}
		}(w)
	}
	
	// 发送任务
	go func() {
		for i := 1; i <= 10; i++ {
			tasks <- i
		}
		close(tasks)
	}()
	
	// 收集结果
	go func() {
		wg.Wait()
		close(results)
	}()
	
	for result := range results {
		fmt.Println(result)
	}
}

// 示例7: 带错误处理的 Worker Pool
type JobWithError struct {
	ID   int
	Data int
}

type ResultWithError struct {
	JobID  int
	Result int
	Err    error
}

func processWithError(job JobWithError) ResultWithError {
	time.Sleep(100 * time.Millisecond)
	
	// 模拟随机错误
	if rand.Float32() < 0.2 {
		return ResultWithError{
			JobID: job.ID,
			Err:   fmt.Errorf("处理失败: 任务 %d", job.ID),
		}
	}
	
	return ResultWithError{
		JobID:  job.ID,
		Result: job.Data * 2,
		Err:    nil,
	}
}

func example7_worker_pool_with_error_handling() {
	fmt.Println("\n=== 示例7: 带错误处理的 Worker Pool ===")
	
	const numWorkers = 3
	jobs := make(chan JobWithError, 10)
	results := make(chan ResultWithError, 10)
	var wg sync.WaitGroup
	
	// 启动 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				result := processWithError(job)
				results <- result
			}
		}(w)
	}
	
	// 发送任务
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- JobWithError{ID: i, Data: i * 10}
		}
		close(jobs)
	}()
	
	// 收集结果
	go func() {
		wg.Wait()
		close(results)
	}()
	
	successCount := 0
	errorCount := 0
	
	for result := range results {
		if result.Err != nil {
			fmt.Printf("❌ 任务 %d 失败: %v\n", result.JobID, result.Err)
			errorCount++
		} else {
			fmt.Printf("✅ 任务 %d 成功: 结果 = %d\n", result.JobID, result.Result)
			successCount++
		}
	}
	
	fmt.Printf("\n统计: 成功 %d, 失败 %d\n", successCount, errorCount)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("======================================")
	fmt.Println("    Go Worker Pool 模式示例")
	fmt.Println("======================================")
	
	example1_simple_worker_pool()
	example2_worker_pool_with_waitgroup()
	example3_structured_worker_pool()
	example4_dynamic_worker_pool()
	example5_worker_pool_with_timeout()
	example6_rate_limited_worker_pool()
	example7_worker_pool_with_error_handling()
	
	fmt.Println("\n所有示例运行完成！")
}
