# Goroutine 核心原理详解

## 1. Goroutine 是什么？

Goroutine 是 Go 语言的轻量级线程，由 Go 运行时管理。它是 Go 并发编程的基础。

> "Goroutines are lightweight threads managed by the Go runtime."

### 1.1 Goroutine vs 线程

| 特性 | Goroutine | 操作系统线程 |
|------|-----------|-------------|
| **内存占用** | 初始 2KB 栈 | 1-2 MB 栈 |
| **创建成本** | ~几微秒 | ~几毫秒 |
| **调度** | Go 运行时（用户态） | 操作系统（内核态） |
| **数量** | 数十万个 | 数千个 |
| **上下文切换** | ~200ns | ~1-2μs |

**示例：**
```go
// 创建 Goroutine 非常简单
go func() {
    fmt.Println("Hello from goroutine!")
}()

// 可以轻松创建大量 goroutine
for i := 0; i < 100000; i++ {
    go func(id int) {
        // 处理任务
    }(i)
}
```

## 2. GMP 调度模型

Go 使用 GMP 模型实现 M:N 调度（M个 goroutine 调度到 N 个 OS 线程）。

### 2.1 核心组件

```
G (Goroutine)
    ↓
M (Machine/OS Thread)
    ↓
P (Processor)
```

#### G - Goroutine

```go
type g struct {
    stack       stack   // 栈内存
    stackguard0 uintptr // 栈保护
    m           *m      // 当前运行的 M
    sched       gobuf   // 调度信息
    atomicstatus uint32 // 状态
    goid         int64  // goroutine ID
    // ... 更多字段
}
```

**Goroutine 状态：**
- `_Gidle`: 刚创建
- `_Grunnable`: 可运行，在队列中
- `_Grunning`: 正在运行
- `_Gsyscall`: 系统调用中
- `_Gwaiting`: 等待中（阻塞）
- `_Gdead`: 已终止

#### M - Machine (OS 线程)

```go
type m struct {
    g0      *g       // 用于调度的特殊 goroutine
    curg    *g       // 当前运行的 goroutine
    p       puintptr // 绑定的 P
    nextp   puintptr // 下一个要绑定的 P
    spinning bool    // 是否处于自旋状态
    // ... 更多字段
}
```

**特点：**
- 对应一个操作系统线程
- 数量由 runtime.GOMAXPROCS 控制
- 可以动态创建和销毁

#### P - Processor (逻辑处理器)

```go
type p struct {
    id          int32
    status      uint32     // P 的状态
    m           muintptr   // 绑定的 M
    runqhead    uint32     // 本地队列头
    runqtail    uint32     // 本地队列尾
    runq        [256]guintptr // 本地运行队列
    runnext     guintptr   // 下一个要运行的 G
    // ... 更多字段
}
```

**特点：**
- 数量 = GOMAXPROCS（默认 = CPU 核心数）
- 维护本地 goroutine 队列
- 减少锁竞争

### 2.2 GMP 工作流程

```
全局队列 (Global Queue)
    ↓
┌────────────┐  ┌────────────┐  ┌────────────┐
│ P1         │  │ P2         │  │ P3         │
│ ┌────────┐ │  │ ┌────────┐ │  │ ┌────────┐ │
│ │ 本地队列 │ │  │ │ 本地队列 │ │  │ │ 本地队列 │ │
│ │  G1  G2 │ │  │ │  G3  G4 │ │  │ │  G5  G6 │ │
│ └────────┘ │  │ └────────┘ │  │ └────────┘ │
│     ↓      │  │     ↓      │  │     ↓      │
│  ┌───┐    │  │  ┌───┐    │  │  ┌───┐    │
│  │ M1│    │  │  │ M2│    │  │  │ M3│    │
│  └───┘    │  │  └───┘    │  │  └───┘    │
└────────────┘  └────────────┘  └────────────┘
     ↓               ↓               ↓
  OS Thread      OS Thread      OS Thread
```

