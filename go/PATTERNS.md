# Go Channel 并发模式图解

本文档使用图解方式展示常见的 Go Channel 并发模式。

## 目录

1. [基础通信模式](#基础通信模式)
2. [Generator 生成器模式](#generator-生成器模式)
3. [Pipeline 流水线模式](#pipeline-流水线模式)
4. [Fan-Out/Fan-In 扇出扇入](#fan-outfan-in-扇出扇入)
5. [Worker Pool 工作池](#worker-pool-工作池)
6. [超时和取消](#超时和取消)
7. [请求-响应模式](#请求-响应模式)

---

## 基础通信模式

### 1. 简单的点对点通信

```
Goroutine A                    Goroutine B
    |                              |
    | ch <- data                   |
    |----------------------------->|
    |         (发送)         (接收) | <- ch
    |                              |
```

**代码：**
```go
ch := make(chan int)

// Goroutine A
go func() {
    ch <- 42
}()

// Goroutine B
value := <-ch
```

---

### 2. 无缓冲 vs 缓冲 Channel

#### 无缓冲 Channel（同步）

```
发送者                 接收者
  |                      |
  | ch <- 1 (阻塞)       |
  |------ 等待 ----------|
  |                      | <- ch
  |<----- 同时 -------->|
  | (发送完成)    (接收完成)
```

#### 缓冲 Channel（异步）

```
发送者      缓冲区[3]       接收者
  |                          |
  | ch <- 1                  |
  |--->[1|_|_]               |
  |                          |
  | ch <- 2                  |
  |--->[1|2|_]               |
  |                          |
  | ch <- 3                  |
  |--->[1|2|3]               |
  |                          |
  | ch <- 4 (阻塞)           |
  |     缓冲区满！            |
  |                          | <- ch
  |                    [2|3|_]<---|
  |<-- 可以发送了！           |
```

---

## Generator 生成器模式

```
              Generator
        ┌──────────────────┐
        │  func gen() {    │
        │    for ... {     │
        │      ch <- val   │
        │    }             │
        │    close(ch)     │
        │  }               │
        └────────┬─────────┘
                 │
                 ▼
            ┌─────────┐
            │ Channel │──────> Consumer
            └─────────┘        (for v := range ch)
```

**代码：**
```go
func generator(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

// 使用
for v := range generator(1, 2, 3, 4, 5) {
    fmt.Println(v)
}
```

---

## Pipeline 流水线模式

```
Generator ──> Stage 1 ──> Stage 2 ──> Stage 3 ──> Consumer
  [1,2,3]      [2,4,6]     [4,8,12]    [8,16,24]   (使用)
   (生成)        (*2)        (*2)         (*2)
```

**完整流程：**
```
┌──────────┐      ┌──────────┐      ┌──────────┐      ┌──────────┐
│ Generate │ ch1  │  Square  │ ch2  │  Double  │ ch3  │ Consumer │
│  1,2,3   │─────>│  1,4,9   │─────>│  2,8,18  │─────>│  Process │
└──────────┘      └──────────┘      └──────────┘      └──────────┘
```

**代码：**
```go
// 阶段1: 生成
gen := func(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

// 阶段2: 平方
sq := func(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

// 连接流水线
nums := gen(1, 2, 3)
squares := sq(nums)
for v := range squares {
    fmt.Println(v)
}
```

---

## Fan-Out/Fan-In 扇出扇入

### Fan-Out (一个输入，多个处理器)

```
                    ┌──> Worker 1 ──┐
                    │                │
Input Channel ──────┼──> Worker 2 ──┼──> Output Channel
                    │                │
                    └──> Worker 3 ──┘
```

### Fan-In (多个输入，一个输出)

```
Worker 1 ──┐
           │
Worker 2 ──┼──> Merge ──> Output Channel
           │
Worker 3 ──┘
```

### 完整的 Fan-Out/Fan-In 流程

```
                      ┌──> Worker 1 ──┐
                      │   (处理 1,4,7) │
Generator ──> Input ──┼──> Worker 2 ──┼──> Merge ──> Consumer
  [1...9]             │   (处理 2,5,8) │
                      └──> Worker 3 ──┘
                          (处理 3,6,9)
```

**代码：**
```go
// Fan-Out: 启动多个 worker
func fanOut(in <-chan int, n int) []<-chan int {
    channels := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        channels[i] = worker(in)
    }
    return channels
}

// Fan-In: 合并多个 channel
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
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
```

---

## Worker Pool 工作池

```
                     ┌─────────────┐
                     │  Worker 1   │
                     │  (空闲)     │
                     └──────┬──────┘
                            │
                     ┌──────▼──────┐
    Tasks ──────────>│ Task Queue  │
    [T1,T2,T3,...]   │  [T2,T3]    │
                     └──────┬──────┘
                            │
                     ┌──────▼──────┐
                     │  Worker 2   │───> Results
                     │  处理 T1    │      Channel
                     └─────────────┘
                            
                     ┌─────────────┐
                     │  Worker 3   │
                     │  (空闲)     │
                     └─────────────┘
```

**工作流程：**
```
1. 任务进入队列
   Tasks ──> [T1|T2|T3|T4|T5]

2. Worker 竞争获取任务
   Worker 1: 获取 T1
   Worker 2: 获取 T2
   Worker 3: 获取 T3

3. 完成后继续获取
   Worker 1: T1 完成 ──> 获取 T4
   Worker 2: T2 完成 ──> 获取 T5
```

**代码：**
```go
func workerPool(numWorkers int, tasks <-chan Task) <-chan Result {
    results := make(chan Result)
    var wg sync.WaitGroup
    
    // 启动固定数量的 worker
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for task := range tasks {
                result := process(task)
                results <- result
            }
        }(i)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
    
    return results
}
```

---

## 超时和取消

### 1. 超时模式

```
Goroutine                      Select
    |                             |
    |--- 发送请求 -------------->  |
    |                             |
    |                        ┌────┴────┐
    |                        │ case    │
    |                        │ result  │
    |<------ 收到结果 ------  │  <-ch   │
    |                        ├─────────┤
    |                        │ case    │
    |                        │timeout  │
    |                        │ <-timer │
    |                        └─────────┘
```

**代码：**
```go
select {
case result := <-ch:
    fmt.Println("成功:", result)
case <-time.After(1 * time.Second):
    fmt.Println("超时")
}
```

### 2. Context 取消传播

```
Parent Context
    │
    ├──> Child Context 1 ──> Worker 1
    │                            │
    ├──> Child Context 2 ──> Worker 2
    │                            │
    └──> Child Context 3 ──> Worker 3
         
cancel() ───> 所有 Worker 收到取消信号
```

**代码：**
```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

for i := 0; i < 3; i++ {
    go func(id int) {
        for {
            select {
            case <-ctx.Done():
                fmt.Printf("Worker %d 退出\n", id)
                return
            default:
                // 工作
            }
        }
    }(i)
}

time.Sleep(2 * time.Second)
cancel()  // 所有 worker 都会退出
```

---

## 请求-响应模式

### 方式1: 通过返回 Channel

```
Client                         Server
  |                               |
  | ch := make(chan Response)    |
  |                               |
  | req := Request{              |
  |   Data: "hello",             |
  |   RespCh: ch                 |
  | }                            |
  |                               |
  | requests <- req ───────────> |
  |                               |
  |                        处理请求 |
  |                               |
  | result := <-ch <──────────── | ch <- response
  |                               |
```

**代码：**
```go
type Request struct {
    Data     string
    Response chan<- string
}

// Server
func server(requests <-chan Request) {
    for req := range requests {
        result := process(req.Data)
        req.Response <- result
    }
}

// Client
response := make(chan string)
requests <- Request{
    Data:     "hello",
    Response: response,
}
result := <-response
```

### 方式2: 每个请求一个 Channel

```
Client 1 ──> req1 ──┐
                    ├──> Router ──> Workers
Client 2 ──> req2 ──┘                │
                                     │
Client 1 <──────────────────────────┘
     (通过 req1.Response)
```

---

## 信号量模式（限制并发）

```
Semaphore [3 个槽位]
┌─────────────────────┐
│  [✓] [✓] [✓]        │
└──────┬──────────────┘
       │
    请求获取
       │
┌──────▼──────────────┐
│  [✓] [✓] [✓]        │  <- 3个任务正在执行
└──────┬──────────────┘
       │
    新请求被阻塞 ❌
       │
┌──────▼──────────────┐
│  [✓] [✓] [_]        │  <- 一个任务完成
└──────┬──────────────┘
       │
    新请求被允许 ✓
```

**代码：**
```go
// 创建信号量（最多3个并发）
semaphore := make(chan struct{}, 3)

for i := 0; i < 10; i++ {
    semaphore <- struct{}{}  // 获取
    go func() {
        defer func() { <-semaphore }()  // 释放
        doWork()
    }()
}
```

---

## 生产者-消费者模式

### 单生产者-单消费者

```
Producer                Consumer
   │                       │
   ├─> [1] ──────────────> │
   ├─> [2] ──────────────> │
   ├─> [3] ──────────────> │
   │                       │
close(ch)                  │
   │                       │
   └─────────────────────> │ (range 自动退出)
```

### 多生产者-多消费者

```
Producer 1 ──┐            ┌──> Consumer 1
             │            │
Producer 2 ──┼─> Queue ───┼──> Consumer 2
             │            │
Producer 3 ──┘            └──> Consumer 3
```

**同步关闭：**
```go
var wg sync.WaitGroup

// 启动生产者
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        produce(ch)
    }()
}

// 等待所有生产者完成，然后关闭
go func() {
    wg.Wait()
    close(ch)
}()

// 消费者
for v := range ch {
    process(v)
}
```

---

## 令牌桶限流器

```
时间轴 ─────────────────────────────────>

t0: [🪙🪙🪙🪙🪙] (满，容量5)
       ↓
t1: [🪙🪙🪙🪙_] (请求1消耗1个)
       ↓
t2: [🪙🪙🪙__] (请求2消耗1个)
       ↓ +1 (补充)
t3: [🪙🪙🪙🪙_]
       ↓
t4: [🪙🪙🪙__] (请求3消耗1个)
       ↓ +1
t5: [🪙🪙🪙🪙_]
```

**代码：**
```go
type TokenBucket struct {
    tokens chan struct{}
}

func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
    tb := &TokenBucket{
        tokens: make(chan struct{}, capacity),
    }
    
    // 初始填满
    for i := 0; i < capacity; i++ {
        tb.tokens <- struct{}{}
    }
    
    // 定期补充
    go func() {
        ticker := time.NewTicker(rate)
        for range ticker.C {
            select {
            case tb.tokens <- struct{}{}:
            default:
            }
        }
    }()
    
    return tb
}

func (tb *TokenBucket) Take() {
    <-tb.tokens
}
```

---

## 选择合适的模式

| 场景 | 推荐模式 | 示例代码 |
|------|---------|---------|
| 数据生成 | Generator | `examples/01_basic_channel.go` |
| 数据转换 | Pipeline | `examples/06_pipeline.go` |
| 并发处理 | Worker Pool | `examples/05_worker_pool.go` |
| 负载分散 | Fan-Out | `examples/07_fan_out_fan_in.go` |
| 结果聚合 | Fan-In | `examples/07_fan_out_fan_in.go` |
| 限制并发 | Semaphore | `examples/02_buffered_channel.go` |
| 超时控制 | Select+Timer | `examples/04_select_statement.go` |
| 取消操作 | Context | `examples/08_context_cancel.go` |
| 请求响应 | Request-Response | `examples/09_web_scraper.go` |
| 限流 | Token Bucket | `examples/10_rate_limiter.go` |

---

## 反模式（应该避免）

### ❌ 不要在接收方关闭 Channel

```
Producer ──> Channel ──> Consumer
                           │
                        close(ch) ❌
                           │
                        panic!
```

### ❌ 不要向已关闭的 Channel 发送

```
Producer                 Channel
   │                        │
   │                     closed
   │                        │
   ├─> send ──────────────> ❌ panic!
```

### ❌ 不要重复关闭 Channel

```
close(ch)
close(ch)  ❌ panic: close of closed channel
```

### ❌ 避免 Goroutine 泄漏

```
func leak() {
    ch := make(chan int)
    go func() {
        <-ch  // 永远等待 ❌
    }()
    // 函数返回，但 goroutine 泄漏
}
```

**正确做法：**
```go
func noLeak() {
    ch := make(chan int)
    done := make(chan struct{})
    
    go func() {
        select {
        case <-ch:
        case <-done:
            return  // 可以退出
        }
    }()
    
    close(done)  // 清理
}
```

---

## 总结

掌握这些模式可以帮助你：

1. ✅ 写出清晰的并发代码
2. ✅ 避免常见的并发陷阱
3. ✅ 提高程序性能
4. ✅ 简化复杂的同步逻辑

**学习路径：**
1. 先掌握基础通信模式
2. 理解 Pipeline 和 Worker Pool
3. 学习 Fan-Out/Fan-In
4. 掌握超时和取消机制
5. 在实际项目中应用

**参考示例代码：**
- `examples/` 目录包含所有模式的完整实现
- 每个文件都是独立可运行的示例
- 建议按照编号顺序学习
