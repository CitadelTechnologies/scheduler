package exception

type Exception struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Error error `json:"-"`
}

func New(code int, message string, err error) *Exception {
    return &Exception{
        Code: code,
        Message: message,
        Error: err,
    }
}
