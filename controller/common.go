package controller

type response struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Result  any			`json:"result"`
}

func newResponse(status int, message string, result interface{}) *response {
    return &response{status, message, result}
}