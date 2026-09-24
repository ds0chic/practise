package practise

import (
	"context"
	"fmt"
)

func ValueDemo() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 13145)
	fmt.Println(ctx.Value("userID"))

}
