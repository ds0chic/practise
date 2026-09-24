package practise

import (
	"context"
	"fmt"
)

func ParentChildDemo() {
	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithCancel(parent)
	defer cancelParent()
	defer cancelChild()
	cancelParent()
	<-child.Done()
	fmt.Println("child canceled")
	fmt.Println(child.Err())
}
