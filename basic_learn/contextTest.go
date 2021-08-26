package main

import (
	"context"
	"fmt"
	"time"
)

func test2() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		go func(index int) {
			time.Sleep(time.Second * 3)
			select {
			case <-ctx.Done():
				fmt.Println("任务已经别人结束")
			default:
				fmt.Printf("%d 来结束任务!", index)
				cancel()
			}
		}(i)
		go func() {
			select {
			case <-time.After(time.Second):
				fmt.Println("1s 过去了")
				go func(second time.Duration) {
					time.Sleep(time.Second * second)
					select {
					case <-ctx.Done():
						fmt.Println("任务已经别人结束")
					default:
						fmt.Printf("特使来结束任务!")
						cancel()
					}
				}(1)
			case <-ctx.Done():
				fmt.Println("任务已经结束")
			}
		}()
	}
	for {
	}
}

func test3() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func main() {
	test2()
}
