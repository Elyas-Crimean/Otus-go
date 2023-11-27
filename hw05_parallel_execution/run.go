package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var (
		errCountM sync.Mutex
		errCount  int
		wg        sync.WaitGroup
		err       error
	)
	ch := make(chan Task)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range ch {
				err := t()
				if err != nil {
					errCountM.Lock()
					errCount++
					errCountM.Unlock()
				}
			}
		}()
	}
	for i := range tasks {
		errCountM.Lock()
		if errCount >= m && m != 0 {
			errCountM.Unlock()
			break
		}
		errCountM.Unlock()
		ch <- tasks[i]
	}
	close(ch)
	wg.Wait()
	if errCount >= m && m != 0 {
		err = ErrErrorsLimitExceeded
	}
	return err
}
