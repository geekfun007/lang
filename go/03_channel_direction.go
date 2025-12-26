package main

import (
	"fmt"
	"time"
)

// 示例1: 单向 Channel 的基本概念
func example1_basic_direction() {
	fmt.Println("\n=== 示例1: 单向 Channel 基础 ===")
	
	// 双向 channel
	ch := make(chan int)
	
	// 只发送 channel
	var sendOnly chan<- int = ch
	
	// 只接收 channel
	var recvOnly <-chan int = ch
	
	go func() {
		sendOnly <- 42
	}()
	
	val := <-recvOnly
	fmt.Printf("接收到值: %d\n", val)
	
	// 编译时会报错的操作:
	// val := <-sendOnly  // 错误：只发送 channel 不能接收
	// recvOnly <- 42     // 错误：只接收 channel 不能发送
}

// 示例2: 函数参数中使用单向 Channel
// 这是最常见的使用场景
func producer(ch chan<- int) {
	fmt.Println("Producer: 开始生产数据")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: 发送 %d\n", i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	fmt.Println("Producer: 生产完成，关闭 channel")
}

func consumer(ch <-chan int) {
	fmt.Println("Consumer: 开始消费数据")
	for val := range ch {
		fmt.Printf("Consumer: 接收 %d\n", val)
	}
	fmt.Println("Consumer: 消费完成")
}

func example2_function_parameters() {
	fmt.Println("\n=== 示例2: 函数参数中的单向 Channel ===")
	
	ch := make(chan int, 10)
	
	go producer(ch)  // 自动转换为 chan<- int
	consumer(ch)     // 自动转换为 <-chan int
}

// 示例3: 只发送 Channel (chan<-)
func sendData(ch chan<- string, data []string) {
	for _, item := range data {
		ch <- item
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func example3_send_only() {
	fmt.Println("\n=== 示例3: 只发送 Channel ===")
	
	ch := make(chan string, 5)
	data := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	
	go sendData(ch, data)
	
	for fruit := range ch {
		fmt.Printf("收到水果: %s\n", fruit)
	}
}

// 示例4: 只接收 Channel (<-chan)
func receiveData(ch <-chan int) int {
	sum := 0
	for val := range ch {
		sum += val
		fmt.Printf("累加 %d, 当前和: %d\n", val, sum)
	}
	return sum
}

func example4_receive_only() {
	fmt.Println("\n=== 示例4: 只接收 Channel ===")
	
	ch := make(chan int)
	
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	
	total := receiveData(ch)
	fmt.Printf("总和: %d\n", total)
}

// 示例5: 多级流水线中的单向 Channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out // 返回只接收 channel
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func example5_pipeline() {
	fmt.Println("\n=== 示例5: 流水线中的单向 Channel ===")
	
	// 构建流水线: 生成器 -> 平方计算
	nums := generator(1, 2, 3, 4, 5)
	squares := square(nums)
	
	// 消费结果
	for result := range squares {
		fmt.Printf("结果: %d\n", result)
	}
}

// 示例6: 多个消费者，单个生产者
func multiProducer(ch chan<- int, id int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i*100 + id
		time.Sleep(50 * time.Millisecond)
	}
}

func multiConsumer(ch <-chan int, id int, done chan<- bool) {
	for val := range ch {
		fmt.Printf("消费者 %d 接收: %d\n", id, val)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true
}

func example6_multiple_consumers() {
	fmt.Println("\n=== 示例6: 多消费者模式 ===")
	
	ch := make(chan int, 10)
	done := make(chan bool)
	
	// 1个生产者
	go func() {
		multiProducer(ch, 1, 10)
		close(ch)
	}()
	
	// 3个消费者
	for i := 1; i <= 3; i++ {
		go multiConsumer(ch, i, done)
	}
	
	// 等待所有消费者完成
	for i := 1; i <= 3; i++ {
		<-done
	}
	fmt.Println("所有消费者已完成")
}

// 示例7: 使用单向 Channel 增强类型安全
type DataService struct {
	output <-chan string // 只允许外部读取
}

func NewDataService() *DataService {
	ch := make(chan string)
	
	// 内部可以写入
	go func() {
		data := []string{"Data1", "Data2", "Data3"}
		for _, item := range data {
			ch <- item
			time.Sleep(200 * time.Millisecond)
		}
		close(ch)
	}()
	
	return &DataService{output: ch}
}

func (ds *DataService) GetChannel() <-chan string {
	return ds.output
}

func example7_type_safety() {
	fmt.Println("\n=== 示例7: 类型安全的设计 ===")
	
	service := NewDataService()
	ch := service.GetChannel()
	
	// 外部只能读取，不能写入
	for data := range ch {
		fmt.Printf("接收到: %s\n", data)
	}
	
	// 编译时会报错:
	// ch <- "invalid"  // 错误：不能发送到只接收 channel
}

// 示例8: 单向 Channel 在结构体中的应用
type Worker struct {
	id     int
	input  <-chan int      // 只接收任务
	output chan<- int      // 只发送结果
}

func NewWorker(id int, input <-chan int, output chan<- int) *Worker {
	return &Worker{
		id:     id,
		input:  input,
		output: output,
	}
}

func (w *Worker) Run() {
	for task := range w.input {
		result := task * 2 // 简单处理：乘以2
		fmt.Printf("Worker %d 处理任务 %d -> %d\n", w.id, task, result)
		w.output <- result
		time.Sleep(100 * time.Millisecond)
	}
}

func example8_struct_with_channels() {
	fmt.Println("\n=== 示例8: 结构体中使用单向 Channel ===")
	
	tasks := make(chan int, 10)
	results := make(chan int, 10)
	
	// 创建3个 worker
	for i := 1; i <= 3; i++ {
		worker := NewWorker(i, tasks, results)
		go worker.Run()
	}
	
	// 发送任务
	go func() {
		for i := 1; i <= 9; i++ {
			tasks <- i
		}
		close(tasks)
	}()
	
	// 收集结果
	count := 0
	for result := range results {
		fmt.Printf("收到结果: %d\n", result)
		count++
		if count == 9 {
			close(results)
		}
	}
}

func main() {
	fmt.Println("======================================")
	fmt.Println("    Go Channel 方向控制示例")
	fmt.Println("======================================")
	
	example1_basic_direction()
	example2_function_parameters()
	example3_send_only()
	example4_receive_only()
	example5_pipeline()
	example6_multiple_consumers()
	example7_type_safety()
	example8_struct_with_channels()
	
	fmt.Println("\n所有示例运行完成！")
}
