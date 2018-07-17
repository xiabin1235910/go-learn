package runner

import (

	"time"
	"os"
)

type Runner struct {
	interrupt chan os.Signal
	
	complete chan error

	timeout <-chan time.Time
	
	tasks []func(int)
}

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {

	r.complete <- r.run()

	select {
	case err := <-r.complete:
		return err
	case <- r.timeout:
		// return timeout error
	}
	
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// check for an interrupt signal from the OS
		
		task(id)
	}

	return nil
}
