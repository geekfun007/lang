package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 示例1: 简单的三阶段流水线
func example1_simple_pipeline() {
	fmt.Println("\n=== 示例1: 简单的三阶段流水线 ===")
	
	// 阶段1: 生成数字
	generator := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}
	
	// 阶段2: 平方计算
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				out <- n * n
			}
		}()
		return out
	}
	
	// 阶段3: 格式化输出
	format := func(in <-chan int) <-chan string {
		out := make(chan string)
		go func() {
			defer close(out)
			for n := range in {
				out <- fmt.Sprintf("结果: %d", n)
			}
		}()
		return out
	}
	
	// 构建流水线
	numbers := generator(1, 2, 3, 4, 5)
	squares := square(numbers)
	results := format(squares)
	
	// 消费结果
	for result := range results {
		fmt.Println(result)
	}
}

// 示例2: 可取消的流水线
func example2_cancellable_pipeline() {
	fmt.Println("\n=== 示例2: 可取消的流水线 ===")
	
	// 生成器（支持取消）
	generator := func(done <-chan struct{}, nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				select {
				case out <- n:
				case <-done:
					fmt.Println("Generator: 收到取消信号")
					return
				}
			}
		}()
		return out
	}
	
	// 处理器（支持取消）
	multiply := func(done <-chan struct{}, in <-chan int, factor int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				select {
				case out <- n * factor:
				case <-done:
					fmt.Println("Multiply: 收到取消信号")
					return
				}
			}
		}()
		return out
	}
	
	done := make(chan struct{})
	
	// 构建流水线
	numbers := generator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	results := multiply(done, numbers, 2)
	
	// 只处理3个结果后取消
	count := 0
	for n := range results {
		fmt.Printf("收到: %d\n", n)
		count++
		if count == 3 {
			fmt.Println("Main: 发送取消信号")
			close(done)
			break
		}
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("流水线已取消")
}

// 示例3: 扇出-扇入模式（并行处理）
func example3_fan_out_fan_in() {
	fmt.Println("\n=== 示例3: 扇出-扇入模式 ===")
	
	// 生成器
	generator := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}
	
	// 处理器（模拟耗时操作）
	process := func(id int, in <-chan int) <-chan string {
		out := make(chan string)
		go func() {
			defer close(out)
			for n := range in {
				time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
				result := fmt.Sprintf("Worker %d 处理了 %d -> %d", id, n, n*n)
				out <- result
			}
		}()
		return out
	}
	
	// 合并多个 channel
	merge := func(channels ...<-chan string) <-chan string {
		out := make(chan string)
		var wg sync.WaitGroup
		
		wg.Add(len(channels))
		for _, ch := range channels {
			go func(c <-chan string) {
				defer wg.Done()
				for v := range c {
					out <- v
				}
			}(ch)
		}
		
		go func() {
			wg.Wait()
			close(out)
		}()
		
		return out
	}
	
	// 构建流水线
	input := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	
	// 扇出：启动3个并行处理器
	c1 := process(1, input)
	c2 := process(2, input)
	c3 := process(3, input)
	
	// 扇入：合并结果
	results := merge(c1, c2, c3)
	
	// 消费结果
	for result := range results {
		fmt.Println(result)
	}
}

// 示例4: 流水线中的错误处理
type Result struct {
	Value int
	Err   error
}

func example4_pipeline_with_errors() {
	fmt.Println("\n=== 示例4: 带错误处理的流水线 ===")
	
	// 生成器
	generator := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}
	
	// 处理器（可能产生错误）
	process := func(in <-chan int) <-chan Result {
		out := make(chan Result)
		go func() {
			defer close(out)
			for n := range in {
				// 模拟错误：奇数会失败
				if n%2 != 0 {
					out <- Result{
						Value: 0,
						Err:   fmt.Errorf("无法处理奇数: %d", n),
					}
				} else {
					out <- Result{
						Value: n * n,
						Err:   nil,
					}
				}
			}
		}()
		return out
	}
	
	// 构建流水线
	numbers := generator(1, 2, 3, 4, 5, 6)
	results := process(numbers)
	
	// 处理结果
	for result := range results {
		if result.Err != nil {
			fmt.Printf("❌ 错误: %v\n", result.Err)
		} else {
			fmt.Printf("✅ 成功: %d\n", result.Value)
		}
	}
}

