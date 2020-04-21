package gpool

type Pool struct {
	work chan func()
	sem  chan struct{}
}

func New(size int) Pool {
	return Pool{
		work: make(chan func(), size),
		sem:  make(chan struct{}, size),
	}
}

func (p Pool) Schedule(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p Pool) worker(task func()) {
	defer func() { <-p.sem }()

	task()

	for task := range p.work {
		task()
	}
}
