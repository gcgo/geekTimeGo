package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRunPrime(t *testing.T) { startTime := time.Now()
	var result []int
	for num := 2; num <= 200000; num++ {
		if isPrime(num) {
			result = append(result, num)
		} }
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时:", finishTime.Sub(startTime))
}
