package main

import (
	"log"
	"sync"
	"time"
	"sync/atomic"
	"errors"
)

const (
    DEFAULT_EXPIRY_DURATION = 1
)

type sig struct{}

type Pool struct {
	// capacity of the pool.
	capacity int32

	// running is the number of the currently running goroutines.
	running int32

	// expiryDuration set the expired time (second) of every worker.
	expiryDuration time.Duration

	// workers is a slice that store the available workers.
	workers []*Worker

	workerCache sync.Pool

	// I donot know how to use this chan var
	// freeSignal chan sig

	// release is used to notice the pool to closed itself.
	release int32

	// lock for synchronous operation.
	lock sync.Mutex	
}

func NewPool(size int32) (*Pool, error){
	if size <= 0 {
		return nil, errors.New("size is invalid.")
	}
	pool := Pool{
		capacity: size,
		expiryDuration: DEFAULT_EXPIRY_DURATION,
	}
	return &pool, nil
}

func (p *Pool) getWorker() *Worker{
	var w *Worker
	waiting := false
	p.lock.Lock()
	workers := p.workers
	n := len(workers) - 1
	// log.Printf("n is %d", n)
	if n < 0 {
		// 没有空闲worker
		if p.running >= p.capacity {
			// 并且工作的worker已达上限，只能等�
			waiting = true
		} else {
			// 但是工作的worker未达上限，可以创建一个新的worker进行工作
			// p.running++
		}
	} else {
		// 有空闲worker，从队列尾部取出一个使�
		w = workers[n]
		workers[n] =  nil
		p.workers = workers[:n]
	}
	p.lock.Unlock()

	if waiting {
		for {
			p.lock.Lock()
			workers = p.workers
			l := len(workers) - 1
			// log.Printf("[waiting] workers' length is %d", l)
			if l < 0 {
				p.lock.Unlock()
				continue
			}
			w = workers[l]
			workers[l] = nil
			p.workers = workers[:l]
			p.lock.Unlock()
			break
		}
	} else {
		if cacheWorker := p.workerCache.Get(); cacheWorker != nil {
			log.Printf("Get cacheworker!!!!")
			w = cacheWorker.(*Worker)
		} else {
			w = &Worker{
				pool: p,
				task: make(chan func()),
			}
		}
		w.run()
	}
	return w
}

func (p *Pool) Submit(task func()) error {
	if p.release > 0 {
		return errors.New("Pool is closed, fail to submit task.")
	}
	w := p.getWorker()
	w.sendTask(task)
	return nil
}

func (p *Pool) putWorker(w *Worker) {
	p.lock.Lock()
	p.workers = append(p.workers, w)
	p.lock.Unlock()
}

type Worker struct {
	// pool who owns this worker.
	pool *Pool

	// task is a job should be done.
	task chan func()

	// recycleTime will be update when putting a worker back into queue.
	recycleTime time.Time
}

func (w *Worker) run() {
	atomic.AddInt32(&w.pool.running, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				atomic.AddInt32(&w.pool.running, -1)
				w.pool.workerCache.Put(w)				
			}
		}()

		for f := range w.task {
			if f == nil {
				atomic.AddInt32(&w.pool.running, -1)
				w.pool.workerCache.Put(w)
			}
			f()
			w.pool.putWorker(w)
		}
	}()
}

func (w *Worker) sendTask(task func()) {
	w.task <- task
}

//--------------------------test--------------------------
func demoFunc() {
	log.Printf("Hello world!")
}
func main() {
	size := 10
	runtimes := 500
	var wg sync.WaitGroup
	pool,_ := NewPool(int32(size))
	dfun := func() {
		demoFunc()
		wg.Done()
	}
	for i:= 0; i < runtimes; i++ {
		wg.Add(1)
		pool.Submit(dfun)
	}
	wg.Wait()
}

