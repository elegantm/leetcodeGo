package leetcode

import (
	"fmt"
	"sync"
)

func thousandeRun() {
	goNum := 100
	queueChan := make([]chan int, goNum) //多个 channel 用来控制协程的执行顺序
	for i := 0; i < goNum; i++ {
		queueChan[i] = make(chan int, 1)
	}

	exitChan := make(chan int, 1)

	res := 1
	ch := 0

	for k := 0; k < goNum; k++ {
		go func(index int) {
			for { // 协程内部是一个循环

				<-queueChan[index]
				if res > 1000 {
					exitChan <- 1
					break // 达到条件后退出
				}

				fmt.Println("go routine ", index+1, "  : ", res)

				res += 1

				if ch >= goNum-1 {
					ch = 0
				} else {
					ch += 1
				}

				queueChan[ch] <- 1
			}
		}(k)

	}

	queueChan[0] <- 1 // 给第一个 channel 发数据 开启任务

	<-exitChan
	fmt.Println("exit")

}

func thousandTwo(max int) {
	lock := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	num := 100

	wg.Add(num)
	res := 1
	for k := 0; k < num; k++ {
		go func(index int) {
			for {

				lock.Lock()
				now := res
				lock.Unlock()

				if now > max {
					wg.Done()
					break
				}

				if (now-1)%num == index {
					fmt.Println("go routine", index+1, now)
					res += 1
				}

			}

		}(k)

	}
	wg.Wait()

}
