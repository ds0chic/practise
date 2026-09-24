package practise

import (
	"context"
	"fmt"
	"time"
)

func TimeoutDemo() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
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
