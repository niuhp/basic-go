package ch14

import (
	"testing"
	"sync"
)

func TestCounter(t *testing.T) {

	var mut sync.Mutex
	var wg sync.WaitGroup

	count := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			count++
		}()
	}

	t.Logf("before wait,count = %d", count)

	wg.Wait()

	t.Logf("after wait,count = %d", count)
}
