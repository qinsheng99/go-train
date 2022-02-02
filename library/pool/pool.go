package pool

import "sync"

type GoFuncPool struct {
	MaxLimit  int
	Gochannel chan struct{}
	sy        sync.WaitGroup
}

//var GoFunc *GoFuncPool
//
//func init() {
//	GoFunc = NewGoPool(WithMaxLimit(10))
//}

type GoFuncPoolOptions func(pool *GoFuncPool)

func WithMaxLimit(max int) GoFuncPoolOptions {
	return func(pool *GoFuncPool) {
		pool.MaxLimit = max

		pool.Gochannel = make(chan struct{}, pool.MaxLimit)

		for i := 0; i < pool.MaxLimit; i++ {
			pool.Gochannel <- struct{}{}
		}
	}
}
func NewGoPool(options ...GoFuncPoolOptions) *GoFuncPool {
	pool := &GoFuncPool{}

	for _, option := range options {
		option(pool)
	}

	return pool
}

func (p *GoFuncPool) Submit(fn func()) {
	channel := <-p.Gochannel
	p.sy.Add(1)
	go func() {
		fn()
		p.Gochannel <- channel
		defer p.sy.Done()
	}()
	p.sy.Wait()
}

func (p *GoFuncPool) Close() {
	for i := 0; i < p.MaxLimit; i++ {
		<-p.Gochannel
	}
	close(p.Gochannel)
}

func (p *GoFuncPool) Size() int {
	return len(p.Gochannel)
}
