package main

import (
	"fmt"
	"time"
)

type Task struct {
	Id   string
	Name string
}

type Example struct {
	TaskChan chan Task
	Tk       *time.Ticker
}

func (e *Example) Init() {
	e.TaskChan = make(chan Task)
	//e.TaskChan <- Task{}
	//e.Tk = time.NewTicker(time.Microsecond)
}

func (e *Example) DoSomething() {
	select {
	case <-e.Tk.C:
		fmt.Println(time.Now())
		return
	case <-e.TaskChan:
		return
	}
}

func main() {
	e := &Example{}
	e.Init()
	sleepChan := make(chan bool)

	go func() {
		e.Tk = time.NewTicker(3*time.Second)
		defer e.Tk.Stop()
		for {
			select {
			case <-e.Tk.C:
				close(e.TaskChan)
				fmt.Println(time.Now())
			case r :=<-e.TaskChan:
				fmt.Printf("%v",r)
			}
		}
	}()
	e.TaskChan <- Task{Id: "123", Name: "hao"}

	<-sleepChan
}
