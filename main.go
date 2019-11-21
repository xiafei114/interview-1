package main

import (
	"errors"
	"sync"
	"sync/atomic"
)

func main() {}


type Task interface{
	Do() error
}

type WorkerPool struct {
	cap uint32
	running uint32
	Workers []*Worker  // 总的工人数
	Workings []*Worker // 正在工作的工人数
	cond *sync.Cond
	lock sync.Locker
}

func (w *WorkerPool) Cap() uint32 {
	return w.cap
}

func (w *WorkerPool) Running() uint32 {
	return w.running
}


func (w *WorkerPool) addRunning() {
	atomic.AddUint32(&w.running, 1)
}

func (w *WorkerPool) addCap() {
	atomic.AddUint32(&w.cap, 1)
}

func (w *WorkerPool) subCap() {
	atomic.AddUint32(&w.cap, ^uint32(0))
}

func (w *WorkerPool) recycle(worker *Worker) {
	w.Workers = append(w.Workers, worker)
	w.addCap()
}

var wp *WorkerPool

// 创建协程池
func NewWorkerPool(cap int) *WorkerPool {
	if wp != nil {
		return wp
	}

	lock := &sync.RWMutex{}
	cond := sync.NewCond(lock)

	return &WorkerPool{
		cap:uint32(cap),
		running: 0,
		Workers: make([]*Worker, cap, cap),
		Workings: make([]*Worker, 0, cap),
		lock: lock,
		cond: cond,
	}
}

func (w *WorkerPool) Get() *Worker {
	w.lock.Lock()
	if w.cap > 0 {
		worker := w.Workers[0:1]
		w.Workers = w.Workers[1:]
		w.subCap()
		w.lock.Unlock()
		return worker[0]
	} else {
	waiting:
		w.cond.Wait()

		if w.cap > 0 {
			worker := w.Workers[0:1]
			if worker != nil {
				w.lock.Unlock()
				return worker[0]
			}
		}

		goto waiting
	}
}

func (w *Worker) Run(task Task) (*Worker, error) {
	if w == nil {
		return nil, errors.New("worker is empty")
	}

	w.mu.Lock()
	w.Task = task
	w.working = true
	err := task.Do()

	if err != nil {
		w.working = false
		w.mu.Unlock()
		return w, err
	}

	w.working = false
	w.mu.Unlock()
	return w, nil
}

type Worker struct {
	Task
	mu sync.Mutex
	working bool
}

type TaskPool chan Task


