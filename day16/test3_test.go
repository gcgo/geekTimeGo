package main

import (
	"log"
	"time"
)

func main() {
	intCh := make(chan int, 10)
	workerCnt := 10
	for i := 0; i < workerCnt; i++ {
		go func(i int) {
			intCh <- i
			log.Println(i, "开始工作", time.Now())
		}(i)
	}
	time.Sleep(2 * time.Second)
	for o := 0; o < workerCnt; o++ {
		go func(o int) {
			out := <-intCh
			log.Printf("%d拿到%d\n", o, out)
		}(o)
	}
	time.Sleep(2 * time.Second)
}
