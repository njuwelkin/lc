package worker_pool

import (
	//"fmt"
	"sync"
)

type JobStatus int

type Job interface {
	Do()
}

type WorkerPool struct {
	jobChannel  chan Job
	quitChannel chan bool
	workers     int
	chanSize    int

	waitGroup sync.WaitGroup

	running      bool
	runningMutex sync.Mutex
}

func NewWorkerPool(workers int, chanSize int) *WorkerPool {
	return &WorkerPool{
		workers:  workers,
		chanSize: chanSize,
		running:  false,
	}
}

func (workerPool *WorkerPool) execute() {
	workerPool.waitGroup.Add(1)
	defer workerPool.waitGroup.Done()
	for {
		// two layer of select to make sure jobChannel's priority is heigher
		// worker can quit only when no jobs pending
		select {
		case job := <-workerPool.jobChannel:
			job.Do()
		default:
			select {
			case job := <-workerPool.jobChannel:
				job.Do()
			case <-workerPool.quitChannel:
				//fmt.Println("quit")
				return
			}
		}
	}
}

func (workerPool *WorkerPool) Run() *WorkerPool {
	workerPool.runningMutex.Lock()
	defer workerPool.runningMutex.Unlock()

	if !workerPool.running {
		workerPool.jobChannel = make(chan Job, workerPool.chanSize)
		workerPool.quitChannel = make(chan bool)
		for i := 0; i < workerPool.workers; i++ {
			go workerPool.execute()
		}
		workerPool.running = true
	}
	return workerPool
}

func (workerPool *WorkerPool) InsertJob(job Job) {
	workerPool.runningMutex.Lock()
	defer workerPool.runningMutex.Unlock()

	if workerPool.running {
		workerPool.jobChannel <- job
	}
}

func (workerPool *WorkerPool) Quit() {
	workerPool.runningMutex.Lock()
	defer workerPool.runningMutex.Unlock()

	if workerPool.running {
		workerPool.running = false
		for i := 0; i < workerPool.workers; i++ {
			workerPool.quitChannel <- true
		}
		// block until all works quit
		workerPool.waitGroup.Wait()
		close(workerPool.jobChannel)
		close(workerPool.quitChannel)
	}
}

type simpleJob struct {
	f func()
}

func (j *simpleJob) Do() {
	j.f()
}

func (workerPool *WorkerPool) InsertFuncJob(f func()) {
	workerPool.InsertJob(&simpleJob{f: f})
}
