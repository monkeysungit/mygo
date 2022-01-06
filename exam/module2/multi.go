package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

//生产者
func produce(out chan<- int, nu int) {
	for {
		cond.L.Lock()
		//产品区满 等待消费者消费
		for len(out) == 3 {
			cond.Wait()
		}
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%dth ***producer produce***,num = %d,len(chan) = %d\n", nu, num, len(out))
		cond.L.Unlock()

		//生产了产品唤醒 消费者线程
		cond.Signal()
		//生产完了歇一会，给其他协程机会
		time.Sleep(time.Second)
	}
}

//消费者
func consume(in <-chan int, nu int) {
	for {
		cond.L.Lock()
		//产品区空 等待生产者生产
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("%dth ###consumer consume###,num = %d,len(chan) = %d\n", nu, num, len(in))
		cond.L.Unlock()
		cond.Signal()

		//消费完了歇一会，给其他协程机会
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	//产品区 使用channel模拟
	product := make(chan int, 3)

	//创建互斥锁和条件变量
	cond.L = new(sync.Mutex)

	//5个消费者
	for i := 0; i < 5; i++ {
		go produce(product, i)
	}
	//3个生产者
	for i := 0; i < 3; i++ {
		go consume(product, i)
	}

	//主协程阻塞 不结束
	<-quit
}
