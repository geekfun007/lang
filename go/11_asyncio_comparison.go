package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ============================================================
// 示例1: 基础异步操作
// Python: async def + await
// Go: goroutine + channel/WaitGroup
// ============================================================

func example1_basic_async() {
	fmt.Println("\n=== 示例1: 基础异步操作 ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
async def hello():
    print("Hello")
    await asyncio.sleep(1)
    print("World")

asyncio.run(hello())
`)
	
	fmt.Println("Go 等价代码 (同步):")
	func() {
		fmt.Println("Hello")
		time.Sleep(1 * time.Second)
		fmt.Println("World")
	}()
	
	fmt.Println("\nGo 等价代码 (异步):")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello (async)")
		time.Sleep(1 * time.Second)
		fmt.Println("World (async)")
	}()
	wg.Wait()
}

// ============================================================
// 示例2: asyncio.gather() - 并发执行多个任务
// ============================================================

func example2_gather() {
	fmt.Println("\n=== 示例2: asyncio.gather() ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
async def task(name, delay):
    await asyncio.sleep(delay)
    return f"Result from {name}"

results = await asyncio.gather(
    task("A", 1),
    task("B", 2),
    task("C", 1.5)
)
`)
	
	fmt.Println("Go 等价代码:")
	
	type TaskResult struct {
		name   string
		result string
	}
	
	task := func(name string, delay time.Duration, results chan<- TaskResult, wg *sync.WaitGroup) {
		defer wg.Done()
		time.Sleep(delay)
		results <- TaskResult{name: name, result: fmt.Sprintf("Result from %s", name)}
	}
	
	results := make(chan TaskResult, 3)
	var wg sync.WaitGroup
	
	start := time.Now()
	
	wg.Add(3)
	go task("A", 1*time.Second, results, &wg)
	go task("B", 2*time.Second, results, &wg)
	go task("C", 1500*time.Millisecond, results, &wg)
	
	go func() {
		wg.Wait()
		close(results)
	}()
	
	var allResults []string
	for result := range results {
		allResults = append(allResults, result.result)
		fmt.Printf("  %s\n", result.result)
	}
	
	fmt.Printf("总耗时: %v (并发执行，取最长时间)\n", time.Since(start))
}

// ============================================================
// 示例3: asyncio.wait_for() - 超时控制
// ============================================================

