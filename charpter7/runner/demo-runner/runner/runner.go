package runner

import (
	"os"
	"os/signal"
	"errors"
)

type Runner struct {
	interrupt chan os.Signal
	complete chan error
	tasks []func(int)
}

func New() *Runner {
	return &Runner{
		interrupt: make(chan os.Signal),
		complete: make(chan error),
	}
}

func (r *Runner) Start() {
	signal.Notify(r.interrupt, os.Interrupt)

	r.run()
}

var ErrTimeout = errors.New("error timeout")

func (r *Runner) run() error {
	for id, task := range r.tasks {
		// got interrupt
		if (r.gotInterrupt()) {
			return ErrTimeout
		}

		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {


	select {
	case <- r.interrupt:
		// stop other signals
		
		return true
	default:
		return false
	}
}
