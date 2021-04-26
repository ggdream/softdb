package pool

import (
	"sync"
)


const (
	buffer = 1 << 13
)

type (
	Drip = func()

	Pool struct {
		count     int
		gainQueue chan Drip
		castQueue chan Drip

		closed		bool
		mu sync.Mutex
		finish	chan bool
	}
)

func New(count int) *Pool {
	return &Pool{
		count:     count,
		gainQueue: make(chan Drip, buffer),
		castQueue: make(chan Drip, buffer),
		closed: false,
		finish: make(chan bool),
	}
}

// gain2cast 水滴中转函数，需要使用gouortine注册
func (p *Pool) gain2cast() {
	// 如果没有close，将是个死循环
	// 当Stop -> close()时，遍历完自动退出循环
	for drip := range p.gainQueue {
		p.castQueue <- drip
	}

	// 注册完所有任务后，close执行管道
	close(p.castQueue)
}

// flow let a worker work 水滴流动
func (p *Pool) flow() {
	for drip := range p.castQueue {
		drip()
	}

	p.finish <- true
}

// Add add a task
func (p *Pool) Add(drip Drip) {
	// 如果调用了Stop(), 则不再接受添加任务
	if p.closed {
		return
	}
	p.gainQueue <- drip
}


// Run start to exec the pool, async
func (p *Pool) Run() {
	go p.gain2cast()
	for i := 0; i < p.count; i++ {
		go p.flow()
	}
}

func (p *Pool) Stop() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}

	close(p.gainQueue)
	p.closed = true
}

// Wait sync wait, block the process
func (p *Pool) Wait() {
	for i := 0; i < p.count; i++ {
		// 阻塞，等待子gouortine完成任务
		<- p.finish
	}
}
