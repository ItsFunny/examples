package concurrent

import (
	"fmt"
)

/*
	线程池的golang实现版本
	角色:
		dispatcher: 任务的分发者,所有的任务都是通过这个发送给旗下的worker
		worker: 工作者
		jobQueue: 存放所有的工作
		WorkerQueue: 存放空闲的工作者
	原理就是通过任务会统一的给一个dispatcher,disPatcher内部会拥有一个bufferQueue的成员变量用于缓存任务
	当任务将缓存占满之后便会发送(模仿tcp中的nagle算法),后台会有一个线程从bufferQueue中取数据
	取得的数据又会给workerQueue中的channel,因为worker对象持有引用,所以会接收到消息从而执行
	注意点:
		worker中对自身jobQueue的处理,在每次执行之前都要重新入队,因为在添加任务的时候(dispatcher)是出队添加的,因而
		需要重新入队
 */

var jobBufferQueue chan interface{}

type Worker struct {
	workerId      int            // 演示,临时添加
	totalConsumer int            // 演示,临时添加
	jobQueue      chan interface{}
	workerQueue   chan chan interface{}
}

type Dispatcher struct {
	workerQueue chan chan interface{}
	Workers     []*Worker
}

func NewWorker(workerId int, workerQueue chan chan interface{}) *Worker {
	return &Worker{workerId: workerId, workerQueue: workerQueue, jobQueue: make(chan interface{})}
}

func (w *Worker) Consume(job interface{}) {
	w.totalConsumer++
	fmt.Printf("worker:[%d]consume the job:%v,totalCosumed:%d \n", w.workerId, job, w.totalConsumer)
}

func (w *Worker) Start() {
	for {
		w.workerQueue <- w.jobQueue
		select {
		case job := <-w.jobQueue:
			// if ok {
			w.Consume(job)
			// }
		}
	}
}

func (w *Worker) Summary() {
	fmt.Printf("worker:[%d]consume totalCosumed:%d \n", w.workerId, w.totalConsumer)
}

func NewDispatcher(workerQueue chan chan interface{}, workers []*Worker) *Dispatcher {
	return &Dispatcher{workerQueue: workerQueue, Workers: workers}
}

func (d *Dispatcher) Start() {
	for i := 0; i < len(d.Workers); i++ {
		go func(index int) {
			d.Workers[index].Start()
		}(i)
	}
	for {
		select {
		case jobBuffer := <-jobBufferQueue:
			go func(job interface{}, workerQueue chan chan interface{}) {
				jobQueue := <-workerQueue
				jobQueue <- job
			}(jobBuffer, d.workerQueue)
		}
	}
}

// 总结函数,实际而定
func (d *Dispatcher) Summary() {
	total:=0
	for i := 0; i < len(d.Workers); i++ {
		d.Workers[i].Summary()
		total+=d.Workers[i].totalConsumer
	}
	fmt.Println(total)
}

func AddJob(job interface{}) {
	jobBufferQueue <- job
}

func init() {
	jobBufferQueue = make(chan interface{}, 1024)
}
