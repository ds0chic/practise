package practise

import (
	"context"
	"fmt"
	"time"
)

func TaskB(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("TaskB 收到取消信号")
			fmt.Println(ctx.Err())
			return

		default:
			fmt.Println("TaskB 正在运行")
			time.Sleep(time.Second)
		}
	}
}

func TaskA(ctx context.Context) {
	fmt.Println("TaskA 开始")

	TaskB(ctx)
}

func TaskChainDemo() {
	ctx, cancel := context.WithCancel(context.Background())

	go TaskA(ctx)

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)
}
