package main

type Job struct {
	str string
}

func main() {

}

func Worker(id int, chTask <-chan Job, chResult chan<- int) {

	for job := range chTask {
		charCount := len([]rune(job.str)) // Task: count the number of characters
		chResult <- charCount
	}
}

func WorkerPool(numWorker int)
