package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// URL 代表一个待抓取的网页
type URL struct {
	Address string
	Depth   int
}

// PageContent 代表抓取的页面内容
type PageContent struct {
	URL     string
	Content string
	Links   []string
	Error   error
}

// 示例1: 简单的并发网页抓取器
func example1_simple_scraper() {
	fmt.Println("\n=== 示例1: 简单并发抓取器 ===")
	
	// 模拟抓取网页
	fetchPage := func(url string) (PageContent, error) {
		// 模拟网络延迟
		time.Sleep(time.Duration(100+rand.Intn(300)) * time.Millisecond)
		
		// 模拟10%的失败率
		if rand.Float32() < 0.1 {
			return PageContent{}, fmt.Errorf("抓取失败: %s", url)
		}
		
		return PageContent{
			URL:     url,
			Content: fmt.Sprintf("内容来自 %s", url),
			Links:   []string{},
		}, nil
	}
	
	urls := []string{
		"https://example.com/page1",
		"https://example.com/page2",
		"https://example.com/page3",
		"https://example.com/page4",
		"https://example.com/page5",
	}
	
	results := make(chan PageContent, len(urls))
	
	// 并发抓取
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fmt.Printf("开始抓取: %s\n", u)
			page, err := fetchPage(u)
			if err != nil {
				results <- PageContent{URL: u, Error: err}
			} else {
				results <- page
			}
		}(url)
	}
	
	// 等待所有抓取完成
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// 收集结果
	successCount := 0
	failCount := 0
	for page := range results {
		if page.Error != nil {
			fmt.Printf("❌ %s: %v\n", page.URL, page.Error)
			failCount++
		} else {
			fmt.Printf("✅ %s: %s\n", page.URL, page.Content)
			successCount++
		}
	}
	
	fmt.Printf("\n统计: 成功 %d, 失败 %d\n", successCount, failCount)
}

// 示例2: 带限流的抓取器
func example2_rate_limited_scraper() {
	fmt.Println("\n=== 示例2: 带限流的抓取器 ===")
	
	// 限流器：每秒最多2个请求
	type RateLimiter struct {
		tokens chan struct{}
	}
	
	NewRateLimiter := func(requestsPerSecond int) *RateLimiter {
		rl := &RateLimiter{
			tokens: make(chan struct{}, requestsPerSecond),
		}
		
		// 定期补充令牌
		go func() {
			ticker := time.NewTicker(time.Second / time.Duration(requestsPerSecond))
			defer ticker.Stop()
			for range ticker.C {
				select {
				case rl.tokens <- struct{}{}:
				default:
				}
			}
		}()
		
		return rl
	}
	
	fetchWithLimit := func(limiter *RateLimiter, url string) PageContent {
		<-limiter.tokens // 获取令牌
		
		fmt.Printf("[%s] 抓取: %s\n", time.Now().Format("15:04:05"), url)
		time.Sleep(100 * time.Millisecond)
		
		return PageContent{
			URL:     url,
			Content: fmt.Sprintf("内容: %s", url),
		}
	}
	
	limiter := NewRateLimiter(2)
	urls := []string{"page1", "page2", "page3", "page4", "page5", "page6"}
	
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			fetchWithLimit(limiter, u)
		}(url)
	}
	
	wg.Wait()
}

// 示例3: 深度优先爬虫（递归抓取链接）
func example3_depth_first_crawler() {
	fmt.Println("\n=== 示例3: 深度优先爬虫 ===")
	
	// 模拟抓取页面（返回链接）
	fetchPageWithLinks := func(url string) PageContent {
		time.Sleep(100 * time.Millisecond)
		
		// 模拟生成子链接
		var links []string
		numLinks := rand.Intn(3)
		for i := 0; i < numLinks; i++ {
			links = append(links, fmt.Sprintf("%s/link%d", url, i+1))
		}
		
		return PageContent{
			URL:     url,
			Content: fmt.Sprintf("页面: %s", url),
			Links:   links,
		}
	}
	
	// 爬虫
	crawl := func(url string, maxDepth int) {
		visited := make(map[string]bool)
		var mu sync.Mutex
		
		var crawlRecursive func(string, int)
		crawlRecursive = func(u string, depth int) {
			if depth > maxDepth {
				return
			}
			
			// 检查是否已访问
			mu.Lock()
			if visited[u] {
				mu.Unlock()
				return
			}
			visited[u] = true
			mu.Unlock()
			
			// 抓取页面
			fmt.Printf("深度 %d: %s\n", depth, u)
			page := fetchPageWithLinks(u)
			
			// 递归抓取链接
			var wg sync.WaitGroup
			for _, link := range page.Links {
				wg.Add(1)
				go func(l string) {
					defer wg.Done()
					crawlRecursive(l, depth+1)
				}(link)
			}
			wg.Wait()
		}
		
		crawlRecursive(url, 0)
	}
	
	crawl("https://example.com", 2)
}

