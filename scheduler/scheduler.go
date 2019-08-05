package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChannel chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	s.workChannel <- request
}

func (s *SimpleScheduler) Run() {
	s.workChannel = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}


func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workChannel
}
