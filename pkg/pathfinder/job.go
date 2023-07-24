package pathfinder

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Job struct {
	Config       *Config
	Mutex        sync.Mutex
	Rate         *rate.Limiter
	Runner       *Runner
	ErrorCounter int
	Done         int
	Total        int
	Running      bool
	queue        []QueueJob
}

type QueueJob struct {
	URL     string
	Payload string
}

func NewJob(config *Config) *Job {
	var job Job

	job.Rate = rate.NewLimiter(rate.Every(time.Second), int(config.Rate))
	job.Running = false
	job.Config = config
	job.Done = 0
	job.Total = 0
	job.ErrorCounter = 0
	job.queue = make([]QueueJob, 0)

	return &job
}

func (job *Job) RateLimit(ctx context.Context, URL string) error {
	if job.checkError() {
		return errors.New("Hit maximum error rate")
	}

	if err := job.Rate.Wait(ctx); err != nil {
		return err
	}

	resp, err := job.Runner.Fetch(URL)
	if err != nil {
		return err
	}

	if resp.StatusCode == 403 {
		job.addError()
	}

	resp.Compare()
	return nil
}

func (job *Job) Start() error {
	var wg sync.WaitGroup

	fmt.Printf("[*] Running %d total queued jobs\n", job.Total)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	job.Runner = NewRunner(job.Config)
	threads := make(chan QueueJob, job.Config.Threads)

	for i := 0; i < job.Config.Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range threads {
				select {
				case <-ctx.Done():
					return
				default:
					if err := job.RateLimit(ctx, task.URL); err != nil {
						// Avoid spamming context errors with all the threads
						if err != context.Canceled {
							fmt.Printf("[-] Error: %v\n", err)
						}
						job.stop(cancel)
						return
					}
					job.addDone()
				}
			}
		}()
	}

	for _, tasks := range job.queue {
		if !job.Runner.Ready {
			if err := job.Runner.CheckReady(tasks.URL); err == nil {
				job.start()
			} else {
				return err
			}
		}

		if !job.Running {
			break
		}

		threads <- tasks
	}

	close(threads)
	wg.Wait()
	return nil
}

func (job *Job) Queue(URL string, payload string) {
	var queue QueueJob

	queue.URL = URL
	queue.Payload = payload
	job.queue = append(job.queue, queue)
	job.Total++
}

func (job *Job) checkError() bool {
	if job.Done > 50 && job.Config.StopError && (float64(job.ErrorCounter)/float64(job.Done) > 0.75) {
		return true
	}
	return false
}

func (job *Job) start() {
	job.Mutex.Lock()
	defer job.Mutex.Unlock()
	job.Running = true
	job.Runner.Ready = true
}

func (job *Job) stop(cancel context.CancelFunc) {
	job.Mutex.Lock()
	defer job.Mutex.Unlock()
	job.Running = false
	cancel()
}

func (job *Job) addDone() {
	job.Mutex.Lock()
	defer job.Mutex.Unlock()
	job.Done++
}

func (job *Job) addError() {
	job.Mutex.Lock()
	defer job.Mutex.Unlock()
	job.ErrorCounter++
}
