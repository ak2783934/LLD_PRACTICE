package main

/*

Modify this so workers can be gracefully stopped using context.Context



*/

// type Job struct {
// 	id int
// }

// func process(job Job, workerID int) {
// 	time.Sleep(700 * time.Millisecond)

// 	fmt.Println("processing the job ", job.id, "on worker ", workerID)
// }

// func worker(jobs <-chan Job, id int, wg *sync.WaitGroup, ctx context.Context) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case job, ok := <-jobs:
// 			if !ok {
// 				return
// 			}
// 			process(job, id)
// 		case <-ctx.Done():
// 			fmt.Println("cancel call came for worker ", id)
// 			return
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	jobChan := make(chan Job, 10)
// 	for i := 0; i < 3; i++ {
// 		go worker(jobChan, i+1, &wg, ctx)
// 	}

// 	// publish 10 tasks and should be consumed by different
// 	// go routines worker and complete the processing.

// 	for j := 0; j < 10; j++ {
// 		jobChan <- Job{id: j}
// 	}

// 	close(jobChan)
// 	wg.Wait()
// }
