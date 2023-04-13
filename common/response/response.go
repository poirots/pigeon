package response

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "SUCCESS",
		Data:    data,
	}
}

func Error(errorMsg string) *Response {
	return &Response{
		Code:    -1,
		Message: errorMsg,
	}
}

func ErrorWithCode(code int, errorMsg string) *Response {
	return &Response{
		Code:    code,
		Message: errorMsg,
	}
}
