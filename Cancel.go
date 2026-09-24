package practise

import (
	"context"
	"fmt"
	"time"
)

func CancelDemo() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务结束")
				return
			default:
				fmt.Println("任务运行")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
