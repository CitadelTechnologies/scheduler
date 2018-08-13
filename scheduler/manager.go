package scheduler

import(
    "net/http"
    "time"
)

var Scheduling *Scheduler

func init() {
    Scheduling = &Scheduler{
        Jobs: make(map[string]*Job, 0),
    }
}

func CreateJob(id, method, url, callAt string) *Job {
    executedAt, err := time.Parse(time.RFC3339, callAt)
    if err != nil {
        panic(err)
    }
    return &Job{
        Id: id,
        Method: method,
        Url: url,
        ExecutedAt: executedAt,
        Timer: time.AfterFunc(time.Until(executedAt), func() {
            PerformJobCall(method, url)
            Scheduling.RemoveJob(id)
        }),
    }
}

func PerformJobCall(method, url string) {
    _, err := http.NewRequest(method, url, nil)
    if err != nil {
        panic(err)
    }
}
