package pool

import (
	"io"
	"errors"
	"log"
	"sync"
)

type Pool struct {
	m sync.Mutex
	resources chan io.Closer
	factory func() (io.Closer, error)
}

var err = errors.New("Pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size < 0 {
		return nil, errors.New("Size cannot be less than 0")
	}

	return &Pool{
		factory: fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire", "++++++++++ Shared Resource:")
		if !ok {
			log.Println("Pool has been closed")
			return nil, err
		}
		return r, nil
	default:
		log.Println("Acquire", "------------ New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	select {
	case p.resources <- r:
		log.Println("rrrrrrrrrrrr Release", ":++++++In Queue:")
	default:
		r.Close()
		log.Println("cccccccccccc Release", "-------Closing")
	}
}
