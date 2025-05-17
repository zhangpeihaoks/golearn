package main

import (
	"log/slog"
	"sync"
)

func Processer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- (id+1)*100 + i
		slog.Info("processer send msg", "id", id, "msg", (id+1)*100+i)
	}
}

func Consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		slog.Info("consumer receive msg", "id", id, "msg", msg)
	}

}

const (
	ProcesserCount = 5
	ConsumerCount  = 2
	BufferSize     = 10
)

func main() {
	ch := make(chan int, 0)
	var processWg sync.WaitGroup
	processWg.Add(ProcesserCount)
	for i := 0; i < ProcesserCount; i++ {
		go Processer(i, ch, &processWg)
	}

	var consumerWg sync.WaitGroup
	consumerWg.Add(ConsumerCount)
	for i := 0; i < ConsumerCount; i++ {
		go Consumer(i, ch, &consumerWg)
	}

	go func() {
		processWg.Wait()
		close(ch)
		slog.Info("all processer done")
	}()

	consumerWg.Wait()
	slog.Info("all consumer done")
}