// 示例4: 广度优先爬虫（使用队列）
func example4_breadth_first_crawler() {
	fmt.Println("\n=== 示例4: 广度优先爬虫 ===")
	
	fetchPageWithLinks := func(url string) PageContent {
		time.Sleep(100 * time.Millisecond)
		
		var links []string
		numLinks := rand.Intn(2) + 1
		for i := 0; i < numLinks; i++ {
			links = append(links, fmt.Sprintf("%s/sub%d", url, i+1))
		}
		
		return PageContent{
			URL:     url,
			Content: fmt.Sprintf("内容: %s", url),
			Links:   links,
		}
	}
	
	crawlBFS := func(startURL string, maxDepth int) {
		visited := make(map[string]bool)
		queue := make(chan URL, 100)
		results := make(chan PageContent, 100)
		
		// 启动 worker pool
		const numWorkers = 3
		var wg sync.WaitGroup
		
		for w := 1; w <= numWorkers; w++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				for url := range queue {
					if url.Depth > maxDepth {
						continue
					}
					
					fmt.Printf("Worker %d 抓取深度 %d: %s\n", id, url.Depth, url.Address)
					page := fetchPageWithLinks(url.Address)
					results <- page
					
					// 将新链接加入队列
					for _, link := range page.Links {
						queue <- URL{Address: link, Depth: url.Depth + 1}
					}
				}
			}(w)
		}
		
		// 发送起始 URL
		visited[startURL] = true
		queue <- URL{Address: startURL, Depth: 0}
		
		// 收集结果
		go func() {
			count := 0
			maxCount := 10 // 限制抓取数量
			for page := range results {
				fmt.Printf("收到: %s (包含 %d 个链接)\n", page.URL, len(page.Links))
				count++
				
				if count >= maxCount {
					close(queue)
					return
				}
			}
		}()
		
		wg.Wait()
		close(results)
	}
	
	crawlBFS("https://example.com", 3)
}

// 示例5: 带超时和重试的抓取器
func example5_scraper_with_retry() {
	fmt.Println("\n=== 示例5: 带超时和重试的抓取器 ===")
	
	fetchWithRetry := func(ctx context.Context, url string, maxRetries int) (PageContent, error) {
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("尝试 %d/%d: %s\n", attempt, maxRetries, url)
			
			// 每次尝试有500ms超时
			attemptCtx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
			
			resultCh := make(chan PageContent, 1)
			errCh := make(chan error, 1)
			
			go func() {
				// 模拟抓取
				delay := time.Duration(rand.Intn(800)) * time.Millisecond
				time.Sleep(delay)
				
				// 模拟失败
				if rand.Float32() < 0.4 {
					errCh <- fmt.Errorf("抓取失败")
					return
				}
				
				resultCh <- PageContent{
					URL:     url,
					Content: fmt.Sprintf("内容: %s", url),
				}
			}()
			
			select {
			case page := <-resultCh:
				cancel()
				fmt.Printf("✅ 成功: %s\n", url)
				return page, nil
			case err := <-errCh:
				cancel()
				fmt.Printf("❌ 失败: %v\n", err)
				if attempt < maxRetries {
					time.Sleep(200 * time.Millisecond)
				}
			case <-attemptCtx.Done():
				cancel()
				fmt.Printf("⏱️  超时\n")
				if attempt < maxRetries {
					time.Sleep(200 * time.Millisecond)
				}
			case <-ctx.Done():
				cancel()
				return PageContent{}, ctx.Err()
			}
		}
		
		return PageContent{}, fmt.Errorf("达到最大重试次数")
	}
	
	ctx := context.Background()
	urls := []string{"page1", "page2", "page3"}
	
	for _, url := range urls {
		_, err := fetchWithRetry(ctx, url, 3)
		if err != nil {
			fmt.Printf("最终失败: %s - %v\n\n", url, err)
		}
	}
}

