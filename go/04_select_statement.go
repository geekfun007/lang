package main

import (
	"fmt"
	"time"
)

// 示例1: Select 基本用法
func example1_basic_select() {
	fmt.Println("\n=== 示例1: Select 基本用法 ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自 channel 1"
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "来自 channel 2"
	}()
	
	// select 会等待第一个就绪的 channel
	select {
	case msg1 := <-ch1:
		fmt.Println("接收:", msg1)
	case msg2 := <-ch2:
		fmt.Println("接收:", msg2)
	}
}

// 示例2: Select 接收多个消息
func example2_multiple_receives() {
	fmt.Println("\n=== 示例2: 接收多个消息 ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(300 * time.Millisecond)
			ch1 <- fmt.Sprintf("CH1-%d", i)
		}
	}()
	
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(500 * time.Millisecond)
			ch2 <- fmt.Sprintf("CH2-%d", i)
		}
	}()
	
	// 接收6个消息
	for i := 0; i < 6; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("从 CH1 接收:", msg1)
		case msg2 := <-ch2:
			fmt.Println("从 CH2 接收:", msg2)
		}
	}
}

// 示例3: Select 的 Default 分支（非阻塞）
func example3_select_default() {
	fmt.Println("\n=== 示例3: Select Default 分支 ===")
	
	ch := make(chan int, 1)
	
	// 非阻塞发送
	select {
	case ch <- 42:
		fmt.Println("发送成功")
	default:
		fmt.Println("缓冲区满，发送失败")
	}
	
	// 非阻塞接收
	select {
	case val := <-ch:
		fmt.Printf("接收到: %d\n", val)
	default:
		fmt.Println("Channel 为空，接收失败")
	}
	
	// 再次尝试接收（此时 channel 为空）
	select {
	case val := <-ch:
		fmt.Printf("接收到: %d\n", val)
	default:
		fmt.Println("Channel 为空，立即返回")
	}
}

// 示例4: 超时处理
func example4_timeout() {
	fmt.Println("\n=== 示例4: 超时处理 ===")
	
	ch := make(chan string)
	
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "结果数据"
	}()
	
	select {
	case result := <-ch:
		fmt.Println("收到结果:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("超时！操作耗时过长")
	}
}

// 示例5: 定时器 (Ticker)
func example5_ticker() {
	fmt.Println("\n=== 示例5: 使用 Ticker 定时执行 ===")
	
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	done := make(chan bool)
	
	go func() {
		time.Sleep(2500 * time.Millisecond)
		done <- true
	}()
	
	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("Tick %d\n", count)
		case <-done:
			fmt.Println("完成！")
			return
		}
	}
}

// 示例6: Select 的随机性（多个 case 同时就绪）
func example6_random_selection() {
	fmt.Println("\n=== 示例6: Select 的随机选择 ===")
	
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	// 两个 channel 都提前准备好数据
	ch1 <- "Channel 1"
	ch2 <- "Channel 2"
	
	// 多次 select，观察随机性
	for i := 1; i <= 5; i++ {
		// 重新填充数据
		select {
		case ch1 <- "Channel 1":
		default:
		}
		select {
		case ch2 <- "Channel 2":
		default:
		}
		
		// 随机选择一个
		select {
		case msg := <-ch1:
			fmt.Printf("尝试 %d: 选择了 %s\n", i, msg)
		case msg := <-ch2:
			fmt.Printf("尝试 %d: 选择了 %s\n", i, msg)
		}
	}
}

// 示例7: 使用 Select 实现优先级
func example7_priority() {
	fmt.Println("\n=== 示例7: 实现 Channel 优先级 ===")
	
	high := make(chan string, 10)
	low := make(chan string, 10)
	
	// 填充数据
	for i := 1; i <= 5; i++ {
		high <- fmt.Sprintf("高优先级-%d", i)
		low <- fmt.Sprintf("低优先级-%d", i)
	}
	close(high)
	close(low)
	
	// 优先处理高优先级
	for {
		select {
		case msg, ok := <-high:
			if !ok {
				// 高优先级处理完，处理低优先级
				for msg := range low {
					fmt.Println("处理:", msg)
				}
				return
			}
			fmt.Println("处理:", msg)
		default:
			// 高优先级没有数据时，处理低优先级
			select {
			case msg, ok := <-low:
				if !ok {
					low = nil // 防止再次选择
				} else {
					fmt.Println("处理:", msg)
				}
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}

// 示例8: Select 中的发送操作
func example8_send_in_select() {
	fmt.Println("\n=== 示例8: Select 中的发送操作 ===")
	
	ch := make(chan int)
	quit := make(chan bool)
	
	// 消费者
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("接收:", <-ch)
		}
		quit <- true
	}()
	
	// 生产者
	x := 0
	for {
		select {
		case ch <- x:
			fmt.Printf("发送: %d\n", x)
			x++
		case <-quit:
			fmt.Println("退出")
			return
		}
	}
}

// 示例9: 使用 Select 实现取消机制
func example9_cancellation() {
	fmt.Println("\n=== 示例9: 取消机制 ===")
	
	cancel := make(chan struct{})
	done := make(chan bool)
	
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		
		count := 0
		for {
			select {
			case <-ticker.C:
				count++
				fmt.Printf("工作中... %d\n", count)
			case <-cancel:
				fmt.Println("收到取消信号，停止工作")
				done <- true
				return
			}
		}
	}()
	
	// 2秒后取消
	time.Sleep(2 * time.Second)
	fmt.Println("发送取消信号")
	close(cancel)
	<-done
}

// 示例10: 组合使用 Select 和 Nil Channel
func example10_nil_channel() {
	fmt.Println("\n=== 示例10: Nil Channel 在 Select 中的行为 ===")
	
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "消息1"
		time.Sleep(500 * time.Millisecond)
		ch1 <- "消息2"
		close(ch1)
	}()
	
	go func() {
		time.Sleep(800 * time.Millisecond)
		ch2 <- "来自CH2"
		close(ch2)
	}()
	
	// 接收所有消息
	for ch1 != nil || ch2 != nil {
		select {
		case msg, ok := <-ch1:
			if !ok {
				fmt.Println("CH1 已关闭")
				ch1 = nil // 设为 nil 后不会再被选中
			} else {
				fmt.Println("从 CH1:", msg)
			}
		case msg, ok := <-ch2:
			if !ok {
				fmt.Println("CH2 已关闭")
				ch2 = nil
			} else {
				fmt.Println("从 CH2:", msg)
			}
		}
	}
	fmt.Println("所有 channel 已处理完毕")
}

func main() {
	fmt.Println("======================================")
	fmt.Println("    Go Select 语句示例")
	fmt.Println("======================================")
	
	example1_basic_select()
	example2_multiple_receives()
	example3_select_default()
	example4_timeout()
	example5_ticker()
	example6_random_selection()
	example7_priority()
	example8_send_in_select()
	example9_cancellation()
	example10_nil_channel()
	
	fmt.Println("\n所有示例运行完成！")
}
