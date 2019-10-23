package engine

import (
	"fmt"
)

type ConCurrentEngine struct {
	Scheduler   Scheduler
	ChanCount   int
	WorkerReady ReadyNotifier
}

type Scheduler interface {
	Run()
	Submit(request Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConCurrentEngine) Run(sends ...Request) {
	out := make(chan ParserResult)

	c.Scheduler.Run()
	for i := 0; i < c.ChanCount; i++ {
		c.CreateWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}
	for _, row := range sends {
		c.Scheduler.Submit(row)
	}

	for {
		result := <-out
		// TODO: 需要改为存储
		for _, row := range result.Items {
			fmt.Printf("%v", row)
		}
		fmt.Println()

		for _, row := range result.Requests {
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