func example3_timeout() {
	fmt.Println("\n=== 示例3: asyncio.wait_for() ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
try:
    result = await asyncio.wait_for(long_task(), timeout=2.0)
except asyncio.TimeoutError:
    print("Timeout!")
`)
	
	fmt.Println("Go 等价代码 (方式1: select + time.After):")
	
	longTask := func() <-chan string {
		result := make(chan string)
		go func() {
			time.Sleep(3 * time.Second)
			result <- "Done"
		}()
		return result
	}
	
	select {
	case result := <-longTask():
		fmt.Println("  收到结果:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("  超时！")
	}
	
	fmt.Println("\nGo 等价代码 (方式2: Context):")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	result := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		select {
		case result <- "Done":
		case <-ctx.Done():
			return
		}
	}()
	
	select {
	case r := <-result:
		fmt.Println("  收到结果:", r)
	case <-ctx.Done():
		fmt.Println("  超时！")
	}
}

// ============================================================
// 示例4: asyncio.Queue - 异步队列
// ============================================================

func example4_queue() {
	fmt.Println("\n=== 示例4: asyncio.Queue ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
async def producer(queue):
    for i in range(5):
        await queue.put(i)

async def consumer(queue):
    while True:
        item = await queue.get()
        print(f"Consumed {item}")
`)
	
	fmt.Println("Go 等价代码:")
	
	producer := func(queue chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			queue <- i
			fmt.Printf("  Produced: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		close(queue)
	}
	
	consumer := func(queue <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for item := range queue {
			fmt.Printf("  Consumed: %d\n", item)
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	queue := make(chan int, 10)
	var wg sync.WaitGroup
	
	wg.Add(2)
	go producer(queue, &wg)
	go consumer(queue, &wg)
	
	wg.Wait()
}

// ============================================================
// 示例5: asyncio.Semaphore - 限制并发数
// ============================================================

func example5_semaphore() {
	fmt.Println("\n=== 示例5: asyncio.Semaphore ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
semaphore = asyncio.Semaphore(3)

async def download(url):
    async with semaphore:
        await asyncio.sleep(2)
        return f"Downloaded {url}"

await asyncio.gather(*[download(f"url{i}") for i in range(10)])
`)
	
	fmt.Println("Go 等价代码:")
	
	download := func(url string, semaphore chan struct{}, results chan<- string, wg *sync.WaitGroup) {
		defer wg.Done()
		
		// 获取信号量
		semaphore <- struct{}{}
		defer func() { <-semaphore }()
		
		fmt.Printf("  开始下载 %s\n", url)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("  完成下载 %s\n", url)
		
		results <- fmt.Sprintf("Downloaded %s", url)
	}
	
	// 最多3个并发
	semaphore := make(chan struct{}, 3)
	results := make(chan string, 10)
	var wg sync.WaitGroup
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go download(fmt.Sprintf("url%d", i), semaphore, results, &wg)
	}
	
	go func() {
		wg.Wait()
		close(results)
	}()
	
	count := 0
	for range results {
		count++
	}
	
	fmt.Printf("  总共下载了 %d 个文件\n", count)
}

// ============================================================
// 示例6: asyncio.create_task() - 创建后台任务
// ============================================================

func example6_create_task() {
	fmt.Println("\n=== 示例6: asyncio.create_task() ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
async def background_task():
    await asyncio.sleep(2)
    print("Background task done")

task = asyncio.create_task(background_task())
# 继续其他工作
await task  # 等待任务完成
`)
	
	fmt.Println("Go 等价代码:")
	
	done := make(chan bool)
	
	// 创建后台任务
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("  Background task done")
		done <- true
	}()
	
	fmt.Println("  主程序继续其他工作...")
	time.Sleep(1 * time.Second)
	fmt.Println("  主程序工作完成，等待后台任务...")
	
	<-done  // 等待任务完成
	fmt.Println("  所有任务完成")
}

// ============================================================
// 示例7: 异步迭代器
// ============================================================

func example7_async_iterator() {
	fmt.Println("\n=== 示例7: 异步迭代器 ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
class AsyncIterator:
    async def __anext__(self):
        await asyncio.sleep(0.5)
        return value

async for item in AsyncIterator():
    print(item)
`)
	
	fmt.Println("Go 等价代码:")
	
	asyncIterator := func(n int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 1; i <= n; i++ {
				time.Sleep(500 * time.Millisecond)
				ch <- i
			}
		}()
		return ch
	}
	
	for item := range asyncIterator(5) {
		fmt.Printf("  收到: %d\n", item)
	}
}

// ============================================================
// 示例8: 取消任务
// ============================================================

func example8_cancel() {
	fmt.Println("\n=== 示例8: 取消任务 ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
task = asyncio.create_task(long_running())
await asyncio.sleep(2)
task.cancel()  # 取消任务
`)
	
	fmt.Println("Go 等价代码:")
	
	ctx, cancel := context.WithCancel(context.Background())
	
	// 长时间运行的任务
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		count := 0
		for {
			select {
			case <-ticker.C:
				count++
				fmt.Printf("  任务运行中... %d\n", count)
			case <-ctx.Done():
				fmt.Println("  任务被取消")
				return
			}
		}
	}()
	
	time.Sleep(2 * time.Second)
	fmt.Println("  发送取消信号...")
	cancel()
	
	time.Sleep(500 * time.Millisecond)
}

// ============================================================
// 示例9: 异步上下文管理器
// ============================================================

func example9_context_manager() {
	fmt.Println("\n=== 示例9: 异步上下文管理器 ===")
	fmt.Println("Python 代码:")
	fmt.Println(`
async with resource:
    await do_work()
# 自动清理资源
`)
	
	fmt.Println("Go 等价代码:")
	
	type Resource struct {
		name string
	}
	
	acquire := func(name string) *Resource {
		fmt.Printf("  获取资源: %s\n", name)
		time.Sleep(100 * time.Millisecond)
		return &Resource{name: name}
	}
	
	release := func(r *Resource) {
		fmt.Printf("  释放资源: %s\n", r.name)
		time.Sleep(100 * time.Millisecond)
	}
	
	// 使用 defer 实现自动清理
	func() {
		resource := acquire("Database")
		defer release(resource)
		
		fmt.Println("  执行工作...")
		time.Sleep(500 * time.Millisecond)
	}()
	
	fmt.Println("  资源已清理")
}

// ============================================================
// 示例10: 复杂场景 - Web 爬虫
// ============================================================

func example10_web_scraper() {
	fmt.Println("\n=== 示例10: Web 爬虫 (复杂场景) ===")
	fmt.Println("Python asyncio 实现:")
	fmt.Println(`
async def fetch(url, session):
    async with session.get(url) as response:
        return await response.text()

async def main():
    async with aiohttp.ClientSession() as session:
        tasks = [fetch(url, session) for url in urls]
        results = await asyncio.gather(*tasks)
`)
	
	fmt.Println("\nGo 实现:")
	
	type Page struct {
		URL     string
		Content string
	}
	
	// 模拟抓取网页
	fetch := func(url string, semaphore chan struct{}) Page {
		semaphore <- struct{}{}
		defer func() { <-semaphore }()
		
		fmt.Printf("  抓取: %s\n", url)
		time.Sleep(time.Duration(100+len(url)*10) * time.Millisecond)
		
		return Page{
			URL:     url,
			Content: fmt.Sprintf("Content from %s", url),
		}
	}
	
	urls := []string{
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
		"https://example.com/page4",
		"https://example.com/page5",
	}
	
	// 限制并发数为3
	semaphore := make(chan struct{}, 3)
	results := make(chan Page, len(urls))
	var wg sync.WaitGroup
	
	start := time.Now()
	
	// 并发抓取
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			page := fetch(u, semaphore)
			results <- page
		}(url)
	}
	
	// 等待完成
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// 收集结果
	var pages []Page
	for page := range results {
		pages = append(pages, page)
		fmt.Printf("  完成: %s\n", page.URL)
	}
	
	fmt.Printf("  总耗时: %v\n", time.Since(start))
	fmt.Printf("  抓取了 %d 个页面\n", len(pages))
}

// ============================================================
// Main
// ============================================================

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                            ║")
	fmt.Println("║        Python asyncio vs Go 并发编程对比示例               ║")
	fmt.Println("║                                                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	
	example1_basic_async()
	example2_gather()
	example3_timeout()
	example4_queue()
	example5_semaphore()
	example6_create_task()
	example7_async_iterator()
	example8_cancel()
	example9_context_manager()
	example10_web_scraper()
	
	fmt.Println("\n╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║  所有示例运行完成！                                         ║")
	fmt.Println("║  查看 PYTHON_ASYNCIO_COMPARISON.md 了解更多详情            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}
