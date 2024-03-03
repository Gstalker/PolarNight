package utility

import (
	"sync"
	"time"
)

// Sleep 异步等待，避免阻塞当前rotine
func Sleep(t time.Duration) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(t)
		wg.Done()
	}()
	wg.Wait()
}
