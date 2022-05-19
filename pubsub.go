package main

import (
	"fmt"
	"sync"
)

var (
	topic   = make(chan int, 10)
	data    = make(chan int)
	done    = make(chan bool)
	mutex   sync.Mutex
	console = make([]string, 0)
)

func publisher() {
	fmt.Println("Starting Publisher")
	defer func() {
		fmt.Println("Publisher is done")
		close(topic)
	}()
	for {
		select {
		case v, ok := <-data:
			if !ok {
				return
			}
			fmt.Println("Published message:", v)
			topic <- v
		}
	}
}

func subscriber(i int) {
	fmt.Println("Starting Subscriber no.", i)
	for msg := range topic {
		fmt.Println("Received message:", msg, "by subscriber no.", i)
	}

	fmt.Println("Subscriber", i, "is done")
	done <- true
}

func main() {
	go publisher()
	numSubscriber := 7
	for i := 0; i < numSubscriber; i++ {
		go subscriber(i + 1)
	}

	for i := 1; i <= 100; i++ {
		data <- (i * i)
	}

	close(data)

	for i := 0; i < numSubscriber; i++ {
		select {
		case <-done:
			//do nothing
		}
	}
	fmt.Println("Main is done")
}
