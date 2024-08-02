package main

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

// func producer(ch chan int) {
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 		wg.Done()
// 	}()
// }

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func consumer(ch chan int) {
	for {
		if ret, ok := <-ch; ok {
			fmt.Println(ret)
			continue
		}
		break
	}
	wg.Done()
}

func TestCloseChannel(t *testing.T) {
	ch := make(chan int)
	wg.Add(3)
	go producer(ch)
	go consumer(ch)
	go consumer(ch)
	wg.Wait()
}

func TestCloseChannelError(t *testing.T) {
	ch := make(chan int)
	producer(ch)
	consumer(ch)
}
