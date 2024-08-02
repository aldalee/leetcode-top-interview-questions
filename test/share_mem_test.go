package main

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func asyncService() chan string {
	// retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result:", ret)
		retCh <- ret
		fmt.Println("Service exited.")
	}()
	return retCh
}

func task() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	t.Log(service())
	task()
}

func TestAsyncService(t *testing.T) {
	retCh := asyncService()
	task()
	t.Log(<-retCh)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
