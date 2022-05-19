package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func prasanth(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: prasanth handler started")
	defer fmt.Println("server: prasanth handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "prasanth\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
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
	sleepAndTalk(ctx, 3*time.Second, "kudumala")
	http.HandleFunc("/prasanth", prasanth)
	http.ListenAndServe(":8090", nil)

}
