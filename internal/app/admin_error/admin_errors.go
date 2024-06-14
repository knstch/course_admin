package adminerror

type ErrorResponse struct {
	Err     error
	Message string `json:"error"`
	Code    uint   `json:"code"`
}

func CreateErrorResponse(err error, code uint) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
		Code:    code,
	}
}
