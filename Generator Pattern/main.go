package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	canceled := false
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	go func() {
		<-sc
		cancel()
		canceled = true
	}()
	for i := range Count(ctx, 1, 99) {
		fmt.Println(i)
	}

	if canceled {
		fmt.Println(os.Stderr, canceled)
	}
}

func Count(ctx context.Context, start int, end int) <-chan int {
	ch := make(chan int)
	go func(ch chan int) {
		defer close(ch)
	loop:
		for i := start; i <= end; i++ {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
	}(ch)
	return ch
}
