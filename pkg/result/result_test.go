package result

import (
	"skillbox_diploma/pkg/simulator"
	"sync"
	"testing"
	"time"
)

// go test -race
func TestResultMultithreading(t *testing.T) {
	go simulator.StartSimulatorServer()
	time.Sleep(5 * time.Second)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			_ = GetResultData()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			_ = GetResultData()
		}
	}()

	wg.Wait()
}
