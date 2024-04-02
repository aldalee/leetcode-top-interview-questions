package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	cancelCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Done")
		}(i, cancelCh)
	}
	cancel(cancelCh)
	time.Sleep(time.Second * 1)
}

func cancel(cancelCh chan struct{}) {
	// cancelCh <- struct{}{}
	close(cancelCh)
}

func isCancelled(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}
