package concurrent

import (
	"fmt"
	"testing"
	"time"
)

var maxWorker = 10 // 最多10个工作者
var maxJobs = 100000

func TestDispatcher_Start(t *testing.T) {
	workerQueue := make(chan chan interface{}, maxWorker)
	workers := make([]*Worker, 0)
	for i := 0; i < maxWorker; i++ {
		workers = append(workers, NewWorker(i, workerQueue))
	}
	dispatcher := NewDispatcher(workerQueue, workers)

	go func() { dispatcher.Start() }()

	// 新建100000个任务
	startTime := time.Now().Unix()
	for i := 0; i < maxJobs; i++ {
		go func(index int) {
			AddJob(index)
		}(i)
	}
	time.Sleep(time.Second*10)
	fmt.Printf("%d 个任务消费完毕,耗时:%v", maxJobs, time.Now().Unix()-startTime)
	fmt.Println("===========总结函数================")
	dispatcher.Summary()
}
