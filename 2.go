package main

import (
	//"fmt"

	"bufio"
	"context"
	"log"
	"os"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	log.Println("Started...")
	select {
	case <-time.After(d):
		log.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()
	sleepAndTalk(ctx, 5*time.Second, "Hello there!")
}