**调度流程：**

1. **创建 Goroutine**
   ```
   go func() → 创建 G → 加入 P 的本地队列
   ```

2. **执行 Goroutine**
   ```
   M 从 P 的本地队列获取 G → 执行 → 完成或阻塞
   ```

3. **Work Stealing（工作窃取）**
   ```
   P1 本地队列为空 → 从 P2 偷取一半 G
   ```

4. **Hand Off（移交）**
   ```
   G 发起系统调用 → M 阻塞 → P 移交给新的 M
   ```

## 3. Goroutine 的生命周期

### 3.1 创建

```go
// 源码简化版
func newproc(fn *funcval) {
    // 1. 从 P 的缓存或全局缓存获取 G
    gp := gfget(_p_)
    if gp == nil {
        // 创建新的 G，栈大小 2KB
        gp = malg(_StackMin)
    }
    
    // 2. 初始化 G
    gp.sched.sp = sp
    gp.sched.pc = pc
    gp.sched.g = gp
    
    // 3. 设置为 runnable 状态
    casgstatus(gp, _Gdead, _Grunnable)
    
    // 4. 加入 P 的本地队列
    runqput(_p_, gp, true)
}
```

### 3.2 调度

**调度时机：**
1. `go` 语句创建新 goroutine
2. GC 垃圾回收
3. 系统调用
4. 同步操作（channel、mutex 等）
5. 主动调用 `runtime.Gosched()`

**调度策略：**
```go
// 获取下一个要运行的 goroutine
func findrunnable() *g {
    // 1. 每 61 次检查一次全局队列（防止饥饿）
    if _g_.m.p.ptr().schedtick%61 == 0 {
        gp := globrunqget(_g_.m.p.ptr(), 1)
        if gp != nil {
            return gp
        }
    }
    
    // 2. 从本地队列获取
    gp := runqget(_g_.m.p.ptr())
    if gp != nil {
        return gp
    }
    
    // 3. 从全局队列获取
    gp = globrunqget(_g_.m.p.ptr(), 0)
    if gp != nil {
        return gp
    }
    
    // 4. 从其他 P 偷取
    gp = stealWork()
    if gp != nil {
        return gp
    }
    
    // 5. 都没有，休眠
    stopm()
}
```

### 3.3 销毁

```go
// Goroutine 执行完毕后
func goexit() {
    // 1. 执行延迟函数
    rundefers()
    
    // 2. 清理 goroutine
    gp := getg()
    casgstatus(gp, _Grunning, _Gdead)
    
    // 3. 放入缓存池（复用）
    gfput(_p_, gp)
    
    // 4. 调度下一个 goroutine
    schedule()
}
```

## 4. 栈管理

### 4.1 动态栈

Go 使用**分段栈**（Segmented Stacks）策略（Go 1.3 之前）和**连续栈**（Contiguous Stacks）策略（Go 1.3+）。

**连续栈工作原理：**

```
初始栈 (2KB)              栈增长 (4KB)              栈增长 (8KB)
┌──────────┐              ┌──────────┐              ┌──────────┐
│          │              │          │              │          │
│  Stack   │──扩容──>     │  Stack   │──扩容──>     │  Stack   │
│  2KB     │              │  4KB     │              │  8KB     │
│          │              │          │              │          │
└──────────┘              └──────────┘              └──────────┘
                          (复制数据)                (复制数据)
```

**栈扩容流程：**
```go
// 函数序言检查栈空间
func morestack() {
    // 1. 检查是否需要扩容
    if sp < stackguard {
        // 2. 分配新栈（通常是当前的 2 倍）
        newstack := allocstack(oldsize * 2)
        
        // 3. 复制旧栈数据
        copystack(gp, newstack)
        
        // 4. 更新指针
        adjustpointers(gp, &adjinfo)
        
        // 5. 释放旧栈
        freestack(oldstack)
    }
}
```

### 4.2 栈缩容

