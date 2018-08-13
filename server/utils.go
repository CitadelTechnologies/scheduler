package server

import(
    "scheduler/exception"
    "encoding/json"
    "io"
    "io/ioutil"
    "log"
    "net/http"
)

func DecodeJsonRequest(r *http.Request) map[string]interface{} {
    var body []byte
    var err error
    if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
        panic(exception.New(http.StatusInternalServerError, "Request body could not be opened", err))
    }
    if err = r.Body.Close(); err != nil {
        panic(exception.New(http.StatusInternalServerError, "Request body could not be closed", err))
    }
    var data map[string]interface{}
    if err = json.Unmarshal(body, &data); err != nil {
        panic(exception.New(http.StatusUnprocessableEntity, "Request body could not be parsed", err))
    }
    return data
}

func SendJsonResponse(w http.ResponseWriter, code int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    if err := json.NewEncoder(w).Encode(&data); err != nil {
        panic(err)
    }
}

func CatchException(w http.ResponseWriter) {
    r := recover()
    if r == nil {
        return
    }
    if exception, ok := r.(*exception.Exception); ok {
        message := ""
        if exception.Error != nil {
            message = "; [Error]: " + exception.Error.Error()
        }
        if exception.Message != "" || message != "" {
            log.Println("[Exception]: " + exception.Message + message)
        }
        SendJsonResponse(w, exception.Code, exception)
        return
    }
    if err, ok := r.(error); ok {
        log.Println("[Error]: " + err.Error())
        SendJsonResponse(w, 500, "Internal server error")
        return
    }
    panic(r)
}
