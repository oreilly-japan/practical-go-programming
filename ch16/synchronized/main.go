package main

import (
	"sync"
	"sync/atomic"
)

// sample:synchronized
type Account struct {
	balance int
	lock    sync.RWMutex
}

func (a *Account) GetBalanace() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.balance
}

func (a *Account) Transfer(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
}

// sample:synchronized

// sample:synchronized2
type Account2 struct {
	balance  int
	transfer chan int
}

func NewAccount() *Account2 {
	transfer := make(chan int)
	r := &Account2{
		transfer: transfer,
	}
	go func() {
		for {
			amount := <-transfer
			r.balance += amount
		}
	}()
	return r
}

func (a Account2) GetBalanace() int {
	return a.balance
}

func (a *Account2) Transfer(amount int) {
	a.transfer <- amount
}

// sample:synchronized2

// sample:synchronized3
type Account3 struct {
	balance int64
}

func (a Account3) GetBalanace() int64 {
	return a.balance
}

func (a *Account3) Transfer(amount int64) {
	atomic.AddInt64(&a.balance, amount)
}

// sample:synchronized3
