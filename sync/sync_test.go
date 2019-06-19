package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment the counter 3 times leaves it at 4", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			//sync.WaitGroup which is a convenient way of synchronising concurrent processes
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		//By waiting for wg.Wait() to finish before making our assertions we can be sure all of our goroutines
		// have attempted to Inc the Counter
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func NewCounter() *Counter {
	return &Counter{}
}
