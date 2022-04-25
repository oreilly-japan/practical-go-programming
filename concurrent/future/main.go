package main

import (
	"log"
	"time"
)

// sample:testcase
func A() int {
	time.Sleep(time.Second)
	return 10
}

func B() int {
	time.Sleep(time.Second * 2)
	return 5
}

func C(a, b int) int {
	time.Sleep(time.Second * 1)
	return a + b
}

func D(a, c int) int {
	time.Sleep(time.Second)
	return a + c
}

// sample:testcase

// sample:define
type Future struct {
	value int
	wait  chan struct{}
}

func (f *Future) IsDone() bool {
	select {
	case <-f.wait:
		return true
	default:
		return false
	}
}

func (f *Future) Get() int {
	<-f.wait
	return f.value
}

type Promise struct {
	f *Future
}

func (p *Promise) Submit(v int) {
	p.f.value = v
	close(p.f.wait)
}

func MakeFuturePromise() (*Future, *Promise) {
	f := &Future{
		wait: make(chan struct{}),
	}
	p := &Promise{
		f: f,
	}
	return f, p
}

// sample:define

func future() {
	// sample:usage
	fa, pa := MakeFuturePromise()
	fb, pb := MakeFuturePromise()
	fc, pc := MakeFuturePromise()
	fd, pd := MakeFuturePromise()

	go func() {
		a := A()
		pa.Submit(a)
	}()
	go func() {
		b := B()
		pb.Submit(b)
	}()
	go func() {
		c := C(fa.Get(), fb.Get())
		pc.Submit(c)
	}()
	go func() {
		d := D(fa.Get(), fc.Get())
		pd.Submit(d)
	}()
	log.Printf("d: %d\n", fd.Get())
	// sample:usage
}

/*
func main() {
	future()
}
*/
