package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
)

//func demoFunc() {
//	time.Sleep(10 * time.Millisecond)
//	fmt.Println("Hello World!")
//}

var sum int32

type tt struct {
	num int
}

func myFunc(i interface{}) {
	n := i.(tt)
	atomic.AddInt32(&sum, int32(n.num))
	fmt.Printf("run with %d\n", n)
}

func main() {
	defer ants.Release()

	runTimes := 1000
	wg := new(sync.WaitGroup)

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	// 协程池容量10，超时时间1s
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(tt{num: i})
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
