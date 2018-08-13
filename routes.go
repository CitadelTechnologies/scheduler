package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"scheduler/scheduler"
	"scheduler/server"
)

type(
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}
	Routes []Route
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
		go func(router *mux.Router, route Route){
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(ErrorHandler(route.HandlerFunc))
		}(router, route)
    }
    return router
}

func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer server.CatchException(w)
        next(w, req)
    })
}

var routes = Routes{
	Route{
		"Create Job",
		"POST",
		"/jobs",
		scheduler.CreateJobAction,
	},
	Route{
		"Get Jobs",
		"GET",
		"/jobs",
		scheduler.GetJobsAction,
	},
	Route{
		"Get Job",
		"GET",
		"/jobs/{id}",
		scheduler.GetJobAction,
	},
}
