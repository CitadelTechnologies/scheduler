package scheduler

import(
    "scheduler/exception"
    "time"
)

type Scheduler struct {
    Jobs map[string]*Job
}

type Job struct {
    Id string `json:"id"`
    Method string `json:"method"`
    Url string `json:"url"`
    ExecutedAt time.Time `json:"executed_at"`
    Timer *time.Timer `json:"-"`
}

func (s *Scheduler) AddJob(job *Job) {
    s.Jobs[job.Id] = job
}

func (s *Scheduler) GetJob(id string) *Job {
    if job, isset := s.Jobs[id]; isset != false {
        return job
    }
    panic(exception.New(404, "Job not found", nil))
}

func (s *Scheduler) RemoveJob(id string) {
    if _, isset := s.Jobs[id]; !isset {
        panic(exception.New(404, "Job not found", nil))
    }
    delete(s.Jobs, id)
}

func (s *Scheduler) CancelJob(id string) {
    job := s.GetJob(id)
    job.Timer.Stop()
    delete(s.Jobs, id)
}
