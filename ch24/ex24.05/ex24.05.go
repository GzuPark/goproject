package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	const numJob int = 10
	var jobList [numJob]Job

	for i := 0; i < numJob; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(numJob)

	for i := 0; i < numJob; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}

	wg.Wait()
}
