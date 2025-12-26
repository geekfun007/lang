package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 示例1: Context 基本取消
func example1_basic_context_cancel() {
	fmt.Println("\n=== 示例1: Context 基本取消 ===")
	
	ctx, cancel := context.WithCancel(context.Background())
	
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine: 收到取消信号，退出")
				return
			default:
				fmt.Println("Goroutine: 工作中...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	
	// 2秒后取消
	time.Sleep(2 * time.Second)
	fmt.Println("Main: 发送取消信号")
	cancel()
	time.Sleep(1 * time.Second)
}

// 示例2: Context 超时
func example2_context_timeout() {
	fmt.Println("\n=== 示例2: Context 超时 ===")
	
	// 创建一个1秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	result := make(chan string, 1)
	
	go func() {
		// 模拟耗时操作
		time.Sleep(2 * time.Second)
		result <- "操作完成"
	}()
	
	select {
	case res := <-result:
		fmt.Println("收到结果:", res)
	case <-ctx.Done():
		fmt.Println("操作超时:", ctx.Err())
	}
}

// 示例3: Context 截止时间
func example3_context_deadline() {
	fmt.Println("\n=== 示例3: Context 截止时间 ===")
	
	// 设置2秒后的截止时间
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	
	go func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()
		
		count := 0
		for {
			select {
			case <-ticker.C:
				count++
				fmt.Printf("Tick %d\n", count)
			case <-ctx.Done():
				fmt.Println("达到截止时间:", ctx.Err())
				return
			}
		}
	}()
	
	time.Sleep(3 * time.Second)
}

// 示例4: Context 传递值
func example4_context_with_value() {
	fmt.Println("\n=== 示例4: Context 传递值 ===")
	
	type key string
	const userKey key = "user"
	const requestIDKey key = "requestID"
	
	// 创建带值的 context
	ctx := context.Background()
	ctx = context.WithValue(ctx, userKey, "张三")
	ctx = context.WithValue(ctx, requestIDKey, "req-12345")
	
	// 定义 doWork 函数
	doWork := func(ctx context.Context) {
		user := ctx.Value(userKey).(string)
		requestID := ctx.Value(requestIDKey).(string)
		fmt.Printf("  子任务 [%s] 用户: %s\n", requestID, user)
	}
	
	processRequest := func(ctx context.Context) {
		user := ctx.Value(userKey).(string)
		requestID := ctx.Value(requestIDKey).(string)
		
		fmt.Printf("处理请求 [%s] 用户: %s\n", requestID, user)
		
		// 传递给子函数
		doWork(ctx)
	}
	
	processRequest(ctx)
}

// 示例5: 父子 Context 级联取消
func example5_context_hierarchy() {
	fmt.Println("\n=== 示例5: Context 层级取消 ===")
	
	// 创建父 context
	parentCtx, parentCancel := context.WithCancel(context.Background())
	
	// 创建子 context
	childCtx1, _ := context.WithCancel(parentCtx)
	childCtx2, _ := context.WithCancel(parentCtx)
	
	worker := func(id int, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Worker %d: 收到取消信号\n", id)
				return
			default:
				fmt.Printf("Worker %d: 工作中...\n", id)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
	
	go worker(1, childCtx1)
	go worker(2, childCtx2)
	
	time.Sleep(2 * time.Second)
	fmt.Println("取消父 Context")
	parentCancel() // 取消父 context，所有子 context 也会被取消
	
	time.Sleep(1 * time.Second)
}

// 示例6: 使用 Context 的 Worker Pool
func example6_worker_pool_with_context() {
	fmt.Println("\n=== 示例6: 带 Context 的 Worker Pool ===")
	
	type Task struct {
		ID   int
		Data string
	}
	
	workerPool := func(ctx context.Context, numWorkers int, tasks <-chan Task) <-chan string {
		results := make(chan string, numWorkers)
		var wg sync.WaitGroup
		
		for w := 1; w <= numWorkers; w++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for {
					select {
					case task, ok := <-tasks:
						if !ok {
							fmt.Printf("Worker %d: 任务队列已关闭\n", id)
							return
						}
						// 模拟处理
						time.Sleep(time.Duration(100+rand.Intn(300)) * time.Millisecond)
						result := fmt.Sprintf("Worker %d 完成任务 %d: %s", id, task.ID, task.Data)
						
						select {
						case results <- result:
						case <-ctx.Done():
							fmt.Printf("Worker %d: Context 取消\n", id)
							return
						}
					case <-ctx.Done():
						fmt.Printf("Worker %d: 收到取消信号\n", id)
						return
					}
				}
			}(w)
		}
		
		go func() {
			wg.Wait()
			close(results)
		}()
		
		return results
	}
	
	// 创建带超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	tasks := make(chan Task, 10)
	results := workerPool(ctx, 3, tasks)
	
	// 发送任务
	go func() {
		for i := 1; i <= 20; i++ {
			select {
			case tasks <- Task{ID: i, Data: fmt.Sprintf("数据-%d", i)}:
			case <-ctx.Done():
				fmt.Println("停止发送任务")
				close(tasks)
				return
			}
		}
		close(tasks)
	}()
	
	// 收集结果
	for result := range results {
		fmt.Println(result)
	}
	
	fmt.Println("Worker Pool 已停止")
}

