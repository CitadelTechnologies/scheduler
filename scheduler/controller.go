package scheduler

import(
    "github.com/gorilla/mux"
    "scheduler/server"
    "net/http"
)

func CreateJobAction(w http.ResponseWriter, r *http.Request) {
    data := server.DecodeJsonRequest(r)
    job := CreateJob(
        data["id"].(string),
        data["method"].(string),
        data["url"].(string),
        data["executed_at"].(string),
    )
    Scheduling.AddJob(job)

    server.SendJsonResponse(w, 201, job)
}

func GetJobAction(w http.ResponseWriter, r *http.Request) {
    server.SendJsonResponse(w, 200, Scheduling.GetJob(mux.Vars(r)["id"]))
}

func GetJobsAction(w http.ResponseWriter, r *http.Request) {
    jobs := make([]*Job, 0, len(Scheduling.Jobs))

    for  _, job := range Scheduling.Jobs {
       jobs = append(jobs, job)
    }
    server.SendJsonResponse(w, 200, jobs)
}
