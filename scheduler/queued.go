package scheduler

import "crawler/engine"

type QueueScheduler struct {
	requestChannel chan engine.Request
	workerChannel  chan chan engine.Request
}

func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(request engine.Request) {
	q.requestChannel <- request
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workerChannel <- w
}

func (q *QueueScheduler) Run() {
	q.requestChannel = make(chan engine.Request)
	q.workerChannel = make(chan chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue [] chan engine.Request
		for {
			var activityRequest engine.Request
			var activityWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activityRequest = requestQueue[0]
				activityWorker = workerQueue[0]
			}
			select {
			case r := <-q.requestChannel:
				requestQueue = append(requestQueue, r)
			case w := <-q.workerChannel:
				workerQueue = append(workerQueue, w)
			case activityWorker <- activityRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
