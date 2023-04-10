package main

import "fmt"

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results) // each routine is executed by one process
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for number := range tasks {
		results <- fibonnaci(number)
	}
}

func fibonnaci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonnaci(position-2) + fibonnaci(position-1)
}