// 示例5: 带缓冲的流水线
func example5_buffered_pipeline() {
	fmt.Println("\n=== 示例5: 带缓冲的流水线 ===")
	
	generator := func(nums ...int) <-chan int {
		out := make(chan int, 10) // 使用缓冲
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
				fmt.Printf("生成: %d\n", n)
			}
			fmt.Println("生成器完成")
		}()
		return out
	}
	
	slow := func(in <-chan int) <-chan int {
		out := make(chan int, 10)
		go func() {
			defer close(out)
			for n := range in {
				time.Sleep(500 * time.Millisecond) // 慢速处理
				out <- n * 2
				fmt.Printf("处理: %d -> %d\n", n, n*2)
			}
			fmt.Println("处理器完成")
		}()
		return out
	}
	
	// 构建流水线
	numbers := generator(1, 2, 3, 4, 5)
	results := slow(numbers)
	
	// 慢速消费
	for result := range results {
		fmt.Printf("收到结果: %d\n", result)
		time.Sleep(300 * time.Millisecond)
	}
}

// 示例6: 带监控的流水线
type Stats struct {
	Processed int
	Errors    int
	Duration  time.Duration
}

func example6_monitored_pipeline() {
	fmt.Println("\n=== 示例6: 带监控的流水线 ===")
	
	generator := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}
	
	monitored := func(in <-chan int) (<-chan int, <-chan Stats) {
		out := make(chan int)
		stats := make(chan Stats)
		
		go func() {
			defer close(out)
			defer close(stats)
			
			start := time.Now()
			processed := 0
			errors := 0
			
			for n := range in {
				// 模拟处理
				time.Sleep(50 * time.Millisecond)
				processed++
				
				if rand.Float32() < 0.2 {
					errors++
				} else {
					out <- n * n
				}
				
				// 定期发送统计
				if processed%3 == 0 {
					stats <- Stats{
						Processed: processed,
						Errors:    errors,
						Duration:  time.Since(start),
					}
				}
			}
			
			// 最终统计
			stats <- Stats{
				Processed: processed,
				Errors:    errors,
				Duration:  time.Since(start),
			}
		}()
		
		return out, stats
	}
	
	// 构建流水线
	numbers := generator(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	results, stats := monitored(numbers)
	
	// 同时消费结果和统计
	done := make(chan bool)
	
	go func() {
		for stat := range stats {
			fmt.Printf("📊 统计: 已处理 %d, 错误 %d, 耗时 %v\n",
				stat.Processed, stat.Errors, stat.Duration)
		}
	}()
	
	go func() {
		for result := range results {
			fmt.Printf("结果: %d\n", result)
		}
		done <- true
	}()
	
	<-done
	time.Sleep(100 * time.Millisecond)
}

// 示例7: 动态流水线（根据数据动态添加阶段）
func example7_dynamic_pipeline() {
	fmt.Println("\n=== 示例7: 动态流水线 ===")
	
	// 基础生成器
	generator := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for _, n := range nums {
				out <- n
			}
		}()
		return out
	}
	
	// 动态应用转换函数
	transform := func(in <-chan int, fn func(int) int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for n := range in {
				out <- fn(n)
			}
		}()
		return out
	}
	
	// 构建流水线
	pipeline := generator(1, 2, 3, 4, 5)
	
	// 动态添加转换阶段
	transformations := []struct {
		name string
		fn   func(int) int
	}{
		{"加10", func(n int) int { return n + 10 }},
		{"乘2", func(n int) int { return n * 2 }},
		{"减5", func(n int) int { return n - 5 }},
	}
	
	for _, t := range transformations {
		fmt.Printf("添加转换: %s\n", t.name)
		pipeline = transform(pipeline, t.fn)
	}
	
	// 消费结果
	fmt.Println("\n结果:")
	for result := range pipeline {
		fmt.Printf("  %d\n", result)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("======================================")
	fmt.Println("    Go Pipeline 模式示例")
	fmt.Println("======================================")
	
	example1_simple_pipeline()
	example2_cancellable_pipeline()
	example3_fan_out_fan_in()
	example4_pipeline_with_errors()
	example5_buffered_pipeline()
	example6_monitored_pipeline()
	example7_dynamic_pipeline()
	
	fmt.Println("\n所有示例运行完成！")
}