// 示例6: 完整的网页爬虫系统
func example6_complete_crawler() {
	fmt.Println("\n=== 示例6: 完整爬虫系统 ===")
	
	type Crawler struct {
		maxWorkers  int
		maxDepth    int
		visited     map[string]bool
		visitedLock sync.RWMutex
		urlQueue    chan URL
		results     chan PageContent
		wg          sync.WaitGroup
	}
	
	NewCrawler := func(maxWorkers, maxDepth int) *Crawler {
		return &Crawler{
			maxWorkers: maxWorkers,
			maxDepth:   maxDepth,
			visited:    make(map[string]bool),
			urlQueue:   make(chan URL, 100),
			results:    make(chan PageContent, 100),
		}
	}
	
	isVisited := func(c *Crawler, url string) bool {
		c.visitedLock.RLock()
		defer c.visitedLock.RUnlock()
		return c.visited[url]
	}
	
	markVisited := func(c *Crawler, url string) bool {
		c.visitedLock.Lock()
		defer c.visitedLock.Unlock()
		
		if c.visited[url] {
			return false
		}
		c.visited[url] = true
		return true
	}
	
	worker := func(c *Crawler, id int, ctx context.Context) {
		defer c.wg.Done()
		
		for {
			select {
			case url, ok := <-c.urlQueue:
				if !ok {
					return
				}
				
				if url.Depth > c.maxDepth {
					continue
				}
				
				// 检查是否已访问
				if !markVisited(c, url.Address) {
					continue
				}
				
				// 抓取页面
				fmt.Printf("Worker %d: 抓取 [深度 %d] %s\n", id, url.Depth, url.Address)
				time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)
				
				// 模拟生成链接
				var links []string
				if url.Depth < c.maxDepth {
					numLinks := rand.Intn(3)
					for i := 0; i < numLinks; i++ {
						link := fmt.Sprintf("%s/page%d", url.Address, i+1)
						if !isVisited(c, link) {
							links = append(links, link)
						}
					}
				}
				
				page := PageContent{
					URL:     url.Address,
					Content: fmt.Sprintf("内容: %s", url.Address),
					Links:   links,
				}
				
				c.results <- page
				
				// 添加新链接到队列
				for _, link := range links {
					c.urlQueue <- URL{Address: link, Depth: url.Depth + 1}
				}
				
			case <-ctx.Done():
				return
			}
		}
	}
	
	start := func(c *Crawler, ctx context.Context, startURL string) <-chan PageContent {
		// 启动 workers
		for i := 1; i <= c.maxWorkers; i++ {
			c.wg.Add(1)
			go worker(c, i, ctx)
		}
		
		// 发送起始 URL
		c.urlQueue <- URL{Address: startURL, Depth: 0}
		
		// 等待完成后关闭 channels
		go func() {
			c.wg.Wait()
			close(c.results)
		}()
		
		return c.results
	}
	
	// 创建爬虫
	crawler := NewCrawler(3, 2)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// 开始爬取
	results := start(crawler, ctx, "https://example.com")
	
	// 收集结果
	pageCount := 0
	for page := range results {
		pageCount++
		fmt.Printf("📄 页面 %d: %s (%d 个链接)\n", pageCount, page.URL, len(page.Links))
	}
	
	fmt.Printf("\n总共抓取 %d 个页面\n", pageCount)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("======================================")
	fmt.Println("    并发网页爬虫实战")
	fmt.Println("======================================")
	
	example1_simple_scraper()
	example2_rate_limited_scraper()
	example3_depth_first_crawler()
	example4_breadth_first_crawler()
	example5_scraper_with_retry()
	example6_complete_crawler()
	
	fmt.Println("\n所有示例运行完成！")
}
