package engine

type ConCurrentEngine struct {
	Scheduler   Scheduler
	ChanCount   int
	WorkerReady ReadyNotifier
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Run()
	Submit(request Request)
	WorkerChan() chan Request
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConCurrentEngine) Run(sends ...Request) {
	out := make(chan ParserResult)

	// 先把调度器跑起来
	c.Scheduler.Run()

	// 创建worker
	for i := 0; i < c.ChanCount; i++ {
		c.CreateWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	// 提交任务给调度器
	for _, row := range sends {
		c.Scheduler.Submit(row)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//fmt.Printf("%v", item)
			// 小心坑
			go func(item Item) {
				c.ItemChan <- item
			}(item)
		}
		//fmt.Println()

		for _, row := range result.Requests {
			// TODO: 需要加个去重
			c.Scheduler.Submit(row)
		}
	}
}

// CreateWorker 输入：Request ; 输出：ParserResult
func (c *ConCurrentEngine) CreateWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			// 告诉scheduler worker已经ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
