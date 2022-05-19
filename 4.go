package main

import (
	//"fmt"

	"context"
	"log"
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
	time.AfterFunc(3*time.Second, cancel)
	mySleepAndTalk(ctx, 5*time.Second, "Hello there!")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	log.Println("Started...")
	select {
	case <-time.After(d):
		log.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
