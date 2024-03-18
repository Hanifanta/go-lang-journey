package Golang_Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncTask(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Run async task")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncTask(group)
	}

	group.Wait()
	fmt.Println("Done")
}
