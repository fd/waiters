package waiters

import "sync"

// Waiter allowes multiple goroutines to Wait until they are triggered.
type Waiter struct {
	mtx sync.RWMutex
	cha chan struct{}
}

// WaitC returns the channel to block receive from. The returned channel is
// closed when the Waiter is triggered.
func (w *Waiter) WaitC() <-chan struct{} {
	w.mtx.RLock()
	if c := w.cha; c != nil {
		w.mtx.RUnlock()
		return c
	}
	w.mtx.RUnlock()

	w.mtx.Lock()
	defer w.mtx.Unlock()

	if w.cha != nil {
		return w.cha
	}

	w.cha = make(chan struct{})
	return w.cha
}

// Trigger the waiter and unblock all waiting goroutines.
func (w *Waiter) Trigger() {
	w.mtx.Lock()
	defer w.mtx.Unlock()

	if w.cha == nil {
		return
	}

	close(w.cha)
	w.cha = make(chan struct{})
}
