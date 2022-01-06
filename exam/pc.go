package main

import "time"

func main() {
	producer := make(chan int, 10)
	defer close(producer)
	for i := 0; i < 10; i++ {
		producer <- i
		//time.Sleep(time.Second)
		println("producer", i)
	}
	go func() {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			println("consumer: ", <-producer)
		}
	}()
	time.Sleep(5 * time.Second)
}
