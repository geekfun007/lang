package main

import (
	"fmt"
	"time"
)

// 示例1: 最基础的 Channel 使用
func example1_simple() {
	fmt.Println("\n=== 示例1: 简单的 Channel 通信 ===")

	// 创建一个无缓冲 channel
	ch := make(chan string)

	// 启动 goroutine 发送数据
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from goroutine!"
	}()

	// 主 goroutine 接收数据（会阻塞直到有数据）
	msg := <-ch
	fmt.Println("收到消息:", msg)
}

// 示例2: 无缓冲 Channel 的同步特性
func example2_synchronization() {
	fmt.Println("\n=== 示例2: 无缓冲 Channel 的同步特性 ===")

	ch := make(chan int)

	go func() {
		fmt.Println("Goroutine: 准备发送数据...")
		ch <- 42 // 会阻塞，直到有接收者
		fmt.Println("Goroutine: 数据已被接收！")
	}()

	time.Sleep(2 * time.Second) // 模拟主 goroutine 做其他事情
	fmt.Println("Main: 准备接收数据...")
	data := <-ch
	fmt.Printf("Main: 收到数据 %d\n", data)
}

// 示例3: 使用 Channel 等待 Goroutine 完成
func example3_done_signal() {
	fmt.Println("\n=== 示例3: 使用 Channel 作为完成信号 ===")

	done := make(chan bool)

	go func() {
		fmt.Println("工作中...")
		time.Sleep(2 * time.Second)
		fmt.Println("工作完成！")
		done <- true // 发送完成信号
	}()

	fmt.Println("等待工作完成...")
	<-done // 阻塞直到收到完成信号
	fmt.Println("所有工作已完成，程序退出")
}

// 示例4: 接收值和状态
func example4_receive_with_status() {
	fmt.Println("\n=== 示例4: 接收值和 Channel 状态 ===")

	ch := make(chan int, 2)

	// 发送一些值
	ch <- 1
	ch <- 2
	close(ch) // 关闭 channel

	// 接收所有值
	for {
		// 第二个返回值表示 channel 是否还开着
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel 已关闭，退出循环")
			break
		}
		fmt.Printf("收到值: %d\n", val)
	}

	// 从已关闭的 channel 接收会返回零值
	val, ok := <-ch
	fmt.Printf("从已关闭的 channel 接收: 值=%d, ok=%v\n", val, ok)
}

// 示例5: 使用 range 遍历 Channel
func example5_range_over_channel() {
	fmt.Println("\n=== 示例5: 使用 range 遍历 Channel ===")

	ch := make(chan int, 5)

	// 发送数据
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * 10
			time.Sleep(200 * time.Millisecond)
		}
		close(ch) // 必须关闭，否则 range 会永久阻塞
	}()

	// 使用 range 接收数据
	fmt.Println("开始接收数据:")
	for val := range ch {
		fmt.Printf("  收到: %d\n", val)
	}
	fmt.Println("Channel 已关闭，range 循环结束")
}

// 示例6: 多个 Goroutine 发送到同一个 Channel
func example6_multiple_senders() {
	fmt.Println("\n=== 示例6: 多个发送者 ===")

	ch := make(chan string, 10)

	// 启动3个发送者
	for i := 1; i <= 3; i++ {
		id := i
		go func() {
			for j := 1; j <= 3; j++ {
				msg := fmt.Sprintf("发送者-%d: 消息-%d", id, j)
				ch <- msg
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	// 接收所有消息
	time.Sleep(1 * time.Second) // 等待所有消息发送完
	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}

// 示例7: Channel 作为函数参数
func sender(ch chan<- int) { // 只能发送
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) { // 只能接收
	for val := range ch {
		fmt.Printf("接收到: %d\n", val)
	}
}

func example7_channel_as_parameter() {
	fmt.Println("\n=== 示例7: Channel 作为函数参数 ===")

	ch := make(chan int)

	go sender(ch)
	receiver(ch)
}

func main() {
	fmt.Println("======================================")
	fmt.Println("    Go Channel 基础示例")
	fmt.Println("======================================")

	example1_simple()
	example2_synchronization()
	example3_done_signal()
	example4_receive_with_status()
	example5_range_over_channel()
	example6_multiple_senders()
	example7_channel_as_parameter()

	fmt.Println("\n所有示例运行完成！")
}