```go
// GC 时检查栈是否过大
func shrinkstack(gp *g) {
    oldsize := gp.stack.hi - gp.stack.lo
    newsize := oldsize / 2
    
    // 使用不到 1/4，缩容
    if used < oldsize/4 {
        copystack(gp, newsize)
    }
}
```

## 5. 调度器的优化

### 5.1 本地队列 (Local Run Queue)

**目的：** 减少锁竞争

```go
// P 的本地队列
type p struct {
    runq [256]guintptr  // 本地队列
    runqhead uint32
    runqtail uint32
}

// 入队（无锁，只有当前 M 访问）
func runqput(p *p, gp *g, next bool) {
    if next {
        // 优先执行（放在 runnext）
        oldnext := p.runnext
        p.runnext.set(gp)
        if oldnext == 0 {
            return
        }
        gp = oldnext.ptr()
    }
    
    // 放入本地队列
retry:
    h := atomic.LoadAcq(&p.runqhead)
    t := p.runqtail
    if t-h < uint32(len(p.runq)) {
        p.runq[t%uint32(len(p.runq))].set(gp)
        atomic.StoreRel(&p.runqtail, t+1)
        return
    }
    
    // 本地队列满，转移一半到全局队列
    if !runqputslow(p, gp, h, t) {
        goto retry
    }
}
```

### 5.2 Work Stealing（工作窃取）

**目的：** 负载均衡

```go
func stealWork(now int64) *g {
    pp := getg().m.p.ptr()
    
    // 随机选择一个 P 开始
    offset := fastrand() % uint32(gomaxprocs)
    
    for i := 0; i < gomaxprocs; i++ {
        p2 := allp[(offset+i)%gomaxprocs]
        if p2 == pp {
            continue
        }
        
        // 偷取一半 goroutine
        gp := runqsteal(pp, p2, true)
        if gp != nil {
            return gp
        }
    }
    
    return nil
}
```

### 5.3 Spinning Thread（自旋线程）

**目的：** 减少延迟

```go
// M 在没有 G 可运行时，会自旋一段时间
func findrunnable() *g {
    // ...
    
    // 自旋等待
    if _g_.m.spinning {
        // 快速路径：检查本地队列
        if gp := runqget(_g_.m.p.ptr()); gp != nil {
            return gp
        }
    }
    
    // ...
}
```

**自旋条件：**
- 至少有一个空闲的 P
- 没有空闲的 M
- 自旋的 M 数量 < GOMAXPROCS

### 5.4 Hand Off 机制

当 Goroutine 进行系统调用时：

```
阻塞前:
┌─────┐
│ P1  │
├─────┤        ┌─────┐
│ M1  │───────>│ G1  │ (发起系统调用)
└─────┘        └─────┘

阻塞时 (Hand Off):
┌─────┐        ┌─────┐
│ P1  │───┐    │ G1  │ (系统调用中)
├─────┤   │    └─────┘
│ M2  │<──┘       ↓
└─────┘        ┌─────┐
   ↓           │ M1  │ (阻塞)
┌─────┐        └─────┘
│ G2  │ (继续执行其他 goroutine)
└─────┘

系统调用返回:
               ┌─────┐
               │ G1  │ (系统调用完成)
               └─────┘
                  ↓
            尝试获取 P 或进入全局队列
```

## 6. GOMAXPROCS 的影响

```go
// 设置 P 的数量
runtime.GOMAXPROCS(n)
```

**影响：**
- 决定并行执行的 goroutine 数量
- 影响 CPU 利用率
- 影响调度开销

**最佳实践：**
```go
// CPU 密集型：设置为 CPU 核心数
runtime.GOMAXPROCS(runtime.NumCPU())

// I/O 密集型：可以设置更大
runtime.GOMAXPROCS(runtime.NumCPU() * 2)
```

## 7. Goroutine 泄漏

### 7.1 常见原因

