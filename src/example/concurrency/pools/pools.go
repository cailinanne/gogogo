package main

import (
	"fmt"
)

type Job struct {
	worker, input, output int
}

func (j Job) String() string {
	return fmt.Sprintf("worker:[%d] input:[%d] output:[%d]", j.worker, j.input, j.output)
}


func worker(id int, jobs <-chan Job, results chan<- Job) {
	fmt.Printf("Worker %d starting \n", id)
	for j := range jobs {
		j.worker = id
		j.output = j.input * id
		fmt.Printf("Worker %d processing job %v\n", id, j)
		results <- j
	}
}

func main(){
	jobs := make(chan Job, 100)
	results := make(chan Job, 100)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Put 10 jobs in the queue
	numJob := 10
	for i := 0; i < numJob; i++{
		jobs <- Job{0, i, 0}
	}
	close(jobs)

	// There is no super elegant way to close the results channel, so
	// we can't do a range on results.  We just have to "know" how
	// many results there are to fetch.
	for i := 0; i < numJob; i++ {
		job := <- results
		fmt.Printf("Got result %v \n", job )
	}

}