// 示例7: Pipeline 中使用 Context
func example7_pipeline_with_context() {
	fmt.Println("\n=== 示例7: Pipeline 中使用 Context ===")
	
	// 生成器
	generator := func(ctx context.Context, nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				select {
				case out <- n:
				case <-ctx.Done():
					fmt.Println("Generator: 取消")
					return
				}
			}
		}()
		return out
	}
	
	// 处理器
	square := func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				select {
				case out <- n * n:
					time.Sleep(500 * time.Millisecond)
				case <-ctx.Done():
					fmt.Println("Square: 取消")
					return
				}
			}
		}()
		return out
	}
	
	// 创建带取消的 context
	ctx, cancel := context.WithCancel(context.Background())
	
	// 构建 pipeline
	numbers := generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squares := square(ctx, numbers)
	
	// 只读取3个结果后取消
	count := 0
	for n := range squares {
		fmt.Printf("收到: %d\n", n)
		count++
		if count == 3 {
			fmt.Println("Main: 取消 Pipeline")
			cancel()
			break
		}
	}
	
	time.Sleep(500 * time.Millisecond)
}

// 示例8: 多个 Goroutine 协调退出
func example8_coordinated_shutdown() {
	fmt.Println("\n=== 示例8: 协调多个 Goroutine 退出 ===")
	
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	
	// 启动多个服务
	services := []string{"WebServer", "Database", "Cache", "Logger"}
	
	for _, name := range services {
		wg.Add(1)
		go func(serviceName string) {
			defer wg.Done()
			defer fmt.Printf("%s: 已停止\n", serviceName)
			
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()
			
			for {
				select {
				case <-ticker.C:
					fmt.Printf("%s: 运行中...\n", serviceName)
				case <-ctx.Done():
					fmt.Printf("%s: 开始清理资源...\n", serviceName)
					time.Sleep(200 * time.Millisecond)
					return
				}
			}
		}(name)
	}
	
	// 2秒后触发关闭
	time.Sleep(2 * time.Second)
	fmt.Println("\n=== 触发关闭信号 ===")
	cancel()
	
	// 等待所有服务停止
	wg.Wait()
	fmt.Println("所有服务已停止")
}

// 示例9: 带重试的 Context
func example9_context_with_retry() {
	fmt.Println("\n=== 示例9: 带重试的操作 ===")
	
	// 模拟可能失败的操作
	doWork := func(ctx context.Context, attempt int) error {
		fmt.Printf("尝试 %d...\n", attempt)
		
		select {
		case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
			// 模拟30%的失败率
			if rand.Float32() < 0.3 {
				return fmt.Errorf("操作失败")
			}
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	
	// 带重试的执行
	executeWithRetry := func(ctx context.Context, maxRetries int) error {
		for i := 1; i <= maxRetries; i++ {
			err := doWork(ctx, i)
			if err == nil {
				fmt.Println("操作成功！")
				return nil
			}
			
			if err == context.DeadlineExceeded || err == context.Canceled {
				return err
			}
			
			fmt.Printf("失败: %v，重试中...\n", err)
			time.Sleep(500 * time.Millisecond)
		}
		return fmt.Errorf("达到最大重试次数")
	}
	
	// 创建带超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err := executeWithRetry(ctx, 5)
	if err != nil {
		fmt.Printf("最终失败: %v\n", err)
	}
}

// 示例10: Context 在 HTTP 请求中的应用
func example10_context_http_simulation() {
	fmt.Println("\n=== 示例10: 模拟 HTTP 请求处理 ===")
	
	type Request struct {
		ID   string
		User string
	}
	
	// 模拟数据库查询
	queryDatabase := func(ctx context.Context, query string) (string, error) {
		select {
		case <-time.After(500 * time.Millisecond):
			return fmt.Sprintf("结果: %s", query), nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	
	// 模拟缓存查询
	queryCache := func(ctx context.Context, key string) (string, error) {
		select {
		case <-time.After(100 * time.Millisecond):
			return "", fmt.Errorf("缓存未命中")
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
	
	// 请求处理器
	handleRequest := func(ctx context.Context, req Request) {
		fmt.Printf("\n处理请求 [%s] 用户: %s\n", req.ID, req.User)
		
		// 尝试从缓存获取
		result, err := queryCache(ctx, req.ID)
		if err == nil {
			fmt.Printf("从缓存返回: %s\n", result)
			return
		}
		
		// 缓存未命中，查询数据库
		fmt.Println("缓存未命中，查询数据库...")
		result, err = queryDatabase(ctx, req.User)
		if err != nil {
			fmt.Printf("请求失败: %v\n", err)
			return
		}
		
		fmt.Printf("请求成功: %s\n", result)
	}
	
	// 模拟3个并发请求
	requests := []Request{
		{ID: "req-1", User: "Alice"},
		{ID: "req-2", User: "Bob"},
		{ID: "req-3", User: "Charlie"},
	}
	
	var wg sync.WaitGroup
	for _, req := range requests {
		wg.Add(1)
		go func(r Request) {
			defer wg.Done()
			
			// 每个请求有1秒超时
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			
			handleRequest(ctx, r)
		}(req)
	}
	
	wg.Wait()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("======================================")
	fmt.Println("    Go Context 取消模式示例")
	fmt.Println("======================================")
	
	example1_basic_context_cancel()
	example2_context_timeout()
	example3_context_deadline()
	example4_context_with_value()
	example5_context_hierarchy()
	example6_worker_pool_with_context()
	example7_pipeline_with_context()
	example8_coordinated_shutdown()
	example9_context_with_retry()
	example10_context_http_simulation()
	
	fmt.Println("\n所有示例运行完成！")
}