**1. Channel 永久阻塞**
```go
// 泄漏
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // 永远等待
    }()
    // 函数返回，但 goroutine 泄漏
}
```

**2. Mutex 死锁**
```go
// 泄漏
var mu sync.Mutex
func leak() {
    mu.Lock()
    go func() {
        mu.Lock()  // 永远等待
        defer mu.Unlock()
    }()
    // 忘记 Unlock
}
```

**3. WaitGroup 计数错误**
```go
// 泄漏
func leak() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        // 忘记 wg.Done()
    }()
    wg.Wait()  // 永远等待
}
```

### 7.2 检测方法

```go
// 1. 运行时检测
before := runtime.NumGoroutine()
// 执行代码
after := runtime.NumGoroutine()
if after > before {
    fmt.Printf("可能泄漏 %d goroutines\n", after-before)
}

// 2. 使用 pprof
import _ "net/http/pprof"
// 访问 http://localhost:6060/debug/pprof/goroutine

// 3. 使用 runtime/trace
import "runtime/trace"
trace.Start(os.Stdout)
defer trace.Stop()
```

## 8. 性能优化

### 8.1 减少 Goroutine 创建

```go
// 不好：频繁创建
for _, item := range items {
    go process(item)
}

// 更好：使用 Worker Pool
pool := make(chan struct{}, 10)
for _, item := range items {
    pool <- struct{}{}
    go func(item Item) {
        defer func() { <-pool }()
        process(item)
    }(item)
}
```

### 8.2 避免 Goroutine 泄漏

```go
// 使用 Context
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // 工作
        }
    }
}
```

### 8.3 合理设置 GOMAXPROCS

```go
// 根据工作负载调整
if cpuBound {
    runtime.GOMAXPROCS(runtime.NumCPU())
} else {
    runtime.GOMAXPROCS(runtime.NumCPU() * 2)
}
```

## 9. 调试工具

### 9.1 runtime 包

```go
// 获取 goroutine 数量
n := runtime.NumGoroutine()

// 获取调度信息
var stats runtime.MemStats
runtime.ReadMemStats(&stats)

// 打印所有 goroutine 的栈
buf := make([]byte, 1<<16)
n := runtime.Stack(buf, true)
```

### 9.2 trace 工具

```go
import "runtime/trace"

f, _ := os.Create("trace.out")
trace.Start(f)
defer trace.Stop()

// 执行代码...

// 查看: go tool trace trace.out
```

### 9.3 pprof

```bash
# Goroutine profile
go tool pprof http://localhost:6060/debug/pprof/goroutine

# 查看调度延迟
GODEBUG=schedtrace=1000 ./myprogram
```

## 10. 总结

### Goroutine 的核心特性

- ✅ **轻量级**: 2KB 初始栈，支持数十万个并发
- ✅ **高效调度**: GMP 模型，M:N 调度
- ✅ **动态栈**: 自动增长和缩容
- ✅ **Work Stealing**: 自动负载均衡
- ✅ **低延迟**: 用户态调度，上下文切换快

### 关键设计决策

| 设计 | 目的 | 实现 |
|------|------|------|
| GMP 模型 | 高效调度 | M:N 多路复用 |
| 本地队列 | 减少锁竞争 | 每个 P 独立队列 |
| Work Stealing | 负载均衡 | 空闲时窃取其他 P 的任务 |
| Hand Off | 避免阻塞 | 系统调用时移交 P |
| 动态栈 | 内存高效 | 按需增长/缩容 |

### 最佳实践

1. ✅ 使用 Worker Pool 限制并发数
2. ✅ 使用 Context 管理生命周期
3. ✅ 避免 Goroutine 泄漏
4. ✅ 合理设置 GOMAXPROCS
5. ✅ 使用工具监控和调试

---

**参考资料：**
- [Go Runtime Scheduler](https://golang.org/src/runtime/proc.go)
- [GMP Model Design Doc](https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw)
- [Go Scheduler Analysis](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
