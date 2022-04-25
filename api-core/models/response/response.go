package response

var ParamValidException = Exception{
	Code:    400,
	Message: "param valid error",
}

type Exception struct {
	Code    int
	Message string
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Get(code int, message string, data interface{}) *Response {
	return &Response{
		code,
		message,
		data,
	}
}

func Success(data interface{}) *Response {
	return Get(200, "success", data)
}

func FromException(exception *Exception) *Response {
	return Get(exception.Code, exception.Message, nil)
}

func FromError(err error) *Response {
	return Get(500, err.Error(), nil)
}
