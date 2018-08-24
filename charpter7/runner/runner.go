package runner

import (

	"time"
	"os"
	"os/signal"
	"errors"
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

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {

	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <- r.timeout:
		return ErrTimeout
		// return timeout error
	}
	
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// check for an interrupt signal from the OS
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		
		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <- r.interrupt:
		// sth
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
