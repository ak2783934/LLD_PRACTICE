package main

/*
Implement a worker pool with:
- 3 workers
- 10 jobs
- Each job prints worker ID and job ID
*/

// type Job struct {
// 	id int
// }

// func process(job Job, workerID int) {
// 	fmt.Println("processing the job ", job.id, "on worker ", workerID)
// }

// func worker(jobs <-chan Job, id int, wg *sync.WaitGroup) {
// 	for job := range jobs {
// 		func() {
// 			defer wg.Done()
// 			process(job, id)
// 		}()
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(10)
// 	jobChan := make(chan Job, 10)
// 	for i := 0; i < 3; i++ {
// 		go worker(jobChan, i+1, &wg)
// 	}

// 	// publish 10 tasks and should be consumed by different
// 	// go routines worker and complete the processing.

// 	for j := 0; j < 10; j++ {
// 		jobChan <- Job{id: j}
// 	}

// 	close(jobChan)
// 	wg.Wait()
// }
