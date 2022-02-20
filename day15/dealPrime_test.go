package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRunPrime(t *testing.T) {
	startTime := time.Now()
	var result []int
	for num := 2; num <= 200000; num++ {
		if isPrime(num) {
			result = append(result, num)
		}
	}
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时:", finishTime.Sub(startTime))
}

func TestRunPrime2(t *testing.T) {
	startTime := time.Now()
	var result []int
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		result = append(result, collectPrime(2, 100000)...)
	}()
	go func() {
		defer wg.Done()
		result = append(result, collectPrime(100001, 200000)...)
	}()
	wg.Wait()
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时:", finishTime.Sub(startTime))
}
