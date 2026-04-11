package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, tasks <-chan string, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d остановлен: %v\n", id, ctx.Err())
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
			select {
			case result <- fmt.Sprintf("Воркер %d обработал: %s", id, task):
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tasks := make(chan string, 10)
	result := make(chan string, 10)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, tasks, result, &wg)
	}

	go func() {
		jobs := []string{"task1", "task2", "task3", "task4", "task5", "task6"}
		for _, j := range jobs {
			select {
			case tasks <- j:
			case <-ctx.Done():
				close(tasks)
				return
			}
		}
		close(tasks)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println(res)
	}

	if ctx.Err() != nil {
		fmt.Println("Работа прервана по таймауту")
	}
}
