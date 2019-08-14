package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var allDone sync.WaitGroup
var consumeCount uint32
var timeoutCount uint32

/*
Alloc: 表示demo程序申请还未被垃圾回收器回收的内存
HeapIdle: 表示demo程序的垃圾回收器持有且没有归还给操作系统的内存
*/
func nowStats() string {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	return fmt.Sprintf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes) NumGoroutine:%d", ms.Alloc, ms.HeapIdle, ms.HeapReleased, runtime.NumGoroutine())
}

func produce(ch chan int) {
	ch <- 1
}

func consume(ch chan int) {
	t := time.NewTimer(5 * time.Second)
	select {
	case <-ch:
		atomic.AddUint32(&consumeCount, 1)
		//case <-time.After(5 * time.Second):
	case <-t.C:
		atomic.AddUint32(&timeoutCount, 1)
	}
	t.Stop()
	allDone.Done()
}

func main() {
	//log.Printf("program begin. %s\n", nowStats())
	//
	//for i := 0; i < 1000*1000; i++ {
	//	allDone.Add(1)
	//	ch := make(chan int, 1)
	//	go consume(ch)
	//	go produce(ch)
	//}
	//allDone.Wait()
	//
	//runtime.GC()
	//log.Printf("all consumer done. consume count:%d timeoutcount:%d stats:%s\n",
	//	atomic.LoadUint32(&consumeCount), atomic.LoadUint32(&timeoutCount), nowStats())
	//log.Printf("sleep...")
	//time.Sleep(10 * time.Second)
	//
	//runtime.GC()
	//log.Printf("program end. %s\n", nowStats())

	//ch := make(chan int)
	//go func() {
	//	outer:
	//	for {
	//		select {
	//		case d, ok := <-ch:
	//			if !ok {
	//				fmt.Println("select break")
	//				break outer
	//			}
	//			fmt.Printf("select:%d\n", d)
	//			//default:
	//			//	fmt.Printf("default\n")
	//		}
	//	}
	//	//for d := range ch{
	//	//	fmt.Printf("select:%d\n", d)
	//	//}
	//}()
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		ch <- i
	//	}
	//	fmt.Println("sleep one second")
	//	time.Sleep(time.Second)
	//	close(ch)
	//}()
	//time.Sleep(time.Minute)

	//var arr = make([]int,0,10)
	//doSomeHappyThings(arr)
	//fmt.Println(arr, len(arr), cap(arr), "after return")
	//
	//m := make(map[int]int,6)
	//m[1] = 1

	//var ch = make(chan int, 100)
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		ch <- 1
	//	}
	//}()
	//
	//for {
	//	// the wrong part
	//	if len(ch) == 100 {
	//		sum := 0
	//		itemNum := len(ch)
	//		for i := 0; i < itemNum; i++ {
	//			sum += <-ch
	//		}
	//		if sum == itemNum {
	//			return
	//		}
	//	}
	//}

	runtime.GOMAXPROCS(runtime.NumCPU())
	go server()
	go printNum()
	var i = 3
	for {
		// will block here, and never go out
		fmt.Printf("main:%d\n",i)
		i++
	}
	fmt.Println("for loop end")
	time.Sleep(time.Second * 3600)
}

func printNum() {
	i := 0
	for {
		fmt.Println(i)
		i++
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func server() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func doSomeHappyThings(arr []int) {
	arr = append(arr, 1)
	fmt.Println(arr, "after append")
}