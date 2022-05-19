package main

import (
	//"fmt"

	"context"
	"log"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	log.Println("Started...")

	val, ok := ctx.Value("name").(string)
	if !ok {
		log.Println("name not present in context")
		return
	}
	log.Println("Name provided:", val)

	select {
	case <-time.After(d):
		log.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx,
		time.Now().Add(10*time.Second))
	ctx = context.WithValue(ctx, "firstname", "santak")
	defer cancel()
	sleepAndTalk(ctx, 5*time.Second, "Hello there!")
}
