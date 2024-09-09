package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	initEmployees()

	workers := 3

	chJobs := make(chan *Employee, len(employees))
	chResults := make(chan string, len(employees))
	chErrors := make(chan error, len(employees))

	for i := 0; i < len(employees); i++ {
		chJobs <- &employees[i]
	}
	close(chJobs)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go IncreaseSalaries(i, &wg, chJobs, chResults, chErrors)
	}

	go func() {
		wg.Wait()
		close(chErrors)
		close(chResults)
	}()

	fmt.Println("RESULTS:")
	for v := range chResults {
		fmt.Println(v)
	}

	fmt.Println("ERRORS:")
	for v := range chErrors {
		fmt.Println(v)
	}

}

func IncreaseSalaries(id int, wg *sync.WaitGroup, jobs <-chan *Employee, results chan<- string, errs chan<- error) {
	defer wg.Done()

	for job := range jobs {

		rng := rand.Intn(100)

		// Simulate an error in 1/3 of cases
		if rng > 33 {
			job.Salary = job.Salary + job.Salary*0.2
			results <- fmt.Sprintf("Goroutine №%d: Successfuly!", id)
		} else {
			errs <- fmt.Errorf("goroutine №:%d Failed", id)
		}

		time.Sleep(1 * time.Second)
	}
}
