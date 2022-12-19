package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for i := 1; i <= 10; i++ {
		jobName := fmt.Sprintf("job_%v", i)
		go job(ctx, jobName)
	}

	time.Sleep(time.Minute * 1)
}

func job(ctx context.Context, name string) {

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("job %v done!!!!!!\n", name)
			return
		default:
			fmt.Printf("doing job %v\n", name)
			time.Sleep(time.Second)
		}
	}
}
