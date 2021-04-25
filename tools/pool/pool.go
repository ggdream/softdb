package pool

const (
	buffer = 1 << 13
)

type (
	Drip = func()

	Pool struct {
		count     int
		gainQueue chan Drip
		castQueue chan Drip
	}
)

func New(count int) *Pool {
	return &Pool{
		count:     count,
		gainQueue: make(chan Drip, buffer),
		castQueue: make(chan Drip, buffer),
	}
}

// gain2cast 水滴中转函数，需要使用gouortine注册
func (p *Pool) gain2cast() {
	for drip := range p.gainQueue {
		p.castQueue <- drip
	}
}

// flow let a worker work 水滴流动
func (p *Pool) flow() {
	for drip := range p.castQueue {
		drip()
	}
}

// Add add a task
func (p *Pool) Add(drip Drip) {
	p.gainQueue <- drip
}

// Wait sync wait, block the process
func (p *Pool) Wait() {}

// Run start to exec the pool, async
func (p *Pool) Run() {
	go p.gain2cast()
	for i := 0; i < p.count; i++ {
		go p.flow()
	}
}
