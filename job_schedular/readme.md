Design a job scheduling system that allows clients to submit jobs and have them executed at the right time


Functional requirements: 
- Able to schedule jobs at any time 
- jobs can be one time or repeated
- how the execution happens is not our concern. 
- implement retries in case of any failures of jobs with backoffs. 
- Jobs can be cancelled as well at any time. 
- 

Entities
Job{
    id string
    payload string
    cronExpression string
}

JobExecution {
    executionID
    jobID
    runTime time.Time
    retryCount int
    status string // enequed, processing, retrying, completed, cancelled, failed
}

JobSchedular{
    jobs []*Job
    JobExecutions []*JobExecution
    jobCount int
    MaxRetryCount int
}

func CreateJob() (string, error) {
    create the job, and schedule one job for us. 
}
func CancelJob(jobID string) error {
    findAllJobExecution and cancell them. 
}

fun RunCron(){
    run a ticker and keep checking for execution time. 
    time.Ticker {
        if any jobs matches with given timestamp, pick those, 
        fetch details of the job
        execute it and shcedule the next job if needed based on cron expresson?
        also keeps updating the retry and other things for on job
    }
}


JobExecuter interface {}
Execute(Job)(error)



- How to handle jobs? repeated jobs? 
    - do we keep adding next run date for these? 









Interfaces


APIs




