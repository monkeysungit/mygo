package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func main() {
	q := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
	for i := 0; i < 3; i++ {
		go q.Enqueue("虎", i)
	}
	for i := 0; i < 3; i++ {
		go q.Dequeue(i)
	}
	quit := make(chan bool)
	//主协程阻塞 不结束
	<-quit
}

func (q *Queue) Enqueue(item string, nu int) {
	for {
		q.cond.L.Lock()
		if len(q.queue) == 5 {
			q.cond.Wait()
		}
		q.queue = append(q.queue, item)
		fmt.Printf("%dth - putting %s to queue, %v\n", nu, item, q.queue)
		q.cond.L.Unlock()
		q.cond.Signal()
		time.Sleep(time.Second * 2)
	}
}

func (q *Queue) Dequeue(nu int) {
	for {
		time.Sleep(time.Second)
		q.cond.L.Lock()
		if len(q.queue) == 0 {
			q.cond.Wait()
		}
		result := q.queue[0]
		q.queue = q.queue[1:]
		fmt.Printf("%dth - getting %s to queue, %v\n", nu, result, q.queue)
		q.cond.L.Unlock()
		q.cond.Signal()
		time.Sleep(time.Second)
	}
}
