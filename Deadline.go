package practise

import (
	"context"
	"fmt"
	"time"
)

func DeadlineDemo() {
	ctx := context.Background()
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println("任务中止")
			return
		default:
			fmt.Println("任务运行中")
			time.Sleep(time.Second)
		}
	}
}
